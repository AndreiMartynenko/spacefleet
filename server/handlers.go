package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func spaceshipHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)

	var resp interface{}

	fmt.Println("URL ", r.URL)

	path := strings.TrimSpace(r.URL.Path)

	if r.Method == "GET" {

		name := strings.TrimSpace(r.URL.Query().Get("name"))
		class := strings.TrimSpace(r.URL.Query().Get("class"))
		status := strings.TrimSpace(r.URL.Query().Get("status"))

		searchParams := []string{}
		if name != "" {
			searchParams = append(searchParams, "Name")
		}
		if class != "" {
			searchParams = append(searchParams, "Class")
		}
		if status != "" {
			searchParams = append(searchParams, "Status")
		}

		if path == "/spaceship/" {

			crafts, err := getSpaceCrafts()

			if err != nil {
				resp = QueryStatus{Success: false}
				json.NewEncoder(w).Encode(resp)
				return
			}

			//Filter

			if len(searchParams) > 0 {
				filteredCrafts := []map[string]string{}
				for _, craft := range crafts {
					filtered := make(map[string]string)
					if name != "" {
						filtered["name"] = craft.Name
					}
					if class != "" {
						filtered["class"] = craft.Class
					}
					if status != "" {
						filtered["status"] = craft.Status
					}
					filteredCrafts = append(filteredCrafts, filtered)
					resp = filteredCrafts
				}
				json.NewEncoder(w).Encode(resp)
				return

			}

			resp = crafts

		} else {

			idStr := strings.TrimSpace(strings.TrimPrefix(path, "/spaceship/"))
			idStr = strings.Split(idStr, "&")[0]

			id, err := strconv.Atoi(idStr)
			if err != nil {
				resp = QueryStatus{Success: false}
				json.NewEncoder(w).Encode(resp)
				return
			}

			craft, err := getSpaceCraftById(id)

			if err != nil {
				fmt.Println(err)
				resp = QueryStatus{Success: false}
				json.NewEncoder(w).Encode(resp)
				return
			}

			if craft == nil {
				resp = QueryStatus{Success: false}
				json.NewEncoder(w).Encode(resp)
				return
			}

			//Filter

			if len(searchParams) > 0 {

				filtered := make(map[string]string)
				if name != "" {
					filtered["name"] = craft.Name
				}
				if class != "" {
					filtered["class"] = craft.Class
				}
				if status != "" {
					filtered["status"] = craft.Status
				}

				resp = filtered

				json.NewEncoder(w).Encode(resp)
				return
			}
			resp = craft
		}

		//POST

	} else if r.Method == "POST" {

		name := r.FormValue("name")
		class := r.FormValue("class")
		crewStr := r.FormValue("crew")
		image := r.FormValue("image")
		valueStr := r.FormValue("value")
		status := r.FormValue("status")
		armaments := r.FormValue("armaments")

		fmt.Println("name: ", name)

		fmt.Println(r.Header)

		crew, err := strconv.Atoi(crewStr)
		if err != nil {
			fmt.Println("Parse crew err ", err)
			resp = QueryStatus{Success: false}
			json.NewEncoder(w).Encode(resp)
			return

		}

		value, err := strconv.ParseFloat(valueStr, 32)
		if err != nil {
			fmt.Println("Parse value err ", err)
			resp = QueryStatus{Success: false}
			json.NewEncoder(w).Encode(resp)
			return
		}

		data := []Armament{}
		if armaments != "" {
			bytes := []byte(armaments)
			err = json.Unmarshal(bytes, &data)
			if err != nil {
				fmt.Println("Unmarshal armoments err ", err)
				resp = QueryStatus{Success: false}
				json.NewEncoder(w).Encode(resp)
				return
			}

		}

		craft := SpaceCraft{
			Name:     name,
			Class:    class,
			Crew:     crew,
			Image:    image,
			Value:    value,
			Status:   status,
			Armament: data,
		}

		err = saveSpaceShip(craft)
		if err != nil {
			fmt.Println(err)
			resp = QueryStatus{Success: false}
			json.NewEncoder(w).Encode(resp)
			return
		}
		resp = QueryStatus{Success: true}

		//DELETE

	} else if r.Method == "DELETE" {

		path := strings.TrimSpace(r.URL.Path)
		idStr := strings.TrimSpace(strings.TrimPrefix(path, "/spaceship/"))
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println(err)
			resp = QueryStatus{Success: false}
			json.NewEncoder(w).Encode(resp)
			return
		}
		err = deleteSpaceCraftById(id)

		if err != nil {
			fmt.Println(err)
			resp = QueryStatus{Success: false}
			json.NewEncoder(w).Encode(resp)
			return
		}

		resp = QueryStatus{Success: true}

		//PUT

	} else if r.Method == "PUT" {

		idStr := strings.TrimSpace(strings.TrimPrefix(path, "/spaceship/"))
		idStr = strings.Split(idStr, "&")[0]

		id, err := strconv.Atoi(idStr)

		if err != nil {
			resp = QueryStatus{Success: false}
			json.NewEncoder(w).Encode(resp)
			return
		}

		params := make(map[string]string)

		name := strings.TrimSpace(r.URL.Query().Get("name"))
		if name != "" {
			params["name"] = name
		}
		class := strings.TrimSpace(r.URL.Query().Get("class"))
		if class != "" {
			params["class"] = class
		}
		crew := strings.TrimSpace(r.URL.Query().Get("crew"))
		if crew != "" {
			params["crew"] = crew
		}
		image := strings.TrimSpace(r.URL.Query().Get("image"))
		if image != "" {
			params["image"] = image
		}
		value := strings.TrimSpace(r.URL.Query().Get("value"))
		if value != "" {
			params["value"] = value
		}
		status := strings.TrimSpace(r.URL.Query().Get("status"))
		if status != "" {
			params["status"] = status
		}

		armaments := strings.TrimSpace(r.URL.Query().Get("armaments"))
		if armaments != "" {
			params["armaments"] = armaments
		}

		err = updateSpaceship(id, params)

		if err != nil {
			fmt.Println(err)
			resp = QueryStatus{Success: false}
			json.NewEncoder(w).Encode(resp)
			return
		}
		resp = QueryStatus{Success: true}

	} else {
		resp = "error: HTTP request method not allowed"
	}

	json.NewEncoder(w).Encode(resp)

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func spaceshipHandler(w http.ResponseWriter, r *http.Request) {
	// resp := Response{Payload: nil, Error: nil}

	var resp interface{}

	fmt.Println(r.Method)

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

		value, err := strconv.Itoa(valueStr)
		if err != nil {
			fmt.Println("Parse value err ", err)
			resp = QueryStatus{Success: false}
			json.NewEncoder(w).Encode(resp)
			return
		}

	}

}

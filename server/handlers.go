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

			// path := strings.TrimSpace(r.URL.Path)
			// fmt.Println(path)

			if path == "/spaceship/" {
				crafts, err := getSpaceCrafts()

				if err != nil {
					//resp.Error = &Error{Type: ERROR_ACCESSING_DATABASE, Message: fmt.Sprintf("Error: %v", err)}
					resp = Error{Error: fmt.Sprintf("error: %v", err)}
					json.NewEncoder(w).Encode(resp)
					return
				}

				// resp.Payload = crafts
				resp = crafts

			} else {
				idStr := strings.TrimSpace(strings.TrimPrefix(path, "/spaceship/"))
				id, err := strconv.Atoi(idStr)
				if err != nil {
					resp = Error{Error: fmt.Sprintf("error: %v", err)}
					json.NewEncoder(w).Encode(resp)
					return
				}
				craft, err := getSpaceCraftById(id)
			}

			// } else {
			// 	resp.Error = &Error{Type: WRONG_HTTP_METHOD, Message: "Error: HTTP request method not allowed"}
			// }

			// json.NewEncoder(w).Encode(resp)

		}
	}
}

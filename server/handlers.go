package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func spaceshipHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Payload: nil, Error: nil}

	if r.Method == "GET" {
		crafts, err := getSpaceCrafts()

		if err != nil {
			resp.Error = &Error{Type: ERROR_ACCESSING_DATABASE, Message: fmt.Sprintf("Error: %v", err)}
			json.NewEncoder(w).Encode(resp)
			return
		}

		path := strings.TrimSpace(r.URL.Path)
		fmt.Println(path)

		if path == "/spaceship/" {
			crafts, err := getSpaceCrafts()

			if err != nil {
				resp.Error = &Error{Type: ERROR_ACCESSING_DATABASE, Message: fmt.Sprintf("Error: %v", err)}
				json.NewEncoder(w).Encode(resp)
				return
			}

			resp.Payload = crafts

		} else {
			resp.Error = &Error{Type: WRONG_HTTP_METHOD, Message: "Error: HTTP request method not allowed"}
		}

		json.NewEncoder(w).Encode(resp)

	}

}

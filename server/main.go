package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB

func main() {
	dbLocal, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/spacefleet")

	db = dbLocal

	if err != nil {
		fmt.Println(err)
	}

	// err = createTable()
	err = createArmamentTable()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir("../")))
	http.HandleFunc("/spaceship/", spaceshipHandler)
	http.ListenAndServe(":8080", nil)
}

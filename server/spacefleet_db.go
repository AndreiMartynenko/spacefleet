package main

import (
	_ "github.com/go-sql-driver/mysql"
)

//     _________Spaceships_____________________________________________________
//     | Id      | Name 	| Class   | Crew     | Image  | Value   | Status  |
//     | INTEGER | TEXT     | TEXT 	  | INTEGER  | TEXT   | INTEGER | TEXT    |

//     _______________Armament________________
//    | id 	    | Craft_id    | Title | Qty  |
//    | INTEGER |  INTEGER    | TEXT  | TEXT |

func createSpacecraftsTable() error {
	query := `CREATE TABLE IF NOT EXISTS spacecrafts (
		Id INTEGER PRIMARY KEY AUTO_INCREMENT,
		Name TEXT NOT NULL,
		Class TEXT NOT NULL,
		Crew INTEGER NOT NULL,
		Image TEXT,
		Value INTEGER NOT NULL,
		Status TEXT NOT NULL
	)
	`
	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil

}

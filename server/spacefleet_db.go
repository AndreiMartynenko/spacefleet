package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//     _________Spacecrafts_____________________________________________________
//     | Id      | Name 	| Class   | Crew     | Image  | Value   | Status  |
//     | INTEGER | TEXT     | TEXT 	  | INTEGER  | TEXT   | INTEGER | TEXT    |

//     _______________Armament________________
//    | id 	    | CraftId    | Title | Qty  |
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

func createArmamentTable() error {
	query := `CREATE TABLE IF NOT EXISTS armament (
		Id INTEGER PRIMARY KEY AUTO_INCREMENT,
		CraftId INTEGER,
		Title TEXT NOT NULL,
		Qty TEXT NOT NULL
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

func saveArmament(craftId int, Title string, Qty string) error {
	query := `INSERT INTO armament_log (CraftId, Title, Qty) VALUES(?,?,?)`

	statement, err := db.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(craftId, Title, Qty)
	if err != nil {
		return err
	}
	return nil
}

func getArmaments(craftId int) ([]Armament, error) {
	query := `
	SELECT
	Title, Qty
	FROM
	armament_log
	WHERE
	CraftId = ?
	`
	rows, err := db.Query(query, craftId)

	if err != nil {
		return nil, err
	}

	armaments := []Armament{}

	for rows.Next() {
		amarment := Armament{}

		err = rows.Scan(&amarment.Title, &amarment.Qty)
		if err != nil {
			return nil, err
		}
		armaments = append(armaments, amarment)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return armaments, nil
}

func insertData() error {

	query := `INSERT INTO spacecrafts VALUES(?,?,?,?,?,?,?,?)`

	/*
	   {
	     "data": [
	       {
	           "id": 1,
	           "name": "Devastator",
	           "status": "operational"
	       },
	       {
	           "id": 2,
	           "name": "Red Five",
	           "status": "damaged"
	       },
	   }

	   {
	       "id": 1,
	       "name": "Devastator",
	       "class": "Star Destroyer",
	       "crew": 35000,
	       "image": "https:\\url.to.image",
	       "value": 1999.99,
	       "status": "operational",
	       "armament": [
	           {
	               "title": "Turbo Laser",
	               "qty": "60"
	           },
	           {
	             "title": "Ion Cannons",
	             "qty": "60",
	           },
	           {
	             "title": "Tractor Beam",
	             "qty": "10",
	           },
	       ]
	   }

	*/
	statement, err := db.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(1, "Devastator", "Star Destroyer", `[{"title": "Turbo Laser", "qty": "60"},{"title": "Ion Cannons","qty": "60"},{"title": "Tractor Beam","qty": "10"}]`, 35000, "https:\\url.to.image", 1999.99, "operational")
	if err != nil {
		return err
	}
	return nil
}

func getSpaceCrafts() ([]SpaceCraft, error) {

	query := `SELECT
	Id, Name, Class, Crew, Image, Value, Status
	FROM
	spacecrafts`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	crafts := []SpaceCraft{}

	for rows.Next() {
		craft := SpaceCraft{}
		err = rows.Scan(&craft.Id, &craft.Name, &craft.Class, &craft.Crew, &craft.Image, &craft.Value, &craft.Status)
		if err != nil {
			return nil, err
		}
		armaments, err := getArmaments(craft.Id)

		if err != nil {
			return nil, err
		}
		craft.Armament = armaments
		crafts = append(crafts, craft)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return crafts, nil
}

func getSpaceCraftById(id int) (*SpaceCraft, error) {
	query := `SELECT
	Id, Name, Class, Crew, Image, Value, Status
	FROM
	spacecrafts
	WHERE
	Id = ? LIMIT 1`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var craft *SpaceCraft
	for rows.Next() {
		craft = &SpaceCraft{}
		err = rows.Scan(&craft.Id, &craft.Name, &craft.Class, &craft.Crew, &craft.Image, &craft.Value, &craft.Status)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	armaments, err := getArmaments(craft.Id)

	if err != nil {
		return nil, err
	}

	craft.Armament = armaments

	return craft, nil
}

func saveSpaceShip(craft SpaceCraft) error {
	query := `INSERT INTO spacecrafts
	(Name, Class, Crew, Image, Value, Status)
	VALUES
	(?,?,?,?,?,?)`

	statement, err := db.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(craft.Name, craft.Class, craft.Crew, craft.Image, craft.Value, craft.Status)
	if err != nil {
		fmt.Print(err)
		return err
	}

	craftId, err := result.LastInsertId()

	if err != nil {
		fmt.Print(err)
		return err
	}

	for _, armament := range craft.Armament {
		err = saveArmament(int(craftId), armament.Title, armament.Qty)
		if err != nil {
			fmt.Print(err)
			return err
		}
	}

	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

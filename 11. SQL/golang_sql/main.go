package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// go get -u github.com/go-sql-driver/mysql

type Colors struct {
	Id    string
	Color string
}

func main() {
	conn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_sample")

	if err != nil {
		panic(err)
	}

	selecting, err := conn.Query("SELECT * FROM colors")

	if err != nil {
		panic(err)
	}

	inserting, err := conn.Query("INSERT INTO colors VALUES (NULL, 'Blue'),(NULL, 'Red'),(NULL, 'Green'),(NULL, 'Yellow')")

	if err != nil {
		panic(err)
	}

	defer inserting.Close()

	if err != nil {
		panic(err)
	}

	updating, err := conn.Query("UPDATE colors SET color = 'Black' WHERE id = 1")

	if err != nil {
		panic(err)
	}

	defer updating.Close()

	tx, err := conn.Begin()

	if err != nil {
		panic(err)
	}

	prep, _ := tx.Prepare("INSERT INTO colors VALUES (NULL, ?)")

	prep.Exec("Magenta")

	tx.Rollback()
	tx.Commit()

	defer prep.Close()

	var colors Colors

	for selecting.Next() {
		err := selecting.Scan(&colors.Id, &colors.Color)

		if err != nil {
			panic(err)
		}

		fmt.Println(colors.Color)
	}

	defer conn.Close()
}

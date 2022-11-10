package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Colors struct {
	Id    string `json:"id"`
	Color string `json:"color"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"status":  "ok",
			"message": "hello world",
		}

		json.NewEncoder(w).Encode(&data)
	})

	http.HandleFunc("/data", data)
	http.HandleFunc("/update", update)

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func data(w http.ResponseWriter, r *http.Request) {
	var data []Colors
	var colors Colors

	conn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_sample")

	if err != nil {
		panic(err)
	}

	selecting, err := conn.Query("SELECT * FROM colors")

	if err != nil {
		panic(err)
	}

	for selecting.Next() {
		err := selecting.Scan(&colors.Id, &colors.Color)

		data = append(data, Colors{
			Id:    colors.Id,
			Color: colors.Color,
		})

		if err != nil {
			panic(err)
		}
	}

	selecting.Close()

	json.NewEncoder(w).Encode(data)
}

func update(w http.ResponseWriter, r *http.Request) {
	conn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_sample")

	if err != nil {
		panic(err)
	}

	tx, err := conn.Begin()

	if err != nil {
		panic(err)
	}

	updating, err := tx.Prepare("UPDATE colors SET color = ? WHERE id = ?")

	if err != nil {
		panic(err)
	}

	updating.Exec("Grey", "2")

	tx.Commit()

	updating.Close()

	data := map[string]string{
		"status":  "ok",
		"message": "updated",
	}

	json.NewEncoder(w).Encode(data)
}

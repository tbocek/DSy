package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var (
	db *sql.DB
)

func main() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=password dbname=postgres host=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8081...")
	err = http.ListenAndServe(":8081", nil)
	fmt.Printf("error: %v\n", err)
}

func handler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var name string
	for rows.Next() {
		var id []uint8
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	fmt.Fprintf(w, "<html><body><h1>Hello, "+name+"!</h1></body></html>")
}

package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var (
	db *sql.DB
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=password dbname=postgres host=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/api/users", getUsers)

	http.ListenAndServe(":8081", nil)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

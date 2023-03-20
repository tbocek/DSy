// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Text: "OST",
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(data)
	})

	port := "8082"
	fmt.Printf("API server listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

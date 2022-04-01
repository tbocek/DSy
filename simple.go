package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var rnd = 0

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
	rnd = rand.Int()
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %d", rnd)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Print("listinging on :" + os.Args[1])
	log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
}

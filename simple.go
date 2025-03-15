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
	rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	rnd = rand.Int()
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<body>
    <h1>Service: %d</h1>
</body>
</html>`, rnd)
		w.Write([]byte(htmlContent))
		fmt.Print("serving /\n")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`"Hi, all good!"`))
		fmt.Print("serving /health\n")
	})

	log.Print("listening on :" + os.Args[1])
	log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
}

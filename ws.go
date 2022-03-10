package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Date struct {
	Now string `json:"now"`
}

func ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	messageType, wsr, err := conn.NextReader()
	if err != nil {
		log.Printf("NextReader error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	req, err := io.ReadAll(wsr)
	if err != nil {
		log.Printf("ReadAll error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Sent from the browser: [%v]", string(req))

	go func(conn *websocket.Conn) {
		for {
			w, err := conn.NextWriter(messageType)
			if err != nil {
				log.Printf("NextWriter error: %v", err)
				conn.Close()
				return
			}

			json, err := json.Marshal(Date{
				Now: time.Now().Format(time.RFC3339),
			})
			if err != nil {
				log.Printf("cannot marshal json: %v", err)
				conn.Close()
				return
			}

			w.Write(json)
			w.Close()
			time.Sleep(1 * time.Second)
		}
	}(conn)
}

func main() {
	log.Print("Starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/ws", ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type DateSSE struct {
	Now string `json:"now"`
}

func sse(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Flush the headers to establish SSE connection
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	} else {
		log.Printf("Streaming not supported")
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Create a ticker to send time updates
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Detect if client closes the connection
	notify := r.Context().Done()

	for {
		select {
		case <-notify:
			log.Println("Client closed connection")
			return
		case <-ticker.C:
			dateJSON, err := json.Marshal(DateSSE{
				Now: time.Now().Format(time.RFC3339),
			})
			if err != nil {
				log.Printf("Error marshaling JSON: %v", err)
				continue
			}

			fmt.Fprintf(w, "data: %s\n\n", dateJSON)

			// Flush the data to the client
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		}
	}
}

func main() {
	log.Print("Starting SSE server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "sse.html")
	})

	http.HandleFunc("/sse", sse)

	log.Printf("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)



func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, map[string]string{
			"status":  "ok",
			"service": "1",
		})
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, map[string]string{
			"status": "healthy",
		})
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, map[string]string{
			"message": "Hello from Service 1",
		})
	})

	log.Println("Service 1 listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func jsonResponse(w http.ResponseWriter, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

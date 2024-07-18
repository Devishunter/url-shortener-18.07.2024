package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/expand", expandHandler)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

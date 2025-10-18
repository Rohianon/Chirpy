package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

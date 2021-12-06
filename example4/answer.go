package main

import (
	"log"
	"net/http"
)

func main() {

	// ListenAndServe starts an HTTP server with a given address and
	// handler defined in NewRouter.
	log.Println("Starting service on port 8080")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

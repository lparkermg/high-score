package main

import (
	"errors"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load services environment config

	// Setup logging
	log.SetPrefix("high-score: ")

	// Setup handler functions

	// Setup server
	err := http.ListenAndServe(":8888", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed.")
	} else if err != nil {
		log.Fatalf("There was a problem starting the server. Error: %s", err)
		os.Exit(1)
	}
}

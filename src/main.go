package main

import (
	"errors"
	"log"
	"mrlparker/high-score/api"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load services environment config

	// Setup logging
	log.SetPrefix("high-score: ")

	// Setup handler functions
	router := gin.Default()

	router.GET("scores", api.GetScore)
	router.POST("scores", api.PostScore)

	// Setup server
	err := router.Run("localhost:8888")

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed.")
	} else if err != nil {
		log.Fatalf("There was a problem starting the server. Error: %s", err)
		os.Exit(1)
	}
}

package main

import (
	"errors"
	"fmt"
	"log"
	"mrlparker/high-score/api"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Load services environment config
	viper.SetEnvPrefix("highscore_service")
	viper.AutomaticEnv()

	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "8888")

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.ReadInConfig()

	// Setup logging
	log.SetPrefix("high-score: ")

	// Setup handler functions
	router := gin.Default()

	router.GET("scores", api.GetScore)
	router.POST("scores", api.PostScore)

	// Setup server
	host := viper.GetString("host")
	port := viper.GetString("port")
	err := router.Run(fmt.Sprintf("%v:%v", host, port))

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed.")
	} else if err != nil {
		log.Fatalf("There was a problem starting the server. Error: %s", err)
		os.Exit(1)
	}
}

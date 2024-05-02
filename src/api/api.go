package api

import (
	data "mrlparker/high-score/data-access"
	validators "mrlparker/high-score/validators"
	"net/http"
)

func GetScore(responseWriter http.ResponseWriter, request *http.Request) {
	apiKey := request.Header.Get("Scorer-Api-Key")
	gameId := request.Header.Get("Scorer-Game-Id")

	valid, err := validators.ValidateGameAndApi(gameId, apiKey)

	if !valid && err != nil {
		// Get request is invalid, we should let the user know.
		responseWriter.WriteHeader(404)
		return
	}

	// Request is valid, we should call the data access layer.
	data.GetScores("", 0, 10)

	responseWriter.WriteHeader(200)
}

func PostScore(responseWriter http.ResponseWriter, request *http.Request) {
	apiKey := request.Header.Get("Scorer-Api-Key")
	gameId := request.Header.Get("Scorer-Game-Id")

	gameApiValid, gameApiErr := validators.ValidateGameAndApi(gameId, apiKey)

	if !gameApiValid && gameApiErr != nil {

		// Post request is invalid, we should let the user know.
		responseWriter.WriteHeader(404)
		return
	}

	valid, err := validators.ValidateGameAndScore(gameId, -5)

	if !valid && err != nil {
		// Post request is invalid, we should let the user know.
		responseWriter.WriteHeader(404)
		return
	}

	// Request is valid, we should call the data access layer.
	data.PostScore(gameId, 0)
	responseWriter.WriteHeader(201)
}

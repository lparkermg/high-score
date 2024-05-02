package api

import (
	data "mrlparker/high-score/data-access"
	validators "mrlparker/high-score/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetScore(ctx *gin.Context) {
	apiKey := ctx.Request.Header.Get("Scorer-Api-Key")
	gameId := ctx.Request.Header.Get("Scorer-Game-Id")

	valid, err := validators.ValidateGameAndApi(gameId, apiKey)

	if !valid && err != nil {
		// Get request is invalid, we should let the user know.
		ctx.Status(http.StatusNotFound)
		return
	}

	// Request is valid, we should call the data access layer.
	data.GetScores("", 0, 10)

	ctx.Status(http.StatusOK)
}

func PostScore(ctx *gin.Context) {
	apiKey := ctx.Request.Header.Get("Scorer-Api-Key")
	gameId := ctx.Request.Header.Get("Scorer-Game-Id")

	gameApiValid, gameApiErr := validators.ValidateGameAndApi(gameId, apiKey)

	if !gameApiValid && gameApiErr != nil {

		// Post request is invalid, we should let the user know.
		ctx.Status(http.StatusNotFound)
		return
	}

	valid, err := validators.ValidateGameAndScore(gameId, -5)

	if !valid && err != nil {
		// Post request is invalid, we should let the user know.
		ctx.Status(http.StatusNotFound)
		return
	}

	// Request is valid, we should call the data access layer.
	data.PostScore(gameId, 0)
	ctx.Status(http.StatusCreated)
}

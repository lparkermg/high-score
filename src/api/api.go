package api

import (
	"log"
	data "mrlparker/high-score/data-access"
	validators "mrlparker/high-score/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Api specific models
type PostModel struct {
	Score int    `json:"score"`
	Name  string `json:"name"`
}

func GetHealth(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func GetScore(ctx *gin.Context) {
	apiKey := ctx.Request.Header.Get("Scorer-Api-Key")
	gameId := ctx.Request.Header.Get("Scorer-Game-Id")
	skip, skipErr := strconv.Atoi(ctx.Request.URL.Query().Get("skip"))
	take, takeErr := strconv.Atoi(ctx.Request.URL.Query().Get("take"))

	if skipErr != nil {
		skip = 0
	}

	if takeErr != nil {
		take = 10
	}

	valid, err := validators.ValidateGameAndApi(gameId, apiKey)

	if !valid && err != nil {
		// Get request is invalid, we should let the user know.
		log.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	// Request is valid, we should call the data access layer.
	scores, scoresErr := data.GetScores(gameId, skip, take)

	if scoresErr != nil {
		log.Println(scoresErr)
		ctx.Status(http.StatusNotFound)
	}

	ctx.IndentedJSON(http.StatusOK, scores)
}

func PostScore(ctx *gin.Context) {
	apiKey := ctx.Request.Header.Get("Scorer-Api-Key")
	gameId := ctx.Request.Header.Get("Scorer-Game-Id")

	gameApiValid, gameApiErr := validators.ValidateGameAndApi(gameId, apiKey)

	if !gameApiValid && gameApiErr != nil {
		// Post request is invalid, we should let the user know.
		log.Println(gameApiErr)
		ctx.Status(http.StatusBadRequest)
		return
	}
	var requestBody PostModel

	if parseErr := ctx.BindJSON(&requestBody); parseErr != nil {
		log.Println(parseErr)
		ctx.Status(http.StatusBadRequest)
		return
	}

	valid, err := validators.ValidateGameAndScore(gameId, requestBody.Name, requestBody.Score)

	if !valid && err != nil {
		// Post request is invalid, we should let the user know.
		log.Println(err)
		ctx.Status(http.StatusNotFound)
		return
	}

	// Request is valid, we should call the data access layer.
	id, err := data.PostScore(gameId, requestBody.Name, requestBody.Score)

	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
	ctx.Value(id)
}

package request_validators

import "errors"

var validGameId = "id"
var validApiKey = "apiKey"

// Validates gameId and apiKey.
func ValidateGameAndApi(gameId string, apiKey string) (bool, error) {
	if gameId != validGameId || apiKey != validApiKey {
		return false, errors.New("game id or api key are not valid")
	}
	return true, nil
}

// Vaslidates score against gameId and the score.
func ValidateGameAndScore(gameId string, score int) (bool, error) {
	return false, errors.New("score is not within limit of game")
}

package request_validators

import "errors"

var validGameId = "id"
var validApiKey = "apiKey"

var nameMaxLength = 3
var maxScore = 500

// Validates gameId and apiKey.
func ValidateGameAndApi(gameId string, apiKey string) (bool, error) {
	if gameId != validGameId || apiKey != validApiKey {
		return false, errors.New("game id or api key are not valid")
	}
	return true, nil
}

// Vaslidates score against gameId and the score.
func ValidateGameAndScore(gameId string, name string, score int) (bool, error) {
	if len(name) > nameMaxLength {
		return false, errors.New("name is too long")
	}

	if score > maxScore {
		return false, errors.New("score is too large")
	}

	return true, nil
}

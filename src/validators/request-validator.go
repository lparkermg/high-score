package request_validators

import (
	"errors"

	"mrlparker/high-score/models"

	"github.com/spf13/viper"
)

// Validates gameId and apiKey.
func ValidateGameAndApi(gameId string, apiKey string) (bool, error) {
	games := getConfig()

	if games[0].GameId != gameId || games[0].ApiKey != apiKey {
		return false, errors.New("game id or api key are not valid")
	}
	return true, nil
}

// Vaslidates score against gameId and the score.
func ValidateGameAndScore(gameId string, name string, score int) (bool, error) {
	games := getConfig()

	if len(name) > games[0].MaxNameLength {
		return false, errors.New("name is too long")
	}

	if score > games[0].MaxScore {
		return false, errors.New("score is too large")
	}

	return true, nil
}

func getConfig() []models.Game {
	var games []models.Game
	viper.UnmarshalKey("games", &games)
	return games
}

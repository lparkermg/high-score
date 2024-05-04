package request_validators

import (
	"errors"
	"fmt"

	"mrlparker/high-score/models"

	"github.com/spf13/viper"
)

// Validates gameId and apiKey.
func ValidateGameAndApi(gameId string, apiKey string) (bool, error) {
	game, err := getConfig(gameId)

	if err != nil {
		return false, err
	}

	if game.ApiKey != apiKey {
		return false, errors.New("game id or api key are not valid")
	}
	return true, nil
}

// Vaslidates score against gameId and the score.
func ValidateGameAndScore(gameId string, name string, score int) (bool, error) {
	game, err := getConfig(gameId)

	if err != nil {
		return false, err
	}

	if len(name) > game.MaxNameLength {
		return false, errors.New("name is too long")
	}

	if score > game.MaxScore {
		return false, errors.New("score is too large")
	}

	return true, nil
}

func getConfig(gameId string) (*models.Game, error) {
	var game *models.Game
	viper.UnmarshalKey(fmt.Sprintf("games.%v", gameId), &game)

	if game == nil {
		return nil, errors.New("gameid is not valid")
	}

	return game, nil
}

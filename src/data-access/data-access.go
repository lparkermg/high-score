package data_access

import models "mrlparker/high-score/models"

func GetScores(gameId string, skip int, take int) []models.Score {
	return []models.Score{
		{GameId: gameId, Name: "Testing", Score: 100},
		{GameId: gameId, Name: "ADB", Score: 99},
	}
}

func PostScore(gameId string, name string, score int) {

}

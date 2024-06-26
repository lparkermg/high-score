package data_access

import (
	"database/sql"
	"fmt"
	models "mrlparker/high-score/models"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var scoreTableName string
var gameTableName string

func GetScores(gameId string, skip int, take int) ([]models.Score, error) {
	db, dbErr := buildConnection()
	if dbErr != nil {
		return nil, dbErr
	}

	var scores []models.Score

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE GameId = ? ORDER BY Score DESC LIMIT %v, %v", scoreTableName, skip, take), gameId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var score models.Score
		if err := rows.Scan(&score.GameId, &score.Name, &score.Score); err != nil {
			return nil, fmt.Errorf("getScores %q: %v", gameId, err)
		}
		scores = append(scores, score)
	}

	return scores, nil
}

func PostScore(gameId string, name string, score int) (int64, error) {
	db, dbErr := buildConnection()
	if dbErr != nil {
		return 0, dbErr
	}

	result, err := db.Exec(fmt.Sprintf("INSERT INTO %s (GameId, Name, Score) VALUES (?, ?, ?)", scoreTableName), gameId, name, score)

	if err != nil {
		return 0, fmt.Errorf("PostScore: %v", err)
	}

	id, resErr := result.LastInsertId()
	if resErr != nil {
		return 0, fmt.Errorf("PostScore: %v", err)
	}

	return id, nil
}

func GetGame(gameId string) (models.Game, error) {
	var game models.Game
	db, dbErr := buildConnection()

	if dbErr != nil {
		return game, dbErr
	}

	row := db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE GameId = ?", gameTableName), gameId)

	if err := row.Scan(&game.GameId, &game.ApiKey, &game.MaxNameLength, &game.MaxScore); err != nil {
		return game, fmt.Errorf("getGame %q: %v", gameId, err)
	}

	return game, nil
}

func buildConnection() (*sql.DB, error) {
	scoreTableName = viper.GetString("DB_SCORETABLE_NAME")
	gameTableName = viper.GetString("DB_GAMETABLE_NAME")
	cfg := mysql.Config{
		User:                 viper.GetString("DB_USER"),
		Passwd:               viper.GetString("DB_PASS"),
		Net:                  "tcp",
		Addr:                 viper.GetString("DB_ADDRESS"),
		DBName:               viper.GetString("DB_NAME"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, err
	}

	return db, nil
}

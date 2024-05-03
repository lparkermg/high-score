package data_access

import (
	"database/sql"
	"fmt"
	models "mrlparker/high-score/models"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var dbName string

func GetScores(gameId string, skip int, take int) ([]models.Score, error) {
	db, dbErr := buildConnection()
	if dbErr != nil {
		return nil, dbErr
	}

	var scores []models.Score

	// TODO: Order by score descending, include skip + take
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE GameId = ?", dbName), gameId)

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

	result, err := db.Exec(fmt.Sprintf("INSERT INTO %s (GameId, Name, Score) VALUES (?, ?, ?)", gameId, name, score))

	if err != nil {
		return 0, fmt.Errorf("PostScore: %v", err)
	}

	id, resErr := result.LastInsertId()
	if resErr != nil {
		return 0, fmt.Errorf("PostScore: %v", err)
	}

	return id, nil
}

func buildConnection() (*sql.DB, error) {
	dbName = viper.GetString("DB_NAME")
	cfg := mysql.Config{
		User:   viper.GetString("DB_USER"),
		Passwd: viper.GetString("DB_PASS"),
		Net:    "tcp",
		Addr:   viper.GetString("DB_ADDRESS"),
		DBName: dbName,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, err
	}

	return db, nil
}

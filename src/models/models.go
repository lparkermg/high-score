package models

type Score struct {
	GameId string `json:"gameId"`
	Name   string `json:"name"`
	Score  int    `json:"score"`
}

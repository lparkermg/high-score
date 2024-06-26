package models

type Score struct {
	GameId string `json:"gameId"`
	Name   string `json:"name"`
	Score  int    `json:"score"`
}

type Game struct {
	GameId        string `json:"gameId"`
	ApiKey        string `json:"apiKey"`
	MaxNameLength int    `json:"maxNameLength"`
	MaxScore      int    `json:"maxScore"`
}

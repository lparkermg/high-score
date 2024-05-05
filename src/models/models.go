package models

type Score struct {
	GameId string `json:"gameId"`
	Name   string `json:"name"`
	Score  int    `json:"score"`
}

type Game struct {
	Id            string `json:"id"`
	ApiKey        string `json:"apiKey"`
	MaxNameLength int    `json:"maxNameLength"`
	MaxScore      int    `json:"maxScore"`
}

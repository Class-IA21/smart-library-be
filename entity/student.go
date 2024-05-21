package entity

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	NPM    string `json:"npm"`
	CardID int    `json:"card_id"`
}

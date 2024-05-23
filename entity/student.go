package entity

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name" validate:"required"`
	NPM    string `json:"npm" validate:"required"`
	CardID int    `json:"card_id" validate:"required"`
}

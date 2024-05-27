package entity

type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"required"`
	NPM       string `json:"npm" validate:"required"`
	AccountID int    `json:"account_id" validate:"required"`
	CardID    int    `json:"card_id" validate:"required"`
}

type StudentResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	NPM    string `json:"npm"`
	CardID int    `json:"card_id"`
}

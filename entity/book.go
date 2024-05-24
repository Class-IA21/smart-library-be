package entity

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title" validate:"required"`
	Author        string `json:"author" validate:"required"`
	Publisher     string `json:"publisher" validate:"required"`
	PublishedDate string `json:"published_date" validate:"required"`
	ISBN          string `json:"isbn" validate:"required"`
	Pages         int    `json:"pages" validate:"required,number"`
	Language      string `json:"language" validate:"required"`
	Genre         string `json:"genre" validate:"required"`
	Description   string `json:"description" validate:"required"`
	CardID        int    `json:"card_id" validate:"required,number"`
}

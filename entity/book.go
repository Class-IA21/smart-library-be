package entity

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	PublishedDate string `json:"published_date"`
	ISBN          string `json:"isbn"`
	Pages         int    `json:"pages"`
	Language      string `json:"language"`
	Genre         string `json:"genre"`
	Description   string `json:"description"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	CardID        int    `json:"card_id"`
}

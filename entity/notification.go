package entity

type Notification struct {
	TransactionID string `json:"transaction_id"`
	BookIDS       []int  `json:"book_ids"`
	DueDate       string `json:"due_date"`
	Message       string `json:"message"`
}

package entity

type Borrow struct {
	TransactionID string `json:"transaction_id"`
	BookIDS       []int  `json:"book_ids" validate:"required"`
	StudentID     int    `json:"student_id" validate:"required"`
	BorrowDate    string `json:"borrow_date"`
	DueDate       string `json:"due_date"`
	ReturnDate    string `json:"return_date"`
}

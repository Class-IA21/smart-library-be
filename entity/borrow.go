package entity

type Borrow struct {
	TransactionID string `json:"transaction_id"`
	BookIDS       []int  `json:"book_ids,omitempty" validate:"required"`
	BookID        int    `json:"book_id,omitempty"`
	StudentID     int    `json:"student_id" validate:"required"`
	BorrowDate    string `json:"borrow_date"`
	DueDate       string `json:"due_date"`
	ReturnDate    string `json:"return_date"`
	Status        string `json:"status"`
}

type BorrowUpdate struct {
	TransactionID string `json:"transaction_id"`
	BookID        int    `json:"book_id" validate:"required"`
	StudentID     int    `json:"student_id" validate:"required"`
	ReturnDate    string `json:"return_date" validate:"required_if=Status returned"`
	Status        string `json:"status" validate:"required,oneof=pending borrowed returned"`
}

type BorrowList struct {
	StudentID    int           `json:"student_id"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID         string `json:"id"`
	BookIDS    []int  `json:"book_ids" validate:"required"`
	BorrowDate string `json:"borrow_date"`
	DueDate    string `json:"due_date"`
	ReturnDate string `json:"return_date"`
}

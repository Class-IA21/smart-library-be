package repository

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
)

type BorrowRepositoryInterface interface {
	GetBorrowsByStudentID(ctx context.Context, db *sql.DB, studentID int) ([]string, *entity.ErrorResponse)
	GetBorrowByID(ctx context.Context, db *sql.DB, transactionID string) (*entity.Borrow, *entity.ErrorResponse)
	GetBorrows(ctx context.Context, db *sql.DB) ([]*entity.Borrow, *entity.ErrorResponse)
	GetBorrowByBookID(ctx context.Context, db *sql.DB, bookID int) ([]*entity.Borrow, *entity.ErrorResponse)
	InsertBorrow(ctx context.Context, tx *sql.Tx, borrow *entity.Borrow) *entity.ErrorResponse
	UpdateBorrow(ctx context.Context, tx *sql.Tx, borrow *entity.BorrowUpdate) *entity.ErrorResponse
}

type BorrowRepository struct{}

func NewBorrowRepository() *BorrowRepository {
	return &BorrowRepository{}
}

func (*BorrowRepository) GetBorrowsByStudentID(ctx context.Context, db *sql.DB, studentID int) (*entity.BorrowList, *entity.ErrorResponse) {

	rows, err := db.QueryContext(ctx, "SELECT transaction_id, student_id, book_id, borrow_date, due_date, return_date, status  FROM borrows WHERE student_id = ?", studentID)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer rows.Close()

	var borrowList entity.BorrowList
	var transactions []entity.Transaction
	var transactionMap = make(map[string]*entity.Transaction)

	for rows.Next() {
		var newTransactionID string
		var bookID int
		var borrowDate string
		var dueDate string
		var returnDate sql.NullString
		var status string
		err := rows.Scan(&newTransactionID, &studentID, &bookID, &borrowDate, &dueDate, &returnDate, &status)
		if err != nil {
			return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan borrow")
		}

		if existingTransaction, ok := transactionMap[newTransactionID]; ok {
			existingTransaction.BookIDS = append(existingTransaction.BookIDS, bookID)
		} else {
			transaction := entity.Transaction{
				ID:         newTransactionID,
				BorrowDate: borrowDate,
				DueDate:    dueDate,
				ReturnDate: returnDate.String,
				BookIDS:    []int{bookID},
				Status:     status,
			}
			transactionMap[newTransactionID] = &transaction
		}
	}

	for _, value := range transactionMap {
		transactions = append(transactions, *value)
	}

	borrowList.Transactions = transactions

	return &borrowList, nil
}

func (*BorrowRepository) GetBorrowByID(ctx context.Context, db *sql.DB, trxID string) (*entity.Borrow, *entity.ErrorResponse) {
	rows, err := db.QueryContext(ctx, "SELECT book_id, student_id, transaction_id, borrow_date, due_date, return_date FROM borrows WHERE transaction_id = ?", trxID)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer rows.Close()

	var borrow entity.Borrow
	var bookIds []int
	var studentId int
	var borrowDate, dueDate string
	var returnDate, transactionID sql.NullString
	for rows.Next() {
		var bookId int
		err := rows.Scan(&bookId, &studentId, &transactionID, &borrowDate, &dueDate, &returnDate)
		if err != nil {
			return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan borrow")
		}
		bookIds = append(bookIds, bookId)
	}

	if !transactionID.Valid {
		return nil, helper.ErrorResponse(http.StatusNotFound, "Transaction ID not found")
	}

	if returnDate.Valid {
		borrow.ReturnDate = returnDate.String
	}

	borrow.TransactionID = transactionID.String
	borrow.StudentID = studentId
	borrow.BookIDS = bookIds
	borrow.BorrowDate = borrowDate
	borrow.DueDate = dueDate

	return &borrow, nil
}

func (*BorrowRepository) GetBorrowByBookID(ctx context.Context, db *sql.DB, bookID int) ([]*entity.Borrow, *entity.ErrorResponse) {
	rows, err := db.QueryContext(ctx, "SELECT transaction_id, borrow_date, due_date, return_date, status, student_id FROM borrows WHERE book_id = ?", bookID)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer rows.Close()

	var borrows []*entity.Borrow
	for rows.Next() {
		var borrow entity.Borrow
		var returnDate sql.NullString
		err := rows.Scan(&borrow.TransactionID, &borrow.BorrowDate, &borrow.DueDate, &returnDate, &borrow.Status, &borrow.StudentID)
		if err != nil {
			return nil, helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
		}
		borrow.ReturnDate = returnDate.String
		borrows = append(borrows, &borrow)
	}

	return borrows, nil
}

func (*BorrowRepository) GetBorrows(ctx context.Context, db *sql.DB) (*entity.BorrowList, *entity.ErrorResponse) {
	rows, err := db.QueryContext(ctx, "SELECT transaction_id, student_id, book_id, borrow_date, due_date, return_date, status  FROM borrows")
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer rows.Close()

	var borrowList entity.BorrowList
	var transactions []entity.Transaction
	var transactionMap = make(map[string]*entity.Transaction)
	var studentID int

	for rows.Next() {
		var newTransactionID string
		var bookID int
		var borrowDate string
		var dueDate string
		var returnDate sql.NullString
		var status string
		err := rows.Scan(&newTransactionID, &studentID, &bookID, &borrowDate, &dueDate, &returnDate, &status)
		if err != nil {
			return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan borrow")
		}

		if existingTransaction, ok := transactionMap[newTransactionID]; ok {
			existingTransaction.BookIDS = append(existingTransaction.BookIDS, bookID)
		} else {
			transaction := entity.Transaction{
				ID:         newTransactionID,
				BorrowDate: borrowDate,
				DueDate:    dueDate,
				ReturnDate: returnDate.String,
				BookIDS:    []int{bookID},
				Status:     status,
			}
			transactionMap[newTransactionID] = &transaction
		}
	}

	for _, value := range transactionMap {
		transactions = append(transactions, *value)
	}

	borrowList.Transactions = transactions

	return &borrowList, nil
}

func (*BorrowRepository) InsertBorrow(ctx context.Context, tx *sql.Tx, borrow *entity.Borrow) *entity.ErrorResponse {

	borrowDate := time.Now()
	dueDate := time.Now().AddDate(0, 0, 7)

	for _, bookID := range borrow.BookIDS {
		_, err := tx.ExecContext(ctx, "INSERT INTO borrows (student_id, transaction_id, book_id, borrow_date, due_date) VALUES (?, ?, ?, ?, ?)",
			borrow.StudentID,
			borrow.TransactionID,
			bookID,
			borrowDate,
			dueDate,
		)
		if err != nil {
			return helper.ErrorResponse(http.StatusInternalServerError, "failed to insert borrow")
		}
	}
	return nil
}

func (*BorrowRepository) UpdateBorrow(ctx context.Context, tx *sql.Tx, borrow *entity.BorrowUpdate) *entity.ErrorResponse {
	query := "UPDATE borrows SET book_id = ?, status = ?"
	var params []interface{}
	params = append(params, borrow.BookID, borrow.Status)

	if borrow.ReturnDate != "" {
		query += ", return_date = ?"
		params = append(params, borrow.ReturnDate)
	}

	query += " WHERE transaction_id = ? AND book_id = ?"
	params = append(params, borrow.TransactionID, borrow.BookID)

	_, err := tx.ExecContext(ctx, query, params...)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return nil
}

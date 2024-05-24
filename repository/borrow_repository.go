package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
)

type BorrowRepositoryInterface interface {
	GetTransactionsByStudentID(ctx context.Context, db *sql.DB, studentID int) ([]string, *entity.ErrorResponse)
	GetBorrowByTransactionID(ctx context.Context, db *sql.DB, transactionID string) (*entity.Borrow, *entity.ErrorResponse)
	InsertBorrow(ctx context.Context, tx *sql.Tx, borrow *entity.Borrow) *entity.ErrorResponse
}

type BorrowRepository struct{}

func NewBorrowRepository() *BorrowRepository {
	return &BorrowRepository{}
}

func (*BorrowRepository) GetTransactionsByStudentID(ctx context.Context, db *sql.DB, studentID int) (*entity.BorrowList, *entity.ErrorResponse) {

	rows, err := db.QueryContext(ctx, "SELECT transaction_id, student_id, book_id, borrow_date, due_date, return_date  FROM borrows WHERE student_id = ?", studentID)
	if err != nil {
		fmt.Println(err)
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
		err := rows.Scan(&newTransactionID, &studentID, &bookID, &borrowDate, &dueDate, &returnDate)
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
			}
			transactionMap[newTransactionID] = &transaction
		}
	}

	for _, value := range transactionMap {
		transactions = append(transactions, *value)
	}

	borrowList.StudentID = studentID
	borrowList.Transactions = transactions

	return &borrowList, nil
}

func (*BorrowRepository) GetBorrowByTransactionID(ctx context.Context, db *sql.DB, trxID string) (*entity.Borrow, *entity.ErrorResponse) {
	rows, err := db.QueryContext(ctx, "SELECT book_id, student_id, transaction_id, borrow_date, due_date, return_date FROM borrows WHERE transaction_id = ?", trxID)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

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

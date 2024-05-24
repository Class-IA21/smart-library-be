package services

import (
	"context"
	"database/sql"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/repository"
	"github.com/google/uuid"
)

type BorrowServiceInterface interface {
	GetBorrowByTransactionID(ctx context.Context, transactionID string) (*entity.Borrow, *entity.ErrorResponse)
	InsertBorrow(ctx context.Context, borrow *entity.Borrow) *entity.ErrorResponse
}

type BorrowServices struct {
	DB *sql.DB
	*repository.BorrowRepository
	*StudentServices
	*BookServices
}

func NewBorrowServices(db *sql.DB, borrowRepo *repository.BorrowRepository, studentService *StudentServices, bookService *BookServices) *BorrowServices {
	return &BorrowServices{DB: db, BorrowRepository: borrowRepo, StudentServices: studentService, BookServices: bookService}
}

func (s *BorrowServices) GetBorrowByTransactionID(ctx context.Context, transactionID string) (*entity.Borrow, *entity.ErrorResponse) {
	return s.BorrowRepository.GetBorrowByTransactionID(ctx, s.DB, transactionID)
}

func (s *BorrowServices) InsertBorrow(ctx context.Context, borrow *entity.Borrow) *entity.ErrorResponse {

	for _, bookID := range borrow.BookIDS {
		_, err := s.BookServices.GetBookByID(ctx, bookID)
		if err != nil {
			return err
		}
	}

	_, errorResponse := s.StudentServices.GetStudentByID(ctx, borrow.StudentID)
	if errorResponse != nil {
		return errorResponse
	}

	uuid := uuid.New().String()
	borrow.TransactionID = uuid

	tx, err := s.DB.Begin()
	if err != nil {
		tx.Rollback()
		return nil
	}

	errorResponse = s.BorrowRepository.InsertBorrow(ctx, tx, borrow)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	tx.Commit()
	return nil
}

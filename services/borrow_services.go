package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/repository"
	"github.com/google/uuid"
)

type BorrowServiceInterface interface {
	GetBorrowsByStudentID(ctx context.Context, studentId int) (transactionID []string, error *entity.ErrorResponse)
	GetBorrowByTransactionID(ctx context.Context, transactionID string) (*entity.Borrow, *entity.ErrorResponse)
	GetBorrowsByBookID(ctx context.Context, bookID int) ([]*entity.Borrow, *entity.ErrorResponse)
	GetBorrows(ctx context.Context) ([]*entity.Borrow, *entity.ErrorResponse)
	InsertBorrow(ctx context.Context, borrow *entity.Borrow) *entity.ErrorResponse
	UpdateBorrow(ctx context.Context, borrow *entity.BorrowUpdate) *entity.ErrorResponse
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

func (s *BorrowServices) GetBorrowsByStudentID(ctx context.Context, studentId int) (*entity.BorrowList, *entity.ErrorResponse) {

	return s.BorrowRepository.GetBorrowsByStudentID(ctx, s.DB, studentId)
}

func (s *BorrowServices) GetBorrowByTransactionID(ctx context.Context, transactionID string) (*entity.Borrow, *entity.ErrorResponse) {
	return s.BorrowRepository.GetBorrowByID(ctx, s.DB, transactionID)
}

func (s *BorrowServices) GetBorrowByBookID(ctx context.Context, bookID int) ([]*entity.Borrow, *entity.ErrorResponse) {
	return s.BorrowRepository.GetBorrowByBookID(ctx, s.DB, bookID)
}

func (s *BorrowServices) GetBorrows(ctx context.Context) ([]*entity.Borrow, *entity.ErrorResponse) {
	return s.BorrowRepository.GetBorrows(ctx, s.DB)
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
		return nil
	}
	defer tx.Commit()

	errorResponse = s.BorrowRepository.InsertBorrow(ctx, tx, borrow)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	return nil
}

func (s *BorrowServices) UpdateBorrow(ctx context.Context, borrow *entity.BorrowUpdate) *entity.ErrorResponse {
	_, errorResponse := s.GetBorrowByTransactionID(ctx, borrow.TransactionID)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer tx.Commit()

	errorResponse = s.BorrowRepository.UpdateBorrow(ctx, tx, borrow)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	return nil
}

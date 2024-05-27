package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/repository"
)

type BookServicesInterface interface {
	GetBooks(ctx context.Context, page, pageSize int) ([]*entity.Book, *entity.ErrorResponse)
	GetBookByID(ctx context.Context, bookID int) (*entity.Book, *entity.ErrorResponse)
	GetBookByCardID(ctx context.Context, cardID int) (*entity.Book, *entity.ErrorResponse)
	DeleteBookByID(ctx context.Context, bookID int) *entity.ErrorResponse
	UpdateBook(ctx context.Context, book *entity.Book) *entity.ErrorResponse
	InsertBook(ctx context.Context, book *entity.Book) *entity.ErrorResponse
}

type BookServices struct {
	*repository.BookRepository
	*sql.DB
}

func NewBookServices(br *repository.BookRepository, db *sql.DB) *BookServices {
	return &BookServices{
		BookRepository: br,
		DB:             db,
	}
}

func (s *BookServices) GetBooks(ctx context.Context, page, pageSize int) ([]*entity.Book, *entity.ErrorResponse) {
	return s.BookRepository.GetBooks(ctx, s.DB, page, pageSize)
}

func (s *BookServices) GetBookByID(ctx context.Context, bookID int) (*entity.Book, *entity.ErrorResponse) {
	if bookID <= 0 {
		return nil, helper.ErrorResponse(http.StatusBadRequest, "invalid book id")
	}
	return s.BookRepository.GetBookByID(ctx, s.DB, bookID)
}

func (s *BookServices) GetBookByCardID(ctx context.Context, cardID int) (*entity.Book, *entity.ErrorResponse) {
	if cardID <= 0 {
		return nil, helper.ErrorResponse(http.StatusBadRequest, "invalid card id")
	}

	return s.BookRepository.GetBookByCardID(ctx, s.DB, cardID)
}

func (s *BookServices) DeleteBookByID(ctx context.Context, bookID int) *entity.ErrorResponse {
	_, errorResponse := s.BookRepository.GetBookByID(ctx, s.DB, bookID)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer tx.Commit()

	if bookID <= 0 {
		return helper.ErrorResponse(http.StatusBadRequest, "invalid book ID")
	}

	errorResponse = s.BookRepository.DeleteBookByID(ctx, tx, bookID)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	return nil
}

func (s *BookServices) UpdateBook(ctx context.Context, book *entity.Book) *entity.ErrorResponse {
	_, errorResponse := s.GetBookByID(ctx, book.ID)
	fmt.Println(errorResponse)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, err.Error())
	}
	defer tx.Commit()

	if book.ID <= 0 {
		return helper.ErrorResponse(http.StatusBadRequest, "invalid book ID")
	}

	if book == nil {
		return helper.ErrorResponse(http.StatusBadRequest, "Payload can't be null")
	}

	errorResponse = s.BookRepository.UpdateBook(ctx, tx, book)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	return nil
}

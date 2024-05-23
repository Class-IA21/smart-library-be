package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/repository"
)

type BookServiceInterface interface {
	GetBooks(ctx context.Context, page, pageSize int) ([]*entity.Book, *entity.ErrorResponse)
	GetBookByID(ctx context.Context, bookID int) (*entity.Book, *entity.ErrorResponse)
	GetBookByCardID(ctx context.Context, cardID int) (*entity.Book, *entity.ErrorResponse)
	DeleteBookByID(ctx context.Context, bookID int) *entity.ErrorResponse
	UpdateBookByID(ctx context.Context, book *entity.Book) *entity.ErrorResponse
	InsertBook(ctx context.Context, book *entity.Book) *entity.ErrorResponse
}

type BookService struct {
	repo *repository.BookRepository
	db   *sql.DB
}

func NewBookService(repo *repository.BookRepository, db *sql.DB) *BookService {
	return &BookService{
		repo: repo,
		db:   db,
	}
}

func (s *BookService) GetBooks(ctx context.Context, page, pageSize int) ([]*entity.Book, *entity.ErrorResponse) {
	return s.repo.GetBooks(ctx, s.db, page, pageSize)
}

func (s *BookService) GetBookByID(ctx context.Context, bookID int) (*entity.Book, *entity.ErrorResponse) {
	if bookID <= 0 {
		return nil, helper.ErrorResponse(http.StatusBadRequest, "invalid book id")
	}
	return s.repo.GetBookByID(ctx, s.db, bookID)
}

func (s *BookService) GetBookByCardID(ctx context.Context, cardID int) (*entity.Book, *entity.ErrorResponse) {
	if cardID <= 0 {
		return nil, helper.ErrorResponse(http.StatusBadRequest, "invalid card id")
	}

	return s.repo.GetBookByCardID(ctx, s.db, cardID)
}

func (s *BookService) DeleteBookByID(ctx context.Context, bookID int) *entity.ErrorResponse {

	_, errorResponse := s.repo.GetBookByID(ctx, s.db, bookID)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.db.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, err.Error())
	}

	if bookID <= 0 {
		return helper.ErrorResponse(http.StatusBadRequest, "invalid book ID")
	}

	errorResponse = s.repo.DeleteBookByID(ctx, tx, bookID)
	if errorResponse != nil {
		tx.Rollback()
		return helper.ErrorResponse(http.StatusInternalServerError, err.Error())
	}

	tx.Commit()
	return nil
}

func (s *BookService) UpdateBookByID(ctx context.Context, book *entity.Book) *entity.ErrorResponse {
	_, errorResponse := s.repo.GetBookByID(ctx, s.db, book.ID)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.db.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, err.Error())
	}

	if book.ID <= 0 {
		return helper.ErrorResponse(http.StatusBadRequest, "invalid book ID")
	}

	if book == nil {
		return helper.ErrorResponse(http.StatusBadRequest, "Payload can't be null")
	}

	errorResponse = s.repo.UpdateBookByID(ctx, tx, book)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	tx.Commit()
	return nil
}

func (s *BookService) InsertBook(ctx context.Context, book *entity.Book) *entity.ErrorResponse {

	if book == nil {
		return helper.ErrorResponse(http.StatusBadRequest, "Payload can't be null")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, err.Error())
	}

	errorResponse := s.repo.InsertBook(ctx, tx, book)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	tx.Commit()
	return nil
}

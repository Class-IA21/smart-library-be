package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
)

type BookCardServicesInterface interface {
	UpdateBook(ctx context.Context, book *entity.Book) *entity.ErrorResponse
	InsertBook(ctx context.Context, book *entity.Book) *entity.ErrorResponse
	DeleteBook(ctx context.Context, bookID int) *entity.ErrorResponse
}

type BookCardServices struct {
	*sql.DB
	*BookServices
	*CardServices
}

func NewBookCardServices(DB *sql.DB, bs *BookServices, cs *CardServices) *BookCardServices {
	return &BookCardServices{
		DB:           DB,
		BookServices: bs,
		CardServices: cs,
	}
}

func (s *BookCardServices) UpdateBook(ctx context.Context, book *entity.Book) *entity.ErrorResponse {
	_, err := s.CardServices.GetCardByID(ctx, book.CardID)
	if err != nil {
		return err
	}

	return s.BookServices.UpdateBook(ctx, book)
}

func (s *BookCardServices) InsertBook(ctx context.Context, book *entity.Book) (int, *entity.ErrorResponse) {
	if book == nil {
		return 0, helper.ErrorResponse(http.StatusBadRequest, "Payload can't be null")
	}

	_, errorResponse := s.CardServices.GetCardByID(ctx, book.CardID)
	if errorResponse != nil {
		return 0, errorResponse
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return 0, helper.ErrorResponse(http.StatusInternalServerError, err.Error())
	}

	id, errorResponse := s.BookRepository.InsertBook(ctx, tx, book)
	if errorResponse != nil {
		tx.Rollback()
		return 0, errorResponse
	}

	tx.Commit()
	return id, nil
}

func (s *BookCardServices) DeleteBook(ctx context.Context, bookID int) *entity.ErrorResponse {
	book, errorResponse := s.BookServices.GetBookByID(ctx, bookID)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, err.Error())
	}
	defer tx.Commit()

	errorResponse = s.BookServices.DeleteCardIDFromBook(ctx, book.ID)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	errorResponse = s.CardServices.DeleteCard(ctx, book.CardID)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	errorResponse = s.BookServices.DeleteBookByID(ctx, book.ID)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	return nil
}

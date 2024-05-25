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
	_, err := s.BookServices.GetBookByID(ctx, book.ID)
	if err != nil {
		return err
	}

	_, err = s.CardServices.GetCardByID(ctx, book.CardID)
	if err != nil {
		return err
	}

	return s.BookServices.UpdateBook(ctx, book)
}

func (s *BookCardServices) InsertBook(ctx context.Context, book *entity.Book) *entity.ErrorResponse {
	if book == nil {
		return helper.ErrorResponse(http.StatusBadRequest, "Payload can't be null")
	}

	_, errorResponse := s.CardServices.GetCardByID(ctx, book.CardID)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, err.Error())
	}

	errorResponse = s.BookRepository.InsertBook(ctx, tx, book)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	tx.Commit()
	return nil
}

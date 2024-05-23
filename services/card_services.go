package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/repository"
)

type CardServiceInterface interface {
	GetCards(ctx context.Context, page, pageSize int) ([]*entity.Card, *entity.ErrorResponse)
	GetCardTypeByUID(ctx context.Context, uid string) (data any, cardType string, err *entity.ErrorResponse)
	GetCardByUID(ctx context.Context, uid string) (*entity.Card, *entity.ErrorResponse)
	GetCardByID(ctx context.Context, id int) (*entity.Card, *entity.ErrorResponse)
}

type CardServices struct {
	db *sql.DB
	*repository.BookRepository
	*repository.CardRepository
	*repository.StudentRepository
}

func NewCardServices(db *sql.DB, bookRepository *repository.BookRepository, cardRepository *repository.CardRepository, studentRepository *repository.StudentRepository) *CardServices {
	return &CardServices{
		db:                db,
		BookRepository:    bookRepository,
		CardRepository:    cardRepository,
		StudentRepository: studentRepository,
	}
}

func (s *CardServices) GetCards(ctx context.Context, page, pageSize int) ([]*entity.Card, *entity.ErrorResponse) {
	return s.CardRepository.GetCards(ctx, s.db, page, pageSize)
}

func (s *CardServices) GetCardTypeByUID(ctx context.Context, uid string) (data any, cardType string, err *entity.ErrorResponse) {
	result, err := s.CardRepository.GetCardByUID(ctx, s.db, uid)
	if err != nil {
		return nil, "", err
	}

	switch result.Type {
	case "student":
		result, err := s.StudentRepository.GetStudentByCardID(ctx, s.db, result.ID)
		if err != nil {
			return nil, "", err
		}
		return result, "student", nil
	case "book":
		result, err := s.BookRepository.GetBookByCardID(ctx, s.db, result.ID)
		if err != nil {
			return nil, "", err
		}
		return result, "book", nil
	default:
		return nil, "", helper.ErrorResponse(http.StatusNotFound, "rfid not registered")
	}
}

func (s *CardServices) GetCardByUID(ctx context.Context, uid string) (*entity.Card, *entity.ErrorResponse) {
	result, err := s.CardRepository.GetCardByUID(ctx, s.db, uid)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *CardServices) GetCardByID(ctx context.Context, id int) (*entity.Card, *entity.ErrorResponse) {
	result, err := s.CardRepository.GetCardByID(ctx, s.db, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

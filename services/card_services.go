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
	GetCardTypeByUID(ctx context.Context, uid string) (id int, cardType string, err *entity.ErrorResponse)
	GetCardByUID(ctx context.Context, uid string) (*entity.Card, *entity.ErrorResponse)
	GetCardByID(ctx context.Context, id int) (*entity.Card, *entity.ErrorResponse)
	InsertCard(ctx context.Context, card *entity.Card) *entity.ErrorResponse
	UpdateCard(ctx context.Context, id int, card *entity.Card) *entity.ErrorResponse
	DeleteCard(ctx context.Context, id int) *entity.ErrorResponse
}

type CardServices struct {
	db *sql.DB
	*repository.CardRepository
	*BookServices
	*StudentServices
}

func NewCardServices(db *sql.DB, cardRepository *repository.CardRepository, studentService *StudentServices, bookService *BookServices) *CardServices {
	return &CardServices{
		db:              db,
		CardRepository:  cardRepository,
		BookServices:    bookService,
		StudentServices: studentService,
	}
}

func (s *CardServices) GetCards(ctx context.Context, page, pageSize int) ([]*entity.Card, *entity.ErrorResponse) {
	return s.CardRepository.GetCards(ctx, s.db, page, pageSize)
}

func (s *CardServices) GetCardTypeByUID(ctx context.Context, uid string) (id int, cardType string, err *entity.ErrorResponse) {
	result, err := s.CardRepository.GetCardByUID(ctx, s.db, uid)
	if err != nil {
		return 0, "", err
	}

	switch result.Type {
	case "student":
		result, err := s.StudentServices.GetStudentByCardID(ctx, result.ID)
		if err != nil {
			return 0, "", err
		}
		return result.ID, "student", nil
	case "book":
		result, err := s.BookServices.GetBookByCardID(ctx, result.ID)
		if err != nil {
			return 0, "", err
		}
		return result.ID, "book", nil
	default:
		return 0, "", helper.ErrorResponse(http.StatusNotFound, "rfid not registered")
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

func (s *CardServices) InsertCard(ctx context.Context, card *entity.Card) *entity.ErrorResponse {
	existingCard, _ := s.GetCardByUID(ctx, card.UID)
	if existingCard != nil {
		return helper.ErrorResponse(http.StatusConflict, "UID already registered.")
	}

	return s.CardRepository.InsertCard(ctx, s.db, card)
}

func (s *CardServices) UpdateCard(ctx context.Context, id int, card *entity.Card) *entity.ErrorResponse {
	existingCard, err := s.CardRepository.GetCardByID(ctx, s.db, id)
	if err != nil {
		return err
	}

	if existingCard == nil {
		return helper.ErrorResponse(http.StatusNotFound, "Card not found")
	}

	err = s.CardRepository.UpdateCard(ctx, s.db, card)
	if err != nil {
		return err
	}

	return nil
}

func (s *CardServices) DeleteCard(ctx context.Context, id int) *entity.ErrorResponse {
	existingCard, err := s.CardRepository.GetCardByID(ctx, s.db, id)
	if err != nil {
		return err
	}

	if existingCard == nil {
		return helper.ErrorResponse(http.StatusNotFound, "Card not found")
	}

	err = s.CardRepository.DeleteCard(ctx, s.db, id)
	if err != nil {
		return err
	}

	return nil
}

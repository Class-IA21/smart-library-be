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
	InsertContainerCard(ctx context.Context, uid string) *entity.ErrorResponse
	GetOnceContainerCard(ctx context.Context) (string, *entity.ErrorResponse)
}

type CardServices struct {
	*sql.DB
	*repository.CardRepository
	*BookServices
	*StudentServices
}

func NewCardServices(db *sql.DB, cardRepository *repository.CardRepository, studentService *StudentServices, bookService *BookServices) *CardServices {
	return &CardServices{
		DB:              db,
		CardRepository:  cardRepository,
		BookServices:    bookService,
		StudentServices: studentService,
	}
}

func (s *CardServices) GetCards(ctx context.Context, page, pageSize int) ([]*entity.Card, *entity.ErrorResponse) {
	return s.CardRepository.GetCards(ctx, s.DB, page, pageSize)
}

func (s *CardServices) GetCardTypeByUID(ctx context.Context, uid string) (id int, cardType string, err *entity.ErrorResponse) {
	result, err := s.CardRepository.GetCardByUID(ctx, s.DB, uid)
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
	result, err := s.CardRepository.GetCardByUID(ctx, s.DB, uid)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *CardServices) GetCardByID(ctx context.Context, id int) (*entity.Card, *entity.ErrorResponse) {
	result, err := s.CardRepository.GetCardByID(ctx, s.DB, id)
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

	return s.CardRepository.InsertCard(ctx, s.DB, card)
}

func (s *CardServices) UpdateCard(ctx context.Context, id int, card *entity.Card) *entity.ErrorResponse {
	existingCard, err := s.CardRepository.GetCardByID(ctx, s.DB, id)
	if err != nil {
		return err
	}

	if existingCard == nil {
		return helper.ErrorResponse(http.StatusNotFound, "Card not found")
	}

	err = s.CardRepository.UpdateCard(ctx, s.DB, card)
	if err != nil {
		return err
	}

	return nil
}

func (s *CardServices) DeleteCard(ctx context.Context, id int) *entity.ErrorResponse {
	existingCard, err := s.CardRepository.GetCardByID(ctx, s.DB, id)
	if err != nil {
		return err
	}

	if existingCard == nil {
		return helper.ErrorResponse(http.StatusNotFound, "Card not found")
	}

	book, _ := s.BookServices.GetBookByCardID(ctx, id)
	if book != nil {
		s.BookServices.DeleteCardIDFromBook(ctx, id)
	}

	student, _ := s.StudentServices.GetStudentByCardID(ctx, id)
	if student != nil {
		s.StudentServices.DeleteCardIDFromStudent(ctx, id)
	}

	err = s.CardRepository.DeleteCard(ctx, s.DB, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *CardServices) InsertContainerCard(ctx context.Context, uid string) *entity.ErrorResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer tx.Commit()

	return s.CardRepository.InsertContainerCard(ctx, tx, uid)
}

func (s *CardServices) GetOnceContainerCard(ctx context.Context) (string, *entity.ErrorResponse) {
	tx, err := s.DB.Begin()
	if err != nil {
		return "", helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer tx.Commit()

	return s.CardRepository.GetOnceContainerCard(ctx, tx)
}

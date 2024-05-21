package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/repository"
)

type CardServiceInterface interface {
	GetCardTypeByUID(ctx context.Context, uid int) (data any, cardType string, err error)
	GetCardByUID(ctx context.Context, uid int) (*entity.Card, error)
	GetCardByID(ctx context.Context, id int) (*entity.Card, error)
}

type CardServices struct {
	db *sql.DB
	*repository.BookRepostitory
	*repository.CardRepository
	*repository.StudentRepository
}

func NewCardServices(db *sql.DB, bookRepository *repository.BookRepostitory, cardRepository *repository.CardRepository, studentRepository *repository.StudentRepository) *CardServices {
	return &CardServices{
		db:                db,
		BookRepostitory:   bookRepository,
		CardRepository:    cardRepository,
		StudentRepository: studentRepository,
	}
}

func (s *CardServices) GetCardTypeByUID(ctx context.Context, uid int) (data any, cardType string, err error) {
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
		result, err := s.BookRepostitory.GetBookByCardID(ctx, s.db, result.ID)
		if err != nil {
			return nil, "", err
		}
		return result, "book", nil
	default:
		return nil, "", errors.New("rfid not registered")
	}
}

func (s *CardServices) GetCardByUID(ctx context.Context, uid int) (*entity.Card, error) {
	result, err := s.CardRepository.GetCardByUID(ctx, s.db, uid)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *CardServices) GetCardByID(ctx context.Context, id int) (*entity.Card, error) {
	result, err := s.CardRepository.GetCardByID(ctx, s.db, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

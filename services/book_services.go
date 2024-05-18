package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/repository"
)

type BookServiceInterface interface {
	GetBookByID(ctx context.Context, bookID int) (*entity.Book, error)
	GetAllBooks(ctx context.Context) ([]*entity.Book, error)
	DeleteBookByID(ctx context.Context, bookID int) error
	UpdateBookByID(ctx context.Context, book *entity.Book) error
	GetBookByCardID(ctx context.Context, cardID int) (*entity.Book, error)
}

type BookService struct {
	repo repository.BookRepositoryInterface
	db   *sql.DB
}

func NewBookService(repo repository.BookRepositoryInterface, db *sql.DB) BookServiceInterface {
	return &BookService{
		repo: repo,
		db:   db,
	}
}

func (s *BookService) GetBookByID(ctx context.Context, bookID int) (*entity.Book, error) {
	if bookID <= 0 {
		return nil, errors.New("invalid book ID")
	}
	return s.repo.GetBookByID(ctx, s.db, bookID)
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]*entity.Book, error) {
	return s.repo.GetAllBooks(ctx, s.db)
}

func (s *BookService) DeleteBookByID(ctx context.Context, bookID int) error {
	if bookID <= 0 {
		return errors.New("invalid book ID")
	}
	return s.repo.DeleteBookByID(ctx, s.db, bookID)
}

func (s *BookService) UpdateBookByID(ctx context.Context, book *entity.Book) error {
	if book == nil {
		return errors.New("book cannot be nil")
	}
	if book.ID <= 0 {
		return errors.New("invalid book ID")
	}
	return s.repo.UpdateBookByID(ctx, s.db, book)
}

func (s *BookService) GetBookByCardID(ctx context.Context, cardID int) (*entity.Book, error) {
	if cardID <= 0 {
		return nil, errors.New("invalid card ID")
	}
	return s.repo.GetBookByCardID(ctx, s.db, cardID)
}

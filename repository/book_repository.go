package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/smart-library-be/entity"
)

type BookRepositoryInterface interface {
	GetBookByID(ctx context.Context, db *sql.DB, bookID int) (*entity.Book, error)
	GetAllBooks(ctx context.Context, db *sql.DB) ([]*entity.Book, error)
	DeleteBookByID(ctx context.Context, db *sql.DB, bookID int) error
	UpdateBookByID(ctx context.Context, db *sql.DB, book *entity.Book) error
	GetBookByCardID(ctx context.Context, db *sql.DB, cardID int) (*entity.Book, error)

type BookRepostitoryImpl struct{}

func NewBookRepositoryImpl() BookRepositoryInterface {
	return &BookRepostitoryImpl{}
}

func (*BookRepostitoryImpl) GetBookByID(ctx context.Context, db *sql.DB, bookID int) (*entity.Book, error) {
	result := db.QueryRowContext(ctx, "SELECT * FROM books WHERE id = ?", bookID)

	var book entity.Book
	err := result.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Publisher,
		&book.PublishedDate,
		&book.ISBN,
		&book.Pages,
		&book.Language,
		&book.Genre,
		&book.Description,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err != nil {
		return nil, errors.New("failed to fetch book")
	}
	return &book, nil
}

func (*BookRepostitoryImpl) GetAllBooks(ctx context.Context, db *sql.DB) ([]*entity.Book, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM books")
	if err != nil {
		return nil, errors.New("failed to fetch books")
	}
	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var book entity.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Publisher,
			&book.PublishedDate,
			&book.ISBN,
			&book.Pages,
			&book.Language,
			&book.Genre,
			&book.Description,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, errors.New("failed to scan books")
		}
		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("error reading books")
	}
	return books, nil
}

func (*BookRepostitoryImpl) DeleteBookByID(ctx context.Context, db *sql.DB, bookID int) error {
	_, err := db.ExecContext(ctx, "DELETE FROM books WHERE id = ?", bookID)
	if err != nil {
		return errors.New("failed to delete book")
	}
	return nil
}

func (*BookRepostitoryImpl) UpdateBookByID(ctx context.Context, db *sql.DB, book *entity.Book) error {
	_, err := db.ExecContext(ctx, "UPDATE books SET title=?, author=?, publisher=?, published_date=?, isbn=?, pages=?, language=?, genre=?, description=?, created_at=?, updated_at=? WHERE id=?",
		book.Title,
		book.Author,
		book.Publisher,
		book.PublishedDate,
		book.ISBN,
		book.Pages,
		book.Language,
		book.Genre,
		book.Description,
		book.CreatedAt,
		book.UpdatedAt,
		book.ID,
	)
	if err != nil {
		return errors.New("failed to update book")
	}
	return nil
}

func (*BookRepostitoryImpl) GetBookByCardID(ctx context.Context, db *sql.DB, cardID int) (*entity.Book, error) {
	result := db.QueryRowContext(ctx, "SELECT * FROM books WHERE card_id = ?", cardID)

	var book entity.Book
	err := result.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Publisher,
		&book.PublishedDate,
		&book.ISBN,
		&book.Pages,
		&book.Language,
		&book.Genre,
		&book.Description,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err != nil {
		return nil, errors.New("failed to fetch book")
	}
	return &book, nil
}

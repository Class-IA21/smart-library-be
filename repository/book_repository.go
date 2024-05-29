package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
)

type BookRepositoryInterface interface {
	GetBookByID(ctx context.Context, db *sql.DB, bookID int) (*entity.Book, *entity.ErrorResponse)
	GetAllBooks(ctx context.Context, db *sql.DB) ([]*entity.Book, *entity.ErrorResponse)
	GetBookByCardID(ctx context.Context, db *sql.DB, cardID int) (*entity.Book, *entity.ErrorResponse)
	DeleteBookByID(ctx context.Context, tx *sql.Tx, bookID int) *entity.ErrorResponse
	DeleteCardIDFromBook(ctx context.Context, tx *sql.Tx, cardID int) *entity.ErrorResponse
	UpdateBook(ctx context.Context, tx *sql.Tx, book *entity.Book) *entity.ErrorResponse
	InsertBook(ctx context.Context, tx *sql.Tx, book *entity.Book) *entity.ErrorResponse
}

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (*BookRepository) GetBooks(ctx context.Context, db *sql.DB, page, pageSize int) ([]*entity.Book, *entity.ErrorResponse) {

	var query string
	if page != 0 && pageSize != 0 {
		offset := (page - 1) * pageSize
		query = fmt.Sprintf("SELECT * FROM books LIMIT %d OFFSET %d", pageSize, offset)
	} else {
		query = "SELECT * FROM books"
	}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, helper.ErrorResponse(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var book entity.Book
		var cardID sql.NullInt64
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
			&cardID,
		)
		if err != nil {
			return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan book")
		}
		book.CardID = int(cardID.Int64)
		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to read books")
	}
	return books, nil
}

func (*BookRepository) GetBookByID(ctx context.Context, db *sql.DB, bookID int) (*entity.Book, *entity.ErrorResponse) {
	result := db.QueryRowContext(ctx, "SELECT * FROM books WHERE id = ?", bookID)

	var book entity.Book
	var cardID sql.NullInt64
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
		&cardID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			message := fmt.Sprintf("book id %d not found", bookID)
			return nil, helper.ErrorResponse(http.StatusNotFound, message)
		}
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan book")
	}
	book.CardID = int(cardID.Int64)
	return &book, nil
}

func (*BookRepository) DeleteBookByID(ctx context.Context, tx *sql.Tx, bookID int) *entity.ErrorResponse {
	_, err := tx.ExecContext(ctx, "DELETE FROM books WHERE id = ?", bookID)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "failed to delete book")
	}

	return nil
}

func (*BookRepository) DeleteCardIDFromBook(ctx context.Context, tx *sql.Tx, cardID int) *entity.ErrorResponse {
	_, err := tx.ExecContext(ctx, "UPDATE books SET card_id = null WHERE card_id = ?", cardID)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	return nil
}

func (*BookRepository) UpdateBook(ctx context.Context, tx *sql.Tx, book *entity.Book) *entity.ErrorResponse {
	_, err := tx.ExecContext(ctx, "UPDATE books SET title=?, author=?, publisher=?, published_date=?, isbn=?, pages=?, language=?, genre=?, description=?, card_id=? WHERE id=?",
		book.Title,
		book.Author,
		book.Publisher,
		book.PublishedDate,
		book.ISBN,
		book.Pages,
		book.Language,
		book.Genre,
		book.Description,
		book.CardID,
		book.ID,
	)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "failed to update book")
	}
	return nil
}

func (*BookRepository) GetBookByCardID(ctx context.Context, db *sql.DB, cardID int) (*entity.Book, *entity.ErrorResponse) {
	result := db.QueryRowContext(ctx, "SELECT * FROM books WHERE card_id = ?", cardID)

	var book entity.Book
	var cardId sql.NullInt64
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
		&cardId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorResponse(http.StatusNotFound, "data not found")
		}
		if isError := strings.Contains(err.Error(), "a foreign key constraint fails"); isError {
			return nil, helper.ErrorResponse(http.StatusUnprocessableEntity, "failed to update book, data card_id not valid.")
		}

		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan book")
	}
	book.CardID = cardID
	return &book, nil
}

func (*BookRepository) InsertBook(ctx context.Context, tx *sql.Tx, book *entity.Book) *entity.ErrorResponse {

	_, err := tx.ExecContext(ctx, "INSERT INTO books (title, author, publisher, published_date, isbn, pages, language, genre, description, card_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		book.Title,
		book.Author,
		book.Publisher,
		book.PublishedDate,
		book.ISBN,
		book.Pages,
		book.Language,
		book.Genre,
		book.Description,
		book.CardID,
	)
	if err != nil {
		if isError := strings.Contains(err.Error(), "a foreign key constraint fails"); isError {
			return helper.ErrorResponse(http.StatusUnprocessableEntity, "failed to insert book, data card_id not valid.")
		}
		return helper.ErrorResponse(http.StatusInternalServerError, "failed to insert book")
	}
	return nil
}

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

type CardRepositoryInterface interface {
	GetCards(ctx context.Context, db *sql.DB, page, pageSize int) ([]*entity.Card, *entity.ErrorResponse)
	GetCardByID(ctx context.Context, db *sql.DB, id int) (*entity.Card, *entity.ErrorResponse)
	GetCardByUID(ctx context.Context, db *sql.DB, uid int) (*entity.Card, *entity.ErrorResponse)
	InsertCard(ctx context.Context, db *sql.DB, rfid *entity.Card) *entity.ErrorResponse
	UpdateCard(ctx context.Context, db *sql.DB, rfid *entity.Card) *entity.ErrorResponse
	DeleteCard(ctx context.Context, db *sql.DB, id int) *entity.ErrorResponse
}

type CardRepository struct{}

func NewCardRepository() *CardRepository {
	return &CardRepository{}
}

func (r *CardRepository) GetCards(ctx context.Context, db *sql.DB, page, pageSize int) ([]*entity.Card, *entity.ErrorResponse) {

	var query string
	if page != 0 && pageSize != 0 {
		offset := (page - 1) * pageSize
		query = fmt.Sprintf("SELECT * FROM card_rfid LIMIT %d OFFSET %d", pageSize, offset)
	} else {
		query = "SELECT * FROM card_rfid"
	}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, err.Error())
	}

	var cards []*entity.Card
	for rows.Next() {
		var card entity.Card
		err := rows.Scan(&card.ID, &card.UID, &card.Type)
		if err != nil {
			return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan card")
		}
		cards = append(cards, &card)
	}

	return cards, nil
}

func (r *CardRepository) GetCardByID(ctx context.Context, db *sql.DB, id int) (*entity.Card, *entity.ErrorResponse) {
	row := db.QueryRowContext(ctx, "SELECT * FROM card_rfid WHERE id = ?", id)

	var rfid entity.Card
	err := row.Scan(&rfid.ID, &rfid.UID, &rfid.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorResponse(http.StatusNotFound, "card not found")
		}
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan card")
	}

	return &rfid, nil
}

func (r *CardRepository) GetCardByUID(ctx context.Context, db *sql.DB, uid string) (*entity.Card, *entity.ErrorResponse) {
	row := db.QueryRowContext(ctx, "SELECT * FROM card_rfid WHERE uid = ?", uid)

	var rfid entity.Card
	err := row.Scan(&rfid.ID, &rfid.UID, &rfid.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorResponse(http.StatusNotFound, "card not found")
		}
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan card")
	}

	return &rfid, nil
}

func (r *CardRepository) InsertCard(ctx context.Context, db *sql.DB, rfid *entity.Card) *entity.ErrorResponse {
	_, err := db.ExecContext(ctx, "INSERT INTO card_rfid (uid, type) VALUES (?, ?)", rfid.UID, rfid.Type)
	if err != nil {
		fmt.Println(err)
		return helper.ErrorResponse(http.StatusInternalServerError, "failed to insert card")
	}

	return nil
}

func (r *CardRepository) UpdateCard(ctx context.Context, db *sql.DB, rfid *entity.Card) *entity.ErrorResponse {
	_, err := db.ExecContext(ctx, "UPDATE card_rfid SET uid = ?, type = ? WHERE id = ?", rfid.UID, rfid.Type, rfid.ID)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "failed to update card")
	}

	return nil
}

func (r *CardRepository) DeleteCard(ctx context.Context, db *sql.DB, id int) *entity.ErrorResponse {
	_, err := db.ExecContext(ctx, "DELETE FROM card_rfid WHERE id = ?", id)
	if err != nil {
		var statusCode int
		if strings.Contains(err.Error(), "foreign key constraint") {
			statusCode = http.StatusConflict
		}
		return helper.ErrorResponse(statusCode, "failed to delete card")
	}

	return nil
}

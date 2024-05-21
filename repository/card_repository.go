package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/smart-library-be/entity"
)

type CardRepositoryInterface interface {
	GetCardByID(ctx context.Context, db *sql.DB, id int) (*entity.Card, error)
	GetCardByUID(ctx context.Context, db *sql.DB, uid int) (*entity.Card, error)
	InsertCard(ctx context.Context, db *sql.DB, rfid *entity.Card) error
	DeleteCard(ctx context.Context, db *sql.DB, id int) error
}

type CardRepository struct{}

func NewCardRepository() *CardRepository {
	return &CardRepository{}
}

func (r *CardRepository) GetCardByID(ctx context.Context, db *sql.DB, id int) (*entity.Card, error) {
	row := db.QueryRowContext(ctx, "SELECT id, uid, type FROM card_rfid WHERE id = ?", id)

	var rfid *entity.Card
	err := row.Scan(&rfid.ID, &rfid.UID, &rfid.Type)
	if err != nil {
		return nil, err
	}

	return rfid, nil
}

func (r *CardRepository) GetCardByUID(ctx context.Context, db *sql.DB, id int) (*entity.Card, error) {
	row := db.QueryRowContext(ctx, "SELECT id, uid, type FROM card_rfid WHERE uid = ?", id)

	var rfid entity.Card
	err := row.Scan(&rfid.ID, &rfid.UID, &rfid.Type)
	if err != nil {
		return nil, err
	}

	return &rfid, nil
}

func (r *CardRepository) InsertCard(ctx context.Context, db *sql.DB, rfid *entity.Card) error {
	_, err := db.ExecContext(ctx, "INSERT INTO card_rfid VALUES (?, ?, ?)", rfid.ID, rfid.UID, rfid.Type)
	if err != nil {
		return err
	}

	return nil
}

func (r *CardRepository) DeleteCard(ctx context.Context, db *sql.DB, id int) error {
	_, err := db.ExecContext(ctx, "DELETE FROM card_rfid WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

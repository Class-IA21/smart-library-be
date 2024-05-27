package repository

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
)

type AccountsRepositoryInterface interface {
	InsertAccount(ctx context.Context, tx *sql.Tx, account *entity.AccountRequest) (accountId int, error *entity.ErrorResponse)
	UpdateAccount(ctx context.Context, tx *sql.Tx, account *entity.AccountRequest) *entity.ErrorResponse
	DeleteAccount(ctx context.Context, tx *sql.Tx, accountID int) *entity.ErrorResponse
	ChangePassword(ctx context.Context, tx *sql.Tx, account *entity.AccountChangePasswordRequest, accountID int) *entity.ErrorResponse
	GetAccountByID(ctx context.Context, db *sql.DB, accountID int) (*entity.AccountResponse, *entity.ErrorResponse)
	GetAccountByEmail(ctx context.Context, db *sql.DB, level string) (*entity.AccountResponse, *entity.ErrorResponse)
}

type AccountsRepository struct{}

func NewAccountsRepository() *AccountsRepository {
	return &AccountsRepository{}
}

func (*AccountsRepository) InsertAccount(ctx context.Context, tx *sql.Tx, account *entity.AccountRequest) (accountId int, error *entity.ErrorResponse) {
	result, err := tx.ExecContext(ctx, "INSERT INTO accounts(email, password, level) VALUES (?,?,?)", account.Email, account.Password, account.Level)
	if err != nil {
		return 0, helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return int(id), nil
}

func (*AccountsRepository) UpdateAccount(ctx context.Context, tx *sql.Tx, account *entity.AccountRequest) *entity.ErrorResponse {
	_, err := tx.ExecContext(ctx, "UPDATE accounts SET email = ?, level = ? WHERE id = ?", account.Email, account.Level, account.ID)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return nil
}

func (*AccountsRepository) DeleteAccount(ctx context.Context, tx *sql.Tx, accountID int) *entity.ErrorResponse {
	_, err := tx.ExecContext(ctx, "DELETE FROM accounts WHERE id = ?", accountID)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return nil
}

func (*AccountsRepository) ChangePassword(ctx context.Context, tx *sql.Tx, account *entity.AccountChangePasswordRequest, accountID int) *entity.ErrorResponse {
	_, err := tx.ExecContext(ctx, "UPDATE accounts SET password = ? WHERE id = ?", account.Password, account.NewPassword, accountID)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return nil
}

func (*AccountsRepository) GetAccountByID(ctx context.Context, db *sql.DB, accountID int) (*entity.AccountResponse, *entity.ErrorResponse) {
	rows := db.QueryRowContext(ctx, "SELECT id, email, level FROM accounts WHERE id = ?", accountID)

	var account entity.Account
	err := rows.Scan(&account.ID, &account.Email, &account.Level)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	var response entity.AccountResponse = entity.AccountResponse{
		ID:    account.ID,
		Email: account.Email,
		Level: account.Level,
	}

	return &response, nil
}

func (*AccountsRepository) GetAccountByEmail(ctx context.Context, db *sql.DB, email string) (*entity.AccountResponse, *entity.ErrorResponse) {
	rows := db.QueryRowContext(ctx, "SELECT id, email, level FROM accounts WHERE email = ?", email)

	var account entity.Account
	err := rows.Scan(&account.ID, &account.Email, &account.Level)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	var response entity.AccountResponse = entity.AccountResponse{
		ID:    account.ID,
		Email: account.Email,
		Level: account.Level,
	}

	return &response, nil
}

func getAccountPasswordHash(ctx context.Context, db *sql.DB, email string) (string, *entity.ErrorResponse) {
	rows := db.QueryRowContext(ctx, "SELECT password FROM accounts WHERE email = ? LIMIT 1", email)

	var passwordHash string
	err := rows.Scan(&passwordHash)
	if err != nil {
		return "", helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return passwordHash, nil
}

func (*AccountsRepository) GetAccountPasswordHash(ctx context.Context, db *sql.DB, email string) (string, *entity.ErrorResponse) {
	return getAccountPasswordHash(ctx, db, email)
}

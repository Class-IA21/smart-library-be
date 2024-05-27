package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"net/http"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/repository"
)

type AccountServicesInterface interface {
	LoginAccount(ctx context.Context, account *entity.Login) (*entity.AccountResponse, *entity.ErrorResponse)
	RegisterAccount(ctx context.Context, account *entity.AccountRequest) *entity.ErrorResponse
	InsertAccount(ctx context.Context, account *entity.AccountRequest) *entity.ErrorResponse
	UpdateAccount(ctx context.Context, account *entity.AccountRequest) *entity.ErrorResponse
	DeleteAccount(ctx context.Context, accountID int) *entity.ErrorResponse
	ChangePassword(ctx context.Context, account *entity.AccountChangePasswordRequest, accountID int) *entity.ErrorResponse
	GetAccountByID(ctx context.Context, accountID int) (*entity.AccountResponse, *entity.ErrorResponse)
	GetAccountByName(ctx context.Context, name string) ([]*entity.AccountResponse, *entity.ErrorResponse)
	GetAccountByEmail(ctx context.Context, email string) (*entity.AccountResponse, *entity.ErrorResponse)
	getAccountPasswordHash(ctx context.Context, email string) (string, *entity.ErrorResponse)
}

type AccountServices struct {
	*sql.DB
	*repository.AccountsRepository
	*StudentServices
}

func NewAccountServices(DB *sql.DB, ar *repository.AccountsRepository, ss *StudentServices) *AccountServices {
	return &AccountServices{
		DB:                 DB,
		AccountsRepository: ar,
		StudentServices:    ss,
	}
}

func (s *AccountServices) LoginAccount(ctx context.Context, account *entity.Login) (*entity.AccountResponse, *entity.ErrorResponse) {
	accountDetail, err := s.GetAccountByEmail(ctx, account.Email)
	if err != nil {
		return nil, err
	}

	passwordHash, err := s.getAccountPasswordHash(ctx, account.Email)
	if err != nil {
		return nil, err
	}
	if isValid := helper.CheckPasswordHash(account.Password, passwordHash); !isValid {
		return nil, helper.ErrorResponse(http.StatusBadRequest, "Email or Password is invalid")
	}

	return accountDetail, nil
}

func (s *AccountServices) RegisterAccount(ctx context.Context, r *entity.AccountRequest) *entity.ErrorResponse {

	hashPassword, err := helper.HashPassword(r.Password)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	fmt.Println(hashPassword)

	account := entity.AccountRequest{
		Name:     r.Name,
		NPM:      r.Name,
		Email:    r.Email,
		Password: hashPassword,
		Level:    r.Level,
	}

	return s.InsertAccount(ctx, &account)
}

func (s *AccountServices) InsertAccount(ctx context.Context, account *entity.AccountRequest) *entity.ErrorResponse {
	isEmailExists, _ := s.GetAccountByEmail(ctx, account.Email)
	if isEmailExists != nil {
		return helper.ErrorResponse(http.StatusConflict, "Email already exists.")
	}

	if account.Level == "student" {
		isNPMExists, _ := s.StudentServices.GetStudentByNPM(ctx, account.NPM)
		if isNPMExists != nil {
			return helper.ErrorResponse(http.StatusConflict, "NPM already exists.")
		}
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer tx.Commit()

	accountId, errorResponse := s.AccountsRepository.InsertAccount(ctx, tx, account)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	var student entity.Student
	student.AccountID = accountId
	student.Name = account.Name
	student.NPM = account.NPM

	errorResponse = s.StudentServices.InsertStudent(ctx, tx, &student)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	return nil
}

func (s *AccountServices) UpdateAccount(ctx context.Context, account *entity.AccountRequest) *entity.ErrorResponse {
	_, errorResponse := s.AccountsRepository.GetAccountByID(ctx, s.DB, account.ID)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer tx.Commit()

	_, errorResponse = s.AccountsRepository.InsertAccount(ctx, tx, account)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	return nil
}

func (s *AccountServices) DeleteAccount(ctx context.Context, accountID int) *entity.ErrorResponse {
	_, errorResponse := s.AccountsRepository.GetAccountByID(ctx, s.DB, accountID)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer tx.Commit()

	errorResponse = s.AccountsRepository.DeleteAccount(ctx, tx, accountID)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	return nil
}

func (s *AccountServices) ChangePassword(ctx context.Context, account *entity.AccountChangePasswordRequest, accountID int) *entity.ErrorResponse {
	_, errorResponse := s.AccountsRepository.GetAccountByID(ctx, s.DB, accountID)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}
	defer tx.Commit()

	errorResponse = s.AccountsRepository.ChangePassword(ctx, tx, account, accountID)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	return nil
}

func (s *AccountServices) GetAccountByID(ctx context.Context, accountID int) (*entity.AccountResponse, *entity.ErrorResponse) {
	return s.AccountsRepository.GetAccountByID(ctx, s.DB, accountID)
}

func (s *AccountServices) GetAccountByEmail(ctx context.Context, email string) (*entity.AccountResponse, *entity.ErrorResponse) {
	return s.AccountsRepository.GetAccountByEmail(ctx, s.DB, email)
}

func (s *AccountServices) getAccountPasswordHash(ctx context.Context, email string) (string, *entity.ErrorResponse) {
	return s.AccountsRepository.GetAccountPasswordHash(ctx, s.DB, email)
}

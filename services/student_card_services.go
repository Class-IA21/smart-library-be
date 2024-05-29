package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
)

type StudentCardServiceInterface interface {
	InsertStudent(ctx context.Context, student *entity.Student) *entity.ErrorResponse
	UpdateStudent(ctx context.Context, student *entity.Student) *entity.ErrorResponse
	DeleteStudent(ctx context.Context, id int) *entity.ErrorResponse
}

type StudentCardServices struct {
	*sql.DB
	*CardServices
	*StudentServices
}

func NewStudentCardServices(DB *sql.DB, cs *CardServices, ss *StudentServices) *StudentCardServices {
	return &StudentCardServices{
		DB:              DB,
		CardServices:    cs,
		StudentServices: ss,
	}
}

func (s *StudentCardServices) InsertStudent(ctx context.Context, student *entity.Student) *entity.ErrorResponse {
	_, errorResponse := s.CardServices.GetCardByID(ctx, student.CardID)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "transaction start failed")
	}
	defer tx.Rollback()

	errResp := s.StudentRepository.InsertStudent(ctx, tx, student)
	if errResp != nil {
		return errResp
	}

	if err := tx.Commit(); err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "transaction commit failed")
	}

	return nil
}

func (s *StudentCardServices) UpdateStudent(ctx context.Context, student *entity.Student) *entity.ErrorResponse {
	_, errorResponse := s.StudentServices.GetStudentByID(ctx, student.ID)
	if errorResponse != nil {
		return errorResponse
	}

	_, errorResponse = s.CardServices.GetCardByID(ctx, student.CardID)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "transaction start failed")
	}
	defer tx.Rollback()

	errorResponse = s.StudentRepository.UpdateStudent(ctx, tx, student.ID, student)
	if errorResponse != nil {
		return errorResponse
	}

	if err := tx.Commit(); err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "transaction commit failed")
	}

	return nil
}

func (s *StudentCardServices) DeleteStudent(ctx context.Context, id int) *entity.ErrorResponse {
	student, errorResponse := s.StudentRepository.GetStudentByID(ctx, s.DB, id)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "transaction start failed")
	}
	defer tx.Commit()

	errorResponse = s.StudentRepository.DeleteStudent(ctx, tx, id)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	errorResponse = s.CardServices.DeleteCard(ctx, student.CardID)
	if errorResponse != nil {
		tx.Rollback()
		return errorResponse
	}

	return nil
}

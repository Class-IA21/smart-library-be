package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/repository"
)

type StudentServiceInterface interface {
	GetStudentByID(ctx context.Context, id int) (*entity.Student, *entity.ErrorResponse)
	DeleteStudent(ctx context.Context, id int) *entity.ErrorResponse
	GetStudents(ctx context.Context, page int, pageSize int) ([]*entity.Student, *entity.ErrorResponse)
}

type StudentServices struct {
	db *sql.DB
	*repository.StudentRepository
}

func NewStudentServices(db *sql.DB, studentRepository *repository.StudentRepository) *StudentServices {
	return &StudentServices{
		db:                db,
		StudentRepository: studentRepository,
	}
}

func (s *StudentServices) GetStudentByCardID(ctx context.Context, cardID int) (*entity.Student, *entity.ErrorResponse) {
	if cardID <= 0 {
		return nil, helper.ErrorResponse(http.StatusBadRequest, "invalid card id")
	}

	return s.StudentRepository.GetStudentByCardID(ctx, s.db, cardID)
}

func (s *StudentServices) GetStudentByID(ctx context.Context, id int) (*entity.Student, *entity.ErrorResponse) {
	return s.StudentRepository.GetStudentByID(ctx, s.db, id)
}

func (s *StudentServices) DeleteStudent(ctx context.Context, id int) *entity.ErrorResponse {
	_, errorResponse := s.StudentRepository.GetStudentByID(ctx, s.db, id)
	if errorResponse != nil {
		return errorResponse
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "transaction start failed")
	}
	defer tx.Rollback()

	errorResponse = s.StudentRepository.DeleteStudent(ctx, tx, id)
	if errorResponse != nil {
		return errorResponse
	}

	if err := tx.Commit(); err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "transaction commit failed")
	}

	return nil
}

func (s *StudentServices) GetStudents(ctx context.Context, page int, pageSize int) ([]*entity.Student, *entity.ErrorResponse) {
	offset := (page - 1) * pageSize
	students, err := s.StudentRepository.GetStudents(ctx, s.db, pageSize, offset)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to retrieve students")
	}
	return students, nil
}

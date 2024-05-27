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
	GetStudentByID(ctx context.Context, id int) (*entity.StudentResponse, *entity.ErrorResponse)
	GetStudentByNPM(ctx context.Context, npm string) (*entity.StudentResponse, *entity.ErrorResponse)
	DeleteStudent(ctx context.Context, id int) *entity.ErrorResponse
	GetStudents(ctx context.Context, page int, pageSize int) ([]*entity.StudentResponse, *entity.ErrorResponse)
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

func (s *StudentServices) GetStudentByCardID(ctx context.Context, cardID int) (*entity.StudentResponse, *entity.ErrorResponse) {
	if cardID <= 0 {
		return nil, helper.ErrorResponse(http.StatusBadRequest, "invalid card id")
	}

	var studentResponse entity.StudentResponse
	student, err := s.StudentRepository.GetStudentByCardID(ctx, s.db, cardID)
	if err != nil {
		return nil, err
	}
	studentResponse.ID = student.ID
	studentResponse.Name = student.Name
	studentResponse.NPM = student.NPM
	studentResponse.CardID = student.CardID

	return &studentResponse, nil
}

func (s *StudentServices) GetStudentByID(ctx context.Context, id int) (*entity.StudentResponse, *entity.ErrorResponse) {
	var studentResponse entity.StudentResponse
	student, err := s.StudentRepository.GetStudentByID(ctx, s.db, id)
	if err != nil {
		return nil, err
	}
	studentResponse.ID = student.ID
	studentResponse.Name = student.Name
	studentResponse.NPM = student.NPM
	studentResponse.CardID = student.CardID

	return &studentResponse, nil
}

func (s *StudentServices) GetStudentByEmail(ctx context.Context, email string) (*entity.StudentResponse, *entity.ErrorResponse) {
	var studentResponse entity.StudentResponse
	student, err := s.StudentRepository.GetStudentByEmail(ctx, s.db, email)
	if err != nil {
		return nil, err
	}
	studentResponse.ID = student.ID
	studentResponse.Name = student.Name
	studentResponse.NPM = student.NPM
	studentResponse.CardID = student.CardID

	return &studentResponse, nil
}

func (s *StudentServices) GetStudentByNPM(ctx context.Context, npm string) (*entity.StudentResponse, *entity.ErrorResponse) {
	var studentResponse entity.StudentResponse
	student, err := s.StudentRepository.GetStudentByNPM(ctx, s.db, npm)
	if err != nil {
		return nil, err
	}
	studentResponse.ID = student.ID
	studentResponse.Name = student.Name
	studentResponse.NPM = student.NPM
	studentResponse.CardID = student.CardID

	return &studentResponse, nil
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

func (s *StudentServices) GetStudents(ctx context.Context, page int, pageSize int) ([]*entity.StudentResponse, *entity.ErrorResponse) {
	offset := (page - 1) * pageSize

	var studentResponse []*entity.StudentResponse
	students, err := s.StudentRepository.GetStudents(ctx, s.db, pageSize, offset)
	if err != nil {
		return nil, err
	}

	for _, student := range students {
		studentResponse = append(studentResponse, &entity.StudentResponse{
			ID:     student.ID,
			Name:   student.Name,
			NPM:    student.NPM,
			CardID: student.CardID,
		})
	}

	return studentResponse, nil
}

package repository

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
)

type StudentRepositoryInterface interface {
	GetAllStudent(ctx context.Context, db *sql.DB) ([]*entity.Student, *entity.ErrorResponse)
	GetStudentByCardID(ctx context.Context, db *sql.DB, cardId int) (*entity.Student, *entity.ErrorResponse)
	GetStudentByID(ctx context.Context, db *sql.DB, id int) (*entity.Student, *entity.ErrorResponse)
	InsertStudent(ctx context.Context, tx *sql.Tx, student *entity.Student) *entity.ErrorResponse
	DeleteStudent(ctx context.Context, tx *sql.Tx, id int) *entity.ErrorResponse
}

type StudentRepository struct{}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (*StudentRepository) GetAllStudent(ctx context.Context, db *sql.DB) ([]*entity.Student, *entity.ErrorResponse) {
	var students []*entity.Student

	rows, err := db.QueryContext(ctx, "SELECT * FROM students")
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to get students")
	}

	for rows.Next() {
		var student entity.Student
		err := rows.Scan(&student.ID, &student.Name, &student.NPM, &student.CardID)
		if err != nil {
			return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan student")
		}
		students = append(students, &student)
	}

	return students, nil
}

func (*StudentRepository) GetStudentByCardID(ctx context.Context, db *sql.DB, cardId int) (*entity.Student, *entity.ErrorResponse) {
	var student entity.Student

	row := db.QueryRowContext(ctx, "SELECT * FROM students WHERE card_id = ?", cardId)
	err := row.Scan(&student.ID, &student.Name, &student.NPM, &student.CardID)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan student")
	}

	return &student, nil
}

func (*StudentRepository) GetStudentByID(ctx context.Context, db *sql.DB, id int) (*entity.Student, *entity.ErrorResponse) {
	var student entity.Student

	row := db.QueryRowContext(ctx, "SELECT * FROM students WHERE id = ?", id)
	err := row.Scan(&student.ID, &student.Name, &student.NPM, &student.CardID)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan student")
	}

	return &student, nil
}

func (*StudentRepository) InsertStudent(ctx context.Context, tx *sql.Tx, student *entity.Student) *entity.ErrorResponse {
	_, err := tx.ExecContext(ctx, "INSERT INTO students VALUES (?, ?, ?, ?)", student.ID, student.Name, student.NPM, student.CardID)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "failed to insert student")
	}

	return nil
}

func (*StudentRepository) DeleteStudent(ctx context.Context, tx *sql.Tx, id int) *entity.ErrorResponse {
	_, err := tx.ExecContext(ctx, "DELETE FROM students WHERE id = ?", id)
	if err != nil {
		if strings.Contains(err.Error(), "foreign key constraint fails") {
			return helper.ErrorResponse(http.StatusConflict, "cannot delete student")
		}
		return helper.ErrorResponse(http.StatusInternalServerError, "failed to delete student")
	}

	return nil
}

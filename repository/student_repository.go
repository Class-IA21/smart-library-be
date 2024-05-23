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

type StudentRepositoryInterface interface {
	GetStudents(ctx context.Context, db *sql.DB, page int, pageSize int) ([]*entity.Student, *entity.ErrorResponse)
	GetStudentByCardID(ctx context.Context, db *sql.DB, cardId int) (*entity.Student, *entity.ErrorResponse)
	GetStudentByID(ctx context.Context, db *sql.DB, id int) (*entity.Student, *entity.ErrorResponse)
	InsertStudent(ctx context.Context, tx *sql.Tx, student *entity.Student) *entity.ErrorResponse
	DeleteStudent(ctx context.Context, tx *sql.Tx, id int) *entity.ErrorResponse
	UpdateStudent(ctx context.Context, tx *sql.Tx, id int, student *entity.Student) *entity.ErrorResponse
}

type StudentRepository struct{}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (*StudentRepository) GetStudents(ctx context.Context, db *sql.DB, page int, pageSize int) ([]*entity.Student, *entity.ErrorResponse) {
	var students []*entity.Student
	var query string

	if page != 0 && pageSize != 0 {
		offset := (page - 1) * pageSize
		query = fmt.Sprintf("SELECT * FROM students LIMIT %d OFFSET %d", pageSize, offset)
	} else {
		query = "SELECT * FROM students"
	}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to get students")
	}
	defer rows.Close()

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
		if err == sql.ErrNoRows {
			return nil, helper.ErrorResponse(http.StatusNotFound, "data not found")
		}
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan student")
	}

	return &student, nil
}

func (*StudentRepository) GetStudentByID(ctx context.Context, db *sql.DB, id int) (*entity.Student, *entity.ErrorResponse) {
	var student entity.Student

	row := db.QueryRowContext(ctx, "SELECT * FROM students WHERE id = ?", id)
	err := row.Scan(&student.ID, &student.Name, &student.NPM, &student.CardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorResponse(http.StatusNotFound, "data not found")
		}
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

func (*StudentRepository) UpdateStudent(ctx context.Context, tx *sql.Tx, id int, student *entity.Student) *entity.ErrorResponse {
	_, err := tx.ExecContext(ctx, "UPDATE students SET name = ?, npm = ?, card_id = ? WHERE id = ?", student.Name, student.NPM, student.CardID, id)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "failed to update student")
	}

	return nil
}

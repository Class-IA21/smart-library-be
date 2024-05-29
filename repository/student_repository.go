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
	GetStudentByAccountID(ctx context.Context, db *sql.DB, accountID int) (*entity.Student, *entity.ErrorResponse)
	GetStudentByNPM(ctx context.Context, db *sql.DB, npm string) (*entity.Student, *entity.ErrorResponse)
	InsertStudent(ctx context.Context, tx *sql.Tx, student *entity.Student) *entity.ErrorResponse
	DeleteStudent(ctx context.Context, tx *sql.Tx, id int) *entity.ErrorResponse
	DeleteCardIDFromStudent(ctx context.Context, tx *sql.Tx, cardID int) *entity.ErrorResponse
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
		query = fmt.Sprintf("SELECT id, name, npm, card_id FROM students LIMIT %d OFFSET %d", pageSize, offset)
	} else {
		query = "SELECT id, name, npm, card_id FROM students"
	}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to get students")
	}
	defer rows.Close()

	for rows.Next() {
		var student entity.Student
		var cardID sql.NullInt64
		err := rows.Scan(&student.ID, &student.Name, &student.NPM, &cardID)
		if err != nil {
			return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan student")
		}
		student.CardID = int(cardID.Int64)
		students = append(students, &student)
	}

	return students, nil
}

func (*StudentRepository) GetStudentByCardID(ctx context.Context, db *sql.DB, cardId int) (*entity.Student, *entity.ErrorResponse) {
	var student entity.Student
	var cardID sql.NullInt64

	row := db.QueryRowContext(ctx, "SELECT id, name, npm, card_id FROM students WHERE card_id = ?", cardId)
	err := row.Scan(&student.ID, &student.Name, &student.NPM, &cardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorResponse(http.StatusNotFound, "data student not found")
		}
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan student")
	}
	student.CardID = int(cardID.Int64)

	return &student, nil
}

func (*StudentRepository) GetStudentByID(ctx context.Context, db *sql.DB, id int) (*entity.Student, *entity.ErrorResponse) {
	var student entity.Student
	var cardID sql.NullInt64

	row := db.QueryRowContext(ctx, "SELECT id, name, npm, card_id FROM students WHERE id = ?", id)
	err := row.Scan(&student.ID, &student.Name, &student.NPM, &cardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorResponse(http.StatusNotFound, "data student not found")
		}
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan student")
	}
	student.CardID = int(cardID.Int64)

	return &student, nil
}

func (*StudentRepository) GetStudentByAccountID(ctx context.Context, db *sql.DB, accountID int) (*entity.Student, *entity.ErrorResponse) {
	var student entity.Student
	var cardID sql.NullInt64

	row := db.QueryRowContext(ctx, "SELECT id, name, npm, card_id FROM students WHERE account_id = ?", accountID)
	err := row.Scan(&student.ID, &student.Name, &student.NPM, &cardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorResponse(http.StatusNotFound, "data student not found")
		}
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan student")
	}
	student.CardID = int(cardID.Int64)

	return &student, nil
}

func (r *StudentRepository) GetStudentByNPM(ctx context.Context, db *sql.DB, npm string) (*entity.Student, *entity.ErrorResponse) {
	var student entity.Student
	var cardID sql.NullInt64

	query := "SELECT id, name, npm, card_id FROM students WHERE npm = ? LIMIT 1"
	row := db.QueryRowContext(ctx, query, npm)
	err := row.Scan(&student.ID, &student.Name, &student.NPM, &cardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorResponse(http.StatusNotFound, "data student not found")
		}
		return nil, helper.ErrorResponse(http.StatusInternalServerError, "failed to scan student")
	}
	student.CardID = int(cardID.Int64)

	return &student, nil
}

func (*StudentRepository) InsertStudent(ctx context.Context, tx *sql.Tx, student *entity.Student) *entity.ErrorResponse {
	var query string
	var args []interface{}

	if student.CardID == 0 {
		query = "INSERT INTO students (name, npm, account_id) VALUES (?, ?, ?)"
		args = []interface{}{student.Name, student.NPM, student.AccountID}
	} else {
		query = "INSERT INTO students (name, npm, card_id, account_id) VALUES (?, ?, ?, ?)"
		args = []interface{}{student.Name, student.NPM, student.CardID, student.AccountID}
	}

	_, err := tx.ExecContext(ctx, query, args...)
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

func (*StudentRepository) DeleteCardIDFromStudent(ctx context.Context, tx *sql.Tx, cardID int) *entity.ErrorResponse {
	_, err := tx.ExecContext(ctx, "UPDATE students SET card_id = null WHERE card_id = ?", cardID)
	if err != nil {
		return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
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

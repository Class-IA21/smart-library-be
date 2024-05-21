package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/smart-library-be/entity"
)

type StudentRepositoryInterface interface {
	GetAllStudent(ctx context.Context, db *sql.DB) ([]*entity.Student, error)
	GetStudentByCardID(ctx context.Context, db *sql.DB, cardId int) (*entity.Student, error)
	GetStudentByID(ctx context.Context, db *sql.DB, id int) (*entity.Student, error)
	InsertStudent(ctx context.Context, tx *sql.Tx, student *entity.Student) error
	DeleteStudent(ctx context.Context, tx *sql.Tx, id int) error
}

type StudentRepository struct{}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (*StudentRepository) GetAllStudent(ctx context.Context, db *sql.DB) ([]*entity.Student, error) {
	var students []*entity.Student

	rows, err := db.QueryContext(ctx, "SELECT * FROM students")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var student entity.Student
		err := rows.Scan(&student.ID, &student.Name, &student.NPM, &student.CardID)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}

func (*StudentRepository) GetStudentByCardID(ctx context.Context, db *sql.DB, cardId int) (*entity.Student, error) {
	var student entity.Student

	row := db.QueryRowContext(ctx, "SELECT * FROM students WHERE card_id = ?", cardId)
	err := row.Scan(&student.ID, &student.Name, &student.NPM, &student.CardID)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (*StudentRepository) GetStudentByID(ctx context.Context, db *sql.DB, id int) (*entity.Student, error) {
	var student entity.Student

	row := db.QueryRowContext(ctx, "SELECT * FROM students WHERE id = ?", id)
	err := row.Scan(&student.ID, &student.Name, &student.NPM, &student.CardID)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (*StudentRepository) InsertStudent(ctx context.Context, tx *sql.Tx, student *entity.Student) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO students VALUES (?, ?, ?, ?)", student.ID, student.Name, student.NPM, student.CardID)
	if err != nil {
		return err
	}

	return nil
}

func (*StudentRepository) DeleteStudent(ctx context.Context, tx *sql.Tx, id int) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM students WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

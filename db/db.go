package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dimassfeb-09/smart-library-be/helper"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	envDB := helper.GetEnvDatabase()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", envDB.User, envDB.Password, envDB.Host, envDB.Port, envDB.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

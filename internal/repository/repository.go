package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DataBase struct {
	//db sqlx.DB
}

func New() (*sqlx.DB, error) {
	var err error
	var db *sqlx.DB
	db, err = sqlx.Connect("postgres", "user=postgres dbname=postgres password=lto host=localhost sslmode=disable port=5555")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

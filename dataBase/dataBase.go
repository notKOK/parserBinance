package dataBase

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DataBase struct {
	Db *sqlx.DB
}

func (*DataBase) New() (*DataBase, error) {
	var base DataBase
	var err error
	base.Db, err = sqlx.Connect("postgres", "user=postgres dbname=postgres password=lto host=localhost sslmode=disable port=5555")
	if err != nil {
		log.Fatal(err)
	}
	err = base.Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return &base, nil
}

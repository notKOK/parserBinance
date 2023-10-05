package useCase

import (
	"log"
	"parser/dataBase"
)

type useCase struct {
}

func AddTicker() {
	base := dataBase.DataBase{}
	d, err := base.New()
	if err != nil {
		log.Fatal(err)
	}
	err = d.Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

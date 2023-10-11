package useCase

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"math"
	"parser/internal/repository"
	"parser/internal/request"
	"parser/models"
	"strconv"
	"sync"
	"time"
)

type useCase struct {
	//interface database
}

var db *sqlx.DB
var mux sync.RWMutex

func AddTicker(ticker *models.Ticker) error {
	mux.Lock()
	defer mux.Unlock()
	var err error
	db, err = repository.New()
	var id int
	var test []models.Ticker
	query := fmt.Sprintf("SELECT id,name FROM tick_n WHERE name=$1") //find entity in database
	err = db.Select(&test, query, ticker.Name)
	if err != nil {
		return err
	}

	if test != nil {
		id = test[0].Id
		err = request.FetchBinance(ticker)

		query = fmt.Sprintf("SELECT rate FROM tick_r WHERE tick_r.ticker_id=$1") //find entity in database
		err = db.Select(&test, query, id)
		if test != nil {
			rate1, err1 := strconv.ParseFloat(test[0].Rate, 64)
			rate2, err2 := strconv.ParseFloat(ticker.Rate, 64)

			if err1 != nil || err2 != nil {
				log.Println("Ошибка при преобразовании строк в числа:", err1, err2)
				return nil
			}
			resultNum := math.Abs(rate1 - rate2)
			resultDiff := strconv.FormatFloat(resultNum, 'f', -1, 64)
			query = fmt.Sprintf("UPDATE tick_r SET rate = $1, upd_time = NOW() + interval '3 hour', diff_rate = $3 WHERE id = $2;") //update in found
			db.QueryRow(query, ticker.Rate, id, resultDiff)
			return nil
		}
	}

	query = fmt.Sprintf("INSERT INTO tick_n (name) values ($1) RETURNING id") //insert ticker.name
	row := db.QueryRow(query, ticker.Name)
	if err := row.Scan(&id); err != nil {
		log.Fatal(err)
	}
	ticker.Id = id

	err = request.FetchBinance(ticker)
	if err != nil {
		query = fmt.Sprintf("DELETE FROM tick_n WHERE id=$1;") //delete ticker if error
		row = db.QueryRow(query, id)
		return err
	}
	resultDiff := "0"
	query = fmt.Sprintf("SELECT rate FROM tick_r WHERE id=$1") //find entity in database
	err = db.Select(&test, query, id)

	if test != nil {

		rate1, err1 := strconv.ParseFloat(test[0].Rate, 64)
		rate2, err2 := strconv.ParseFloat(ticker.Rate, 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Ошибка при преобразовании строк в числа:", err1, err2)
			return nil
		}
		resultNum := math.Abs(rate1 - rate2)
		resultDiff = strconv.FormatFloat(resultNum, 'f', -1, 64)
	}

	query = fmt.Sprintf("INSERT INTO tick_r (rate, id, ticker_id, diff_rate, upd_time) values ($1, $2, $3, $4, NOW())") //insert other values
	row = db.QueryRow(query, ticker.Rate, ticker.Id, ticker.Id, resultDiff)
	return err
}

func FetchTicker(ticker *models.Ticker) error {
	mux.Lock()
	defer mux.Unlock()
	var err error
	db, err = repository.New()

	var extTicker []models.Ticker
	err = db.Select(&extTicker, `SELECT name, rate, diff_rate, upd_time FROM tick_n
	    inner join tick_r on tick_n.id = tick_r.id where name = $1;`, ticker.Name)

	if err != nil {
		return err
	}
	if extTicker == nil {
		log.Println("DATABASE EMPTY")
		return err
	}
	ticker.Rate = extTicker[0].Rate
	ticker.DiffRate = extTicker[0].DiffRate
	ticker.UpdTime = extTicker[0].UpdTime

	return nil
}

func RunUpdate(ticker *models.Ticker) {
	tick := time.NewTicker(5 * time.Second)

	for _ = range tick.C {
		err := AddTicker(ticker)
		if err != nil {
			return
		}
	}
}

package models

type Ticker struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"ticker" db:"name"`
	Rate     string `json:"rate" db:"rate"`
	DiffRate string `db:"diff_rate"`
	UpdTime  string `db:"upd_time"`
}

type Received struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

package models

type Ticker struct {
	Name string `json:"ticker"`
}

type Rate struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

var schema = `
CREATE TABLE tick_name IF NOT EXISTS(
    id SERIAL PRIMARY KEY,
    ticker_name text,
);

CREATE TABLE tick_rate IF NOT EXISTS(
    id integer PRIMARY KEY,
	ticker_rate text,
	FOREIGN KEY (id) REFERENCES tick_name (id) ON DELETE CASCADE
);`

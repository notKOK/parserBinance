CREATE TABLE tick_n (
    id SERIAL PRIMARY KEY,
    name text
);

CREATE TABLE tick_r (
    id integer PRIMARY KEY,
	rate text,
	diff_rate text,
	upd_time timestamp,
    ticker_id int,
    FOREIGN KEY (ticker_id) REFERENCES tick_n (id) ON DELETE CASCADE
);


package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type rate []struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func main() {
	http.HandleFunc("/addticker", addTicker)
	http.HandleFunc("/fetchticker", fetchTicker)

	fmt.Printf("Server started at http://localhost:8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func addTicker(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodGet, "https://api.binance.com/api/v3/ticker/price?symbols=[BTCUSDT]", nil)
	if err != nil {
		log.Fatal(err)
	}
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	var symbRate rate
	err = json.NewDecoder(res.Body).Decode(&symbRate)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	println(symbRate[0].Symbol)
	println(symbRate[0].Price)
}

func fetchTicker(w http.ResponseWriter, r *http.Request) {

}

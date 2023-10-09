package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"parser/models"
)

func FetchBinance(ticker *models.Ticker) error {
	var responseTicker []models.Received
	query := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbols=[\"%s\"]", ticker.Name) //to config
	response, err := http.Get(query)
	err = json.NewDecoder(response.Body).Decode(&responseTicker)
	ticker.Rate = responseTicker[0].Price
	ticker.Name = responseTicker[0].Symbol
	return err
}

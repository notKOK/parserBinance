package http

import (
	"github.com/gofiber/fiber/v2"
	"parser/models"
	"parser/useCase"
)

var defaultPath = "https://api.binance.com/api/v3/ticker/price?symbols=" + `["BTCUSDT"]`

type handler struct {
}

func (*handler) addTicker(c *fiber.Ctx) error {
	var ticker models.Ticker
	err := c.BodyParser(&ticker)
	if err != nil {
		return err
	}
	err = c.JSON(ticker)
	useCase.AddTicker()
	if err != nil {
		return err
	}
	return nil
	/*	var newRate rate
		err := json.Unmarshal(body, &newRate)
		println(newRate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"err": err,
			})
		}
		return c.Status(statusCode).JSON(newRate)

		if err := c.BodyParser(newRate); err != nil {
			return err
		}
		err = c.SendString(Agent.Name)
		if err != nil {
			return err
		}
		return nil*/
}

func (*handler) fetchTicker(c *fiber.Ctx) error {

	return nil
}

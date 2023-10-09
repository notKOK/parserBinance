package http

import (
	"github.com/gofiber/fiber/v2"
	"parser/internal/useCase"
	"parser/models"
)

type handler struct {
	//interface usecase
}

func (*handler) addTicker(c *fiber.Ctx) error {
	var ticker models.Ticker
	err := c.BodyParser(&ticker)
	if err != nil {
		return err
	}
	err = useCase.AddTicker(&ticker)
	err = c.JSON(&ticker)
	if err != nil {
		return err
	}
	go useCase.RunUpdate(&ticker)
	return nil
}

func (*handler) fetchTicker(c *fiber.Ctx) error {
	var ticker models.Ticker
	ticker.Name = c.Params("name")

	err := useCase.FetchTicker(&ticker)
	if err != nil {
		return err
	}
	err = c.JSON(&ticker)
	if err != nil {
		return err
	}
	return nil
}

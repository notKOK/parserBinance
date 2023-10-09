package http

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	handl handler
	app   *fiber.App
}

func New() (*Router, error) {
	resultRouter := Router{}
	resultRouter.app = fiber.New()
	resultRouter.app.Post("/addticker", resultRouter.handl.addTicker)
	resultRouter.app.Get("/fetchticker/:name", resultRouter.handl.fetchTicker)
	return &resultRouter, nil
}

func (Router *Router) Run() error {
	err := Router.app.Listen("0.0.0.0:8000")
	if err != nil {
		return err
	}
	return nil
}

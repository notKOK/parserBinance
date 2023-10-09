package httpServer

import (
	"parser/internal/http"
)

type App struct {
	router *http.Router
}

func New() (*App, error) {
	app := App{}
	var err error
	app.router, err = http.New()
	return &app, err
}

func (app *App) Run() error {
	err := app.router.Run()
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"github.com/gofiber/fiber/v2/log"
	"parser/httpServer"
)

func main() {
	application, err := httpServer.New()
	if err != nil {
		log.Fatal(err)
	}

	err = application.Run()
	if err != nil {
		log.Fatal(err)
	}
}

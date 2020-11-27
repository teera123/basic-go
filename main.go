package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/labstack/echo"

	"academy/internal/app"
	"academy/internal/handler"
)

func main() {
	state := flag.String("state", "dev", "environment")
	flag.Parse()

	cv := app.NewConfig()
	if err := cv.Init(*state); err != nil {
		log.Println("unable to init configuration", err)
		return
	}

	e := echo.New()
	if err := handler.NewRoutes(e, cv); err != nil {
		log.Println("new routes error", err)
		return
	}
	e.Start(fmt.Sprintf(":%d", cv.Port))
}

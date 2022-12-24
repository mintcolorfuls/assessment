package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

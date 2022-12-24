package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/mintcolorfuls/assessment/handlers/expenses"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(initialTable)

	e.POST("/expenses", expenses.Create)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

func initialTable(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := os.Getenv("DATABASE_URL")
		db, err := sql.Open("postgres", url)
		if err != nil {
			c.Error(err)
			return err
		}

		defer db.Close()

		query := `CREATE TABLE IF NOT EXISTS expenses ( id SERIAL PRIMARY KEY, title TEXT, amount FLOAT, note TEXT, tags TEXT[] );`
		_, err = db.Exec(query)
		if err != nil {
			c.Error(err)
			return err
		}

		log.Println("Create table successful")
		return nil
	}
}

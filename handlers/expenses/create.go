package expenses

import (
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	return c.JSON(200, nil)
}

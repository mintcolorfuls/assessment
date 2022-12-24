package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mintcolorfuls/assessment/packages/expenses"
)

func Create(c echo.Context) error {
	var expense expenses.Expense
	err := c.Bind(&expense)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if expense.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, "Amount should be greater than 0")
	}

	newExpense := expenses.Insert(expense.Title, expense.Note, expense.Amount, expense.Tags)
	return c.JSON(http.StatusCreated, newExpense)
}

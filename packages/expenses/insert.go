package expenses

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

func Insert(title, note string, amount int, tags []string) Expense {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connection Failure")
	}

	defer db.Close()

	query := `INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING id, title, amount, note, tags`
	row := db.QueryRow(query, title, amount, note, pq.Array(tags))

	var expense Expense

	err = row.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		log.Fatal("Can't insert data", err)
	}

	return expense
}

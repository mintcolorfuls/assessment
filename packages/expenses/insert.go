package expenses

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Insert() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connection Failure")
	}

	defer db.Close()
}

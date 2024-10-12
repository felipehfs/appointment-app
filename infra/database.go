package infra

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func CreateDatabase() (*sql.DB, error) {
	return sql.Open("postgres", "host=localhost port=5432 user=postgres password=123456  sslmode=disable")
}

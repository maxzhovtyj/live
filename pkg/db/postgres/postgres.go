package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewConn() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "postgres://postgres:1111@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return conn, nil
}

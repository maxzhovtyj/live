package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewConn() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "postgres://max:1111@localhost:5432/live?sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

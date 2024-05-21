package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewConn(connURI string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", connURI)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

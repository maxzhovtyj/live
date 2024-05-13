package postgres

import (
	"fmt"
	"testing"
)

func Test_NewConn(t *testing.T) {
	conn, err := NewConn()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(conn.Ping())
}

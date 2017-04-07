package goflake

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {

	conn := fmt.Sprintf("snowflake://testing.snowflakecomputing.com")

	db, err := sql.Open("snowflake", conn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		t.Skipf("skipping test: %v", err)
	}
}

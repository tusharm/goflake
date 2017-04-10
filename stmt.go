package goflake

import "database/sql/driver"

type Stmt struct {
}

func (c *RestConnection) Prepare(query string) (driver.Stmt, error) {
	return nil, ErrNotImplemented
}

// Query executes a query that may return rows, such as a
// SELECT.
func (d *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	return nil, ErrNotImplemented
}

// Exec executes a query that doesn't return rows, such
// as an INSERT or UPDATE.
func (d *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, ErrNotImplemented
}

func (d *Stmt) Close() error {
	return ErrNotImplemented
}
func (d *Stmt) NumInput() int {
	return 0
}

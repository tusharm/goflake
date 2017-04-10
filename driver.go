package goflake

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"time"
)

// DefaultConfig is probably only for temporary working
var DefaultConfig = &Configuration{
	Protocol: "http",
}

// Driver implements database/sql/driver interface.
type Driver struct {
	// Defaults
	proto, laddr, raddr, user, passwd, db string
	timeout                               time.Duration

	initCmds []string
}

// Open creates a new connection. The uri needs to have the following syntax:
//   snowflake://yourcompany.snowflakecomputing.com:DBNAME/USER/PASSWD
func (d *Driver) Open(uri string) (driver.Conn, error) {
	log.Println(uri)
	return &RestConnection{}, ErrNotImplemented
}

// Close ...
func (d *Driver) Close() error {
	return ErrNotImplemented
}

func init() {
	sql.Register("snowflake", &Driver{})
}

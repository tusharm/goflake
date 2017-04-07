package goflake

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"log"
	"net"
	"time"
)

// DefaultConfig is probably only for temporary working
var DefaultConfig = &Configuration{
	Protocol: "http",
}

// Dialer can be used to dial connections to Snowflake. If Dialer returns (nil, nil)
// the hook is skipped and normal dialing proceeds. user and dbname are there
// only for logging.
type Dialer func(proto, laddr, raddr, user, dbname string, timeout time.Duration) (net.Conn, error)

// Driver implements database/sql/driver interface.
type Driver struct {
	// Defaults
	proto, laddr, raddr, user, passwd, db string
	timeout                               time.Duration
	dialer                                Dialer

	initCmds []string
}

// Open creates a new connection. The uri needs to have the following syntax:
//   snowflake://yourcompany.snowflakecomputing.com:DBNAME/USER/PASSWD
func (d *Driver) Open(uri string) (driver.Conn, error) {
	log.Println(uri)
	return nil, errors.New("Wrong protocol part of URI")
}

// Close ...
func (d *Driver) Close() error {
	return nil
}

func init() {
	sql.Register("snowflake", &Driver{})
}

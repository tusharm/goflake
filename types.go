package goflake

import (
	"net"
	"time"
)

// Dialer can be used to dial connections to MySQL. If Dialer returns (nil, nil)
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

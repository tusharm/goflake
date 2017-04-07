package goflake

import "time"

// Configuration is a per connection thing
type Configuration struct {
	Dsn      string
	User     string
	Password string
	Host     string
	Port     int
	Database string

	Protocol  string
	Warehouse string
	Region    string
	Account   string
	Schema    string
	Role      string

	Timeout time.Time

	Autocommit bool
}

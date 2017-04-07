package goflake

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"log"
)

// Open creates a new connection. The uri needs to have the following syntax:
//
//   snowflake://yourcompany.snowflakecompiuting.com:DBNAME/USER/PASSWD
func (d *Driver) Open(uri string) (driver.Conn, error) {
	log.Println(uri)
	return nil, errors.New("Wrong protocol part of URI")
}

func init() {
	sql.Register("snowflake", &Driver{})
}

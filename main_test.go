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

	/*
			cs.execute("SELECT current_version()")
		        one = cs.fetchone()
		        print one[0]
		        ctx.cursor().execute("USE warehouse CM_TEST_WH;")
		        cs.execute("copy into \"LOADER\".\"PUBLIC\".\"RAW_EVENT\" "
		                    "FROM '@\"LOADER\".\"PUBLIC\".\"STAGE_JSON_FROM_S3\"' "
		                    "FILE_FORMAT = '\"LOADER\".\"PUBLIC\".\"JSON_FORMAT\"' "
		                    "PURGE = TRUE;")
		        results = cs.fetchall()
		        for result in results:
		            print result
	*/
}

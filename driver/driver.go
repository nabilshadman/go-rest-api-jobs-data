package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/************
* DATABASE *
************/
// ConnectDB connects to a PostgreSQL database using standard libray's sql.DB type.
// Logs an error if the connection fails.
func ConnectDB() *sql.DB {
	// retrieve value of environment variable of postgres database
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	// connect to database
	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}

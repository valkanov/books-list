package driver

import (
	"books-list/helpers"
	"database/sql"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

// ConnectDB connects to a database
func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	helpers.LogFatal(err)

	db, err = sql.Open("postgres", pgURL)
	helpers.LogFatal(err)

	err = db.Ping()
	helpers.LogFatal(err)
	return db
}

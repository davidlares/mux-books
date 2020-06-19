package driver

import (
  "github.com/lib/pq"
  "database/sql"
  "log"
  "os"
)

// db
var db *sql.DB

// logging fatal errors
func logFatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

// connect
func ConnectDB() *sql.DB {
  // db connection
  pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
  logFatal(err)
  // connection string
  db, err = sql.Open("postgres", pgUrl)
  logFatal(err)
  // ping
  err = db.Ping()
  logFatal(err)
  // printing the DB stats
  log.Println(pgUrl)

  return db
}

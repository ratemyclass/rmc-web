package rmcpostgres

import (
	"database/sql"
	"fmt"

	// pq is just the driver, always interact with database/sql api
	_ "github.com/lib/pq"
)

// RMCPostgres contains the database handler as well as other things the database may need
type RMCPostgres struct {
	// The SQL handler for the database
	dbHandler *sql.DB

	dbName     string
	dbUsername string
	dbPassword string
	dbHost     string
	dbPort     uint16
	// Rest is TODO
}

// InitializeDB initializes the database's credentials and socket, so they can be used in the SetupClient() function.
// This does not set up a database connection
func (db *RMCPostgres) InitializeDB(dbName string, dbUsername string, dbPassword string, dbHost string, dbPort uint16) (err error) {
	db.dbName = dbName
	db.dbUsername = dbUsername
	db.dbPassword = dbPassword
	db.dbHost = dbHost
	db.dbPort = dbPort

	return
}

// SetupClient sets up the PostgreSQL database client
func (db *RMCPostgres) SetupClient() (err error) {

	// We set up the connection string, ssl mode is require because why would you not want it to be
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=require", db.dbUsername, db.dbPassword, db.dbName, db.dbHost, db.dbPort)
	if db.dbHandler, err = sql.Open("postgres", connStr); err != nil {
		return
	}

	return
}

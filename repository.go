package repository

import (
	"database/sql"

	// PostgresSQL
	_ "github.com/lib/pq"
)

// InitializeDatabase creates a database connection and initializes all Repositories
var OpenDatabaseConnection = func(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	return db
}

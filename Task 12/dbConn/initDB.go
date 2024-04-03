package dbconn

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	// Try to connect Postgres Database and get the sql DB
	conn, err := sql.Open("postgres", "postgresql://postgres:123@localhost:5432/bacancy?sslmode=disable")
	if err != nil {
		return nil, err
	}
	// Ping and Check the connection is Active
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database Connection Established")
	// Return the connection and error
	return conn, nil
}

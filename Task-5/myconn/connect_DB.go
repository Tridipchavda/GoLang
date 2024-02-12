package myconn

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Function to check error and panic
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Function to connect DB by port, username and password, dbname
func ConnectToDB(_port int, _user string, _password string, _dbname string) (db *sql.DB) {
	host := "localhost"
	port := _port
	user := _user
	password := _password
	dbname := _dbname

	// Construct the connection string
	connstr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connecting the postgres driver
	db, err := sql.Open("postgres", connstr)
	CheckErr(err)

	CheckErr(db.Ping())
	fmt.Println("Successfully comm database")

	return db
}

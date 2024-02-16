package dbconn

import (
	"crud_postgres/models"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgresDB() (*App, error) {
	// Load the ENV Fle for Credentials
	godotenv.Load()

	// Store the credentials in variables
	host := os.Getenv("DATABASE_HOST")
	port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")

	// Make the connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connect to the Postgres using GORM
	postgresDB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Make Table Book if not exist in database
	postgresDB.AutoMigrate(&models.Book{})

	// Handling the error in DB connnection and ping
	db, err := postgresDB.DB()

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// return APP struct wuth DB
	return &App{DB: postgresDB}, nil
}

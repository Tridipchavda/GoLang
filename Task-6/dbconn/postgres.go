package dbconn

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Tridipchavda/middleware/models"
	"github.com/Tridipchavda/middleware/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initalize the database connection on postgres
func InitDB() (*gorm.DB, error) {
	// Load the environment variables
	m, err := utils.LoadEnv()
	if err != nil {
		return nil, err
	}
	// store the environment variables
	host := m["DATABASE_HOST"]
	port, _ := strconv.Atoi(m["DATABASE_PORT"])
	user := m["DATABASE_USER"]
	password := m["DATABASE_PASSWORD"]
	dbname := m["DATABASE_NAME"]

	// Make the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open Database connection in gorm with postgres driver
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Create Database if not present
	db.AutoMigrate(models.User{})
	log.Println("Connecting to postgres Successfull")

	//return DB and Nil in error
	return db, nil
}

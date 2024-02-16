package dbconn

import (
	"log"

	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func NewPostgresDBConn() (*App, error) {
	// connect Postgres DB with handling error
	a, err := ConnectPostgresDB()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) ClosePostgresConn() error {
	// Closing the connection in Postgres
	log.Println("Closing Connection to database")

	postdb, err := a.DB.DB()
	if err != nil {
		return err
	}
	err = postdb.Close()
	if err != nil {
		return err
	}
	return nil
}

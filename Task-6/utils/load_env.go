package utils

import (
	"os"

	"github.com/joho/godotenv"
)

// Function to load the environment variable
func LoadEnv() (map[string]string, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Store the environment variable in map and return it
	m := map[string]string{}

	m["DATABASE_USER"] = os.Getenv("DATABASE_USER")
	m["DATABASE_PORT"] = os.Getenv("DATABASE_PORT")
	m["DATABASE_NAME"] = os.Getenv("DATABASE_NAME")
	m["DATABASE_HOST"] = os.Getenv("DATABASE_HOST")
	m["DATABASE_PASSWORD"] = os.Getenv("DATABASE_PASSWORD")

	return m, nil
}

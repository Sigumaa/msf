package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	key, username, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

}

func loadConfig() (string, string, error) {
	if err := godotenv.Load(); err != nil {
		return "", "", err
	}
	key := os.Getenv("API_KEY")
	username := os.Getenv("USER_NAME")
	if key == "" {
		return "", "", errors.New("API_KEY is not set")
	}
	if username == "" {
		return "", "", errors.New("USER_NAME is not set")
	}
	return key, username, nil
}

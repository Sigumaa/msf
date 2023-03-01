package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/Sigumaa/lfu"
	"github.com/joho/godotenv"
)

func main() {
	key, username, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	client := lfu.New(username, key)

	// Get Friends List
	friends, err := client.Friends(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(friends); err != nil {
		log.Fatal(err)
	}
	log.Println(buf.String())
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

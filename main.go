package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/Sigumaa/lfu"
	"github.com/joho/godotenv"
)

func main() {
	key, username, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	client := lfu.New(username, key)
	log.Println("---friends---")
	if err := friends(client); err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	log.Println("---info---")
	if err := info(client); err != nil {
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

func printJSON(v interface{}) error {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		return err
	}
	log.Println(buf.String())
	return nil
}

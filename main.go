package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	key := os.Getenv("API_KEY")
	if key == "" {
		log.Fatal("API_KEY is not set")
	}
	fmt.Println(key)
}

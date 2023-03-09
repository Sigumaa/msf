package main

import (
	"context"
	"errors"
	"fmt"
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

	nowPlaying, err := client.NowPlayingTrack(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Now playing: %s - %s", nowPlaying.Name, nowPlaying.ArtistName())
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

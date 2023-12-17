package main

import (
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/rep-co/hola-quiz-backend/hola-quiz-api/config"
)

var cfg *config.Config

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg = &config.Config{}
	err = env.Parse(cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	storage, err := NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}
	if err := storage.Init(); err != nil {
		log.Fatal(err)
	}

	client = &http.Client{Timeout: 10 * time.Second}

	server := NewAPIServer(":8080", storage)
	server.Run()
}

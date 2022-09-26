package main

import (
	"log"
	"os"

	env "github.com/joho/godotenv"

	"github.com/FaridehGhani/ompfinex_challenge/api"
	"github.com/FaridehGhani/ompfinex_challenge/infra/mongodb"
)

func init() {
	// load env variable
	if err := env.Load(); err != nil {
		log.Fatalf("error loading env file: %v", err)
	}

	// load mongodb
	mongodb.NewClient()
}

func main() {
	// run api server
	router := api.ASCIIArtRouter()
	err := router.Run(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("http server running error: %v", err)
	}
}

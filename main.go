package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/noazaj/crypto-bot/config"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get .env variables
	apiKey := os.Getenv("ApiKey")
	apiSec := os.Getenv("ApiSec")

	// Pass .env variables to config
	config.SetConfig(apiKey, apiSec)

	fmt.Println("Main function running...")
}

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

	// Pass .env variables to config
	config.SetConfig(os.Getenv("KRAKEN_ID"), os.Getenv("KRAKEN_SECRET"))

	// Get the balance response from Kraken
	balance, err := GetBalance()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(balance)
}

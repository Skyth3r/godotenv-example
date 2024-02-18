package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Loaded .env file successfully")
}

func main() {
	text := os.Getenv("TEXT")
	fmt.Printf("%s\n", text) // "Hello, World!
}

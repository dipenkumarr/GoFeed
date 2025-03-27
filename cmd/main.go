package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello")

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Warning: Error loading .env file:", err)
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT is missing in env")
	}

	fmt.Println(portStr)
}
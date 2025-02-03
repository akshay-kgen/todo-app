package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	// load envs
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	app := NewApp()

	app.Start()
}

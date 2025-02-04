package main

import (
	"fmt"

	"github.com/akshay-kgen/todo-app/config"
	"github.com/joho/godotenv"
)

func main() {

	// load envs
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// load configs
	config.Initialize()
	serverConfig := config.GetInstance()

	app := NewApp(serverConfig)

	app.Start()
}

package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// go get github.com/joho/godotenv

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	app := os.Getenv("APP_NAME")
	version := os.Getenv("APP_VER")
	mail := os.Getenv("APP_OWNER_MAIL")

	fmt.Println("Application name:", app)
	fmt.Println("Application version:", version)
	fmt.Println("Application owner:", mail)
}

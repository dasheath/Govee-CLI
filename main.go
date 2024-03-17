/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"govee/cmd"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd.Execute()
}

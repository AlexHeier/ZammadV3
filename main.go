package main

import (
	getgroups "ZammadV3/getGroups"
	getusers "ZammadV3/getUsers"
	"ZammadV3/global"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Loading the ENV variables
func init() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatalf("Error loading .env file: %v", envErr)
	}

	global.ZAMMAD_TOKEN = os.Getenv("ZAMMMAD_TOKEN")
	if global.ZAMMAD_TOKEN == "" {
		log.Fatal("ZAMMAD_TOKEN is not set in the .env file")
	}
}

func main() {
	_ = getgroups.GetGroup()
	_ = getusers.GetUsers()
	global.ClearScreen()
}

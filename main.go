package main

import (
	getgroups "ZammadV3/getGroups"
	getusers "ZammadV3/getUsers"
	"ZammadV3/global"
	terminaloptions "ZammadV3/terminalOptions"
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
	global.IsLoading = true
	go global.LoadingScreen()
	groups := getgroups.GetGroup()
	users := getusers.GetUsers()
	global.IsLoading = false
	terminaloptions.Terminaloptions(groups, users)

}

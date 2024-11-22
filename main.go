package main

import (
	getgroups "ZammadV3/getGroups"
	getusers "ZammadV3/getUsers"
	"ZammadV3/global"
	terminaloptions "ZammadV3/terminalOptions"
	"log"

	"github.com/joho/godotenv"
)

// Loading the ENV variables
func init() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatalf("Error loading .env file: %v", envErr)
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

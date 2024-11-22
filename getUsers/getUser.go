package getusers

import (
	"ZammadV3/global"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Base URL for Zammad users API
var zammadUserEndpoint string = "/api/v1/users"

// GetAllUsers retrieves all users, handling pagination
func GetUsers() []global.User {
	var users []global.User

	page := 1
	perPage := 100

	for {
		// Construct the paginated URL
		url := fmt.Sprintf("%s?page=%d&per_page=%d", global.ZammadBaseURL+zammadUserEndpoint, page, perPage)

		// Create HTTP client and request
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}

		// Add authorization header
		req.Header.Add("Authorization", "Bearer "+global.ZAMMAD_TOKEN)

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error making request: %v", err)
		}
		defer resp.Body.Close()

		// Check for successful response
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Failed to get users, status code: %d", resp.StatusCode)
		}

		// Read and parse the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		var pageUsers []global.User
		if err := json.Unmarshal(body, &pageUsers); err != nil {
			log.Fatalf("Error parsing JSON: %v", err)
		}

		// Break the loop if no users are returned
		if len(pageUsers) == 0 {
			break
		}

		// Add the users from the current page to the main list
		users = append(users, pageUsers...)

		// Move to the next page
		page++
	}

	return users
}

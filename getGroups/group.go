package getgroups

import (
	"ZammadV3/global"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var zammadGroupURL string = "https://zammad.login.no/api/v1/groups"

func GetGroup() []global.Group {
	// Create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", zammadGroupURL, nil)
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
		log.Fatalf("Failed to get groups, status code: %d", resp.StatusCode)
	}

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var groups []global.Group
	if err := json.Unmarshal(body, &groups); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	return groups

}

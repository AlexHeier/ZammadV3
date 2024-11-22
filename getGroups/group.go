package getgroups

import (
	"ZammadV3/global"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

var zammadGroupEndpoint string = "/api/v1/groups"

func GetGroup() []global.Group {
	// Create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", global.ZammadBaseURL+zammadGroupEndpoint, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Add authorization header
	req.Header.Add("Authorization", "Bearer "+os.Getenv("ZAMMMAD_TOKEN"))

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

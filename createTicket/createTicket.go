package createticket

import (
	"ZammadV3/global"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var zammadTicketEndpoint string = "/api/v1/tickets"

func CereateTicket(mailTitle string, mailText []string, mailGroup global.Group, mailOwner []global.User, companies []global.Company) (bool, int) {

	global.ClearScreen()

	emailContent := strings.Join(mailText, "\n")

	for i, comcompny := range companies {
		fmt.Printf("Sending: \r%s", strings.Repeat(" ", 10))
		fmt.Printf("\r(%d/%d)", i+1, len(companies))

		owner := mailOwner[i%len(mailOwner)]
		//customer := fmt.Sprintf("guess:%s", comcompny.Emails)

		payload := map[string]interface{}{
			"title":    mailTitle,
			"group_id": mailGroup.ID,
			"customer": comcompny.Emails,
			"owner_id": owner.ID,
			"article": map[string]interface{}{
				"subject":      mailTitle,
				"body":         emailContent,
				"type":         "email",
				"content_type": "text/html",
				"to":           comcompny.Emails,
				"from":         owner.Email,
				"sender":       "System",
			},
		}

		fmt.Print(payload)

		jsonData, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Error encoding JSON: %v\n", err)
			return false, i
		}

		req, err := http.NewRequest("POST", global.ZammadBaseURL+zammadTicketEndpoint, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			return false, i
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+os.Getenv("ZAMMMAD_TOKEN"))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			return false, i
		}

		defer resp.Body.Close()

		fmt.Print("\n\n")
		fmt.Print(resp)
		fmt.Print("\n\n")

		if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
			continue
		} else {
			companies = append(companies, comcompny)
		}
	}
	return true, 0
}

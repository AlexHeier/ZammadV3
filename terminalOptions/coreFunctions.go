package terminaloptions

import (
	"ZammadV3/global"
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SetEmailContent(oldEmailText []string) []string {

	global.ClearScreen()
	print("1) See current email content\n2) Set email content with file\n3) Type new email content\nQ) exit\n")
	var emailText []string
	for {
		fmt.Printf("\nChoise: ")
		reader := bufio.NewReader(os.Stdin)
		choise, _ := reader.ReadString('\n')
		choise = strings.ToLower(strings.TrimSpace(choise))
		switch choise {
		case "1":
			if len(oldEmailText) <= 0 && len(emailText) <= 0 {
				fmt.Print("There is no email content yet")
				continue
			}

			if len(emailText) > 0 {
				for _, lines := range emailText {
					fmt.Println(lines)
				}
			} else {
				for _, lines := range oldEmailText {
					fmt.Println(lines)
				}
			}
			continue

		case "2":
			fmt.Print("\nNew absolute path: ")
			csvPath, _ := reader.ReadString('\n')
			csvPath = strings.TrimSpace(csvPath)

			file, err := os.Open(csvPath)
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
				return emailText
			}
			defer file.Close()

			// Create a scanner to read each line
			scanner := bufio.NewScanner(file)

			// Read each line and append it to the slice
			for scanner.Scan() {
				emailText = append(emailText, scanner.Text())
			}

			// Check for errors during scanning
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading file: %v\n", err)
			}

			fmt.Printf("\nHere is the new email content:\n\n")
			for _, line := range emailText {
				fmt.Println(line)
			}
			continue

		case "3":

			reader := bufio.NewReader(os.Stdin)
			var inputLines []string

			fmt.Print("\nEnter your email content text here. Type 'DONE' on a new line when finished. \nType 'DONE' with no new lines to exit without saving:\n")

			for {
				fmt.Print("> ")
				line, _ := reader.ReadString('\n')
				line = strings.TrimSpace(line)

				if strings.ToUpper(line) == "DONE" {
					break
				}

				inputLines = append(inputLines, line)
			}
			if len(inputLines) == 0 {
				continue
			} else {
				emailText = inputLines
			}

		case "q":
			if len(emailText) <= 0 {
				return oldEmailText
			}
			return emailText

		default:
			continue
		}

	}
}

// Setter email titelen
func SetEmailTitle(oldTitle string) (newTitle string) {
	global.ClearScreen()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("The title of the email\n")
	fmt.Printf("\nCurrent: %v", oldTitle)
	fmt.Print("\nNew (Q to exit): ")
	newTitle, _ = reader.ReadString('\n')
	newTitle = strings.TrimSpace(newTitle)
	if strings.ToLower(newTitle) == "q" {
		return oldTitle
	}
	return newTitle
}

// pathen til csv filen med bedrifter
func SetCsvPath(oldCsvPath string, oldCompanies []global.Company) (csvPath string, companiesObject []global.Company) {
	global.ClearScreen()

	fmt.Print("\nCSV format: Emails, CC")
	fmt.Print("\nExample line: x@stud.ntnu.no, y@stud.ntnu.no z@stud.ntnu.no ")
	fmt.Print("\nThere has be 1 customer email and 0...n CC\n")
	fmt.Printf("\nEmails in current .CSV file: %v\n", len(oldCompanies))
	var companies []global.Company

	read := bufio.NewReader(os.Stdin)
	fmt.Printf("\nCurrent absolute path: %v", oldCsvPath)
	fmt.Print("\nNew absolute path (Q to exit): ")
	newCsvPath, _ := read.ReadString('\n')
	newCsvPath = strings.TrimSpace(newCsvPath)

	if strings.ToLower(newCsvPath) == "q" {
		return oldCsvPath, oldCompanies
	}

	file, err := os.Open(newCsvPath)
	if err != nil {
		return "Didn't find CSV path", nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return "Error reading CSV", nil
	}

	for i, record := range records {
		// Skip the header row
		if i == 0 {
			continue
		}

		CCList := strings.Split(record[1], " ")

		// Create a Company struct and populate it
		company := global.Company{
			Emails: record[0],
			CC:     CCList,
		}

		// Add the company to the list
		companies = append(companies, company)
	}
	return newCsvPath, companies
}

// setter email group basert på user input og funn fra Zammad
func SetEmailGroup(oldMailGroup global.Group, groups []global.Group) (newMailGroup global.Group, change bool) {
	reader := bufio.NewReader(os.Stdin)

	for {
		global.ClearScreen()
		fmt.Print("Set e-mail group")
		fmt.Print("\nq to quit")
		fmt.Printf("\n\nCurrent: %s\n", newMailGroup.Name)
		for i, group := range groups {
			fmt.Printf("\n%v) %s", i+1, group.Name)
		}

		fmt.Print("\n\nChoise: ")
		choise, _ := reader.ReadString('\n')
		choise = strings.TrimSpace(choise)

		if strings.ToLower(choise) == "q" {
			break
		}

		number, err := strconv.Atoi(choise)
		if err != nil {
			continue
		}
		if number-1 > len(groups) || number <= 0 {
			continue
		}

		newMailGroup = groups[number-1]

	}
	if oldMailGroup == newMailGroup {
		return newMailGroup, false
	}
	return newMailGroup, true
}

func SetMailOwner(currentMailOwners, users []global.User, mailGroup global.Group) []global.User {
	var groupUser []global.User
	reader := bufio.NewReader(os.Stdin)

	global.IsLoading = true
	go global.LoadingScreen()

	for _, user := range users {
		if strings.Contains(strings.ToLower(user.Department), strings.ToLower(mailGroup.Name)) {
			groupUser = append(groupUser, user)
		}
	}
	global.IsLoading = false

	for {
		global.ClearScreen()

		fmt.Print("\nChose email owner(s)")
		fmt.Print("\nCurrent users:")
		for _, currentUser := range currentMailOwners {
			fmt.Printf("\n%s %s", currentUser.Firstname, currentUser.Lastname)
		}

		fmt.Print("\n\n1) Add owner(s)\n2) Remove owner(s)\nQ) exit\n\nChoise: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		choice = strings.ToLower(choice)

		switch choice {
		case "1":
			for {
				var avalibleUsers []global.User
				currentMailOwnersMap := make(map[global.User]bool)
				for _, user := range currentMailOwners {
					currentMailOwnersMap[user] = true
				}
				for _, user := range groupUser {
					if !currentMailOwnersMap[user] {
						avalibleUsers = append(avalibleUsers, user)
					}
				}
				fmt.Print("Avalibale user:\n")
				for i, user := range avalibleUsers {
					fmt.Printf("\n%v) %s %s", i+1, user.Firstname, user.Lastname)
				}
				fmt.Print("\n\nChoise: ")
				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(choice)
				choice = strings.ToLower(choice)

				if choice == "q" {
					break
				}

				number, _ := strconv.Atoi(choice)
				if number-1 > len(avalibleUsers) || number <= 0 {
					continue
				}
				currentMailOwners = append(currentMailOwners, avalibleUsers[number-1])
			}
		case "2":
			for {
				global.ClearScreen()
				fmt.Print("\nQ to quit")
				fmt.Print("\nWhich do you want to remove:\n")

				for i, user := range currentMailOwners {
					fmt.Printf("\n%v) %s %s", i+1, user.Firstname, user.Lastname)
				}

				fmt.Print("\nChoise: ")
				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(choice)

				if choice == "q" || choice == "Q" {
					break
				}

				number, _ := strconv.Atoi(choice)

				if number-1 > len(currentMailOwners) || number <= 0 {
					continue
				}
				for i, user := range currentMailOwners {
					var temp []global.User
					if i != number-1 {
						temp = append(temp, user)
					}
					currentMailOwners = temp
				}
			}
		case "q":
			return currentMailOwners
		default:
			continue
		}
	}
}

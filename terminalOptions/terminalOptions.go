package terminaloptions

import (
	createticket "ZammadV3/createTicket"
	"ZammadV3/global"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Terminaloptions(groups []global.Group, users []global.User) {
	var mailTitle string
	var mailGroup global.Group
	var mailText []string
	var mailCustomerPath string
	var companies []global.Company
	var mailOwner []global.User

	reader := bufio.NewReader(os.Stdin)

	for {
		emailOptions(mailGroup.Name, mailTitle, companies, mailOwner, len(mailText) > 0)

		answer, err := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		switch answer {
		case "1":
			change := false
			mailGroup, change = SetEmailGroup(mailGroup, groups)
			if change {
				mailOwner = nil
			}
		case "2":
			if mailGroup.Name == "" {
				fmt.Print("\nYou need to set a group first")
				reader.ReadString('\n')
				break
			}
			mailOwner = SetMailOwner(mailOwner, users, mailGroup)
		case "3":
			mailTitle = SetEmailTitle(mailTitle)
		case "4":
			mailText = SetEmailContent(mailText)
		case "5":
			mailCustomerPath, companies = SetCsvPath(mailCustomerPath, companies)
		case "s":
			success, amount := createticket.CereateTicket(mailTitle, mailText, mailGroup, mailOwner, companies)

			if success {
				fmt.Printf("\nSendt %d emails...\nEnter to exit: ", len(companies))
				reader.ReadString('\n')
			} else {
				fmt.Printf("\nSomething went wrong. Could not finnish sending emails.\nSendt %d out of %d", amount, len(companies))
				fmt.Printf("\n\nEnter to continue: ")
				reader.ReadString('\n')
			}
		case "q":
			global.ClearScreen()
			fmt.Print("\nTakk for at du brukte Zammad V3 :)\n\n")
			os.Exit(0)
		default:
		}

	}

}

func emailOptions(mailGroup, mailTitle string, companies []global.Company, mailOwner []global.User, hasText bool) {
	global.ClearScreen()
	fmt.Printf("\n1) Set email group. Current:                                %s", defaultIfENotSet(mailGroup))
	fmt.Printf("\n2) Set email owner(s). Has to have group first. Count:      %v", len(mailOwner))
	fmt.Printf("\n3) Set email title. Current:                                %s", defaultIfENotSet(mailTitle))
	fmt.Printf("\n4) Set email content or path. Have:                         %s", strconv.FormatBool(hasText))
	fmt.Printf("\n5) Set email customer and CC file path. Current emails:     %v", len(companies))
	fmt.Printf("\nS) to send the mail(s)")
	fmt.Printf("\nQ) to quit the program")
	fmt.Print("\n\nChoice: ")
}

// Hjelpe funksjoner

func defaultIfENotSet(s string) string {
	if s == "" {
		return "Not Set"
	}
	return s
}

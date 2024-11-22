package terminaloptions

import (
	"ZammadV3/global"
	"bufio"
	"fmt"
	"os"
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
			mailOwner = SetMailOwner(mailOwner, users, mailGroup)
		case "2":
			change := false
			mailGroup, change = SetEmailGroup(mailGroup, groups)
			if change {
				mailOwner = nil
			}
		case "3":
			mailTitle = SetEmailTitle(mailTitle)
		case "4":
			mailText = SetEmailContent(mailText)
		case "5":
			mailCustomerPath, companies = SetCsvPath(mailCustomerPath, companies)
		case "s":
			// send
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
	fmt.Printf("\n1) Set email owner(s). Has to have group first. Current:    %v", len(mailOwner))
	fmt.Printf("\n2) Set email group. Current:                                %s", defaultIfENotSet(mailGroup))
	fmt.Printf("\n3) Set email title. Current:                                %s", defaultIfENotSet(mailTitle))
	fmt.Printf("\n4) Set email content or path. Have:                         %s", boolTrueFalseString(hasText))
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

func boolTrueFalseString(b bool) string {
	if b {
		return "True"
	}
	return "False"
}

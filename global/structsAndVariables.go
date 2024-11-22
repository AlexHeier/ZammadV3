package global

var ZAMMAD_TOKEN string
var ZammadBaseURL string = "https://zammad.login.no"

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID         int    `json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Department string `json:"department"` // må bruke contains når jeg skal hente ut
}

type Company struct {
	Emails string
	CC     []string
}

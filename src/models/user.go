package models

type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Mobile    string `json:"mob"`
}

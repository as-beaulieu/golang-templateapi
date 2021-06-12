package models

type Account struct {
	User
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Dob       string `json:"dob"`
}

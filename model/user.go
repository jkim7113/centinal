package model

type User struct {
	UUID     string
	Username string
	Email    string
	Pw       string
	Bio      string
	Date     string
	PFP      string
	Role     string
	Verified bool
}

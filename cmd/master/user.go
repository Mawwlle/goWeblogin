package main

type User struct {
	Id      int
	Email   string
	EncPass string
}

func NewUser() *User {
	return &User{
		Email:   "TEST123",
		EncPass: "USER",
	}
}

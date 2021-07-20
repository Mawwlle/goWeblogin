package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type User struct {
	Id            int64
	Email         string
	EncryptedPass string
}

func NewUser() *User {
	return &User{
		Email:         "aboba@amogus",
		EncryptedPass: "lolkek",
	}
}

var databaseUrl = "postgres://testuser:1111@localhost/testmooc?sslmode=disable"

// ErrorState Function checking error state and shutting down program if state is err
func ErrorState(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CheckDataBaseConn DataBase conn use lazy creating, cause this
//function ping db and return error if status unavailable
func CheckDataBaseConn(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := sql.Open("postgres", databaseUrl)
	ErrorState(err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			ErrorState(err)
		}
	}(db)

	err = CheckDataBaseConn(db)
	if err != nil {
		ErrorState(err)
	}

	u := NewUser()
	err = db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPass,
	).Scan(&u.Id)
	if err != nil {
		ErrorState(err)
	}
	fmt.Printf("User with %v added", u.Id)
}

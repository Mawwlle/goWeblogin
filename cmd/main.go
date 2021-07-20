package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var database_url = "postgres://testuser:testpassword@localhost/testmooc?sslmode=disable"

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
	db, err := sql.Open("postgres", database_url)
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
	// ... use db here
}
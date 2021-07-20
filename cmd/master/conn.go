package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func ConnDB(passwd, addr, app string) (*sql.DB, error) {
	if passwd == "" {
		return nil, errors.New("db password not specified")
	}

	connStr := fmt.Sprintf("user=postgres password=%s sslmode=disable host=%s  application_name=%s", passwd, addr, app)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, err
	}

	return db, nil
}

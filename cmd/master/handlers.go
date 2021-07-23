package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Adder ToDo add error handler for reboot page
func Adder(app App, u *User) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := FindInDb(app, u)
		if err != nil {
			logrus.Infof("User already exist!")
			return err
		}

		err = app.dbConn.QueryRow(context.Background(),
			"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
			u.Email,
			u.EncPass,
		).Scan(&u.Id)
		if err != nil {
			return err
		}

		logrus.Infof("INTO table users has been added user with %v id %s email and password", u.Id, u.Email)
		return nil
	}
}
func FindInDb(app App, u *User) error {
	err := app.dbConn.QueryRow(context.Background(),
		"SELECT id, email, encrypted_password FROM users WHERE id = $1",
		u.Id,
	).Scan(
		&u.Id,
		&u.Email,
		&u.EncPass,
	)
	if u.Id == 0 {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func Getter(app App, u *User) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := FindInDb(app, u); err != nil {
			return err
		}
		logrus.Infof("User's email with %v id is %s", u.Id, u.Email)
		return nil
	}
}

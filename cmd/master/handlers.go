package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

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

func Adder(app App) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := make(map[string]string)
		u := new(User)
		if err := c.Bind(&data); err != nil {
			return err
		}
		u.Email = data["email"]
		u.EncPass = data["encpass"]

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
		result := fmt.Sprintf("Into table users has been added user with %v id %s email and password", u.Id, u.Email)
		logrus.Infof(result)
		newData, err := json.MarshalIndent(u, "", "\t")
		fmt.Printf("%s", newData)
		if err != nil {
			return err
		}
		return c.String(http.StatusCreated, result)
	}
}

func Getter(app App) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(User)
		data := make(map[string]string)
		if err := c.Bind(&data); err != nil {
			return err
		}
		i, err := strconv.Atoi(data["ID"])
		if err != nil {
			return err
		}
		u.Id = i
		if err := FindInDb(app, u); err != nil {
			return err
		}
		logrus.Infof("User's email with %v id is %s", u.Id, u.Email)

		newData, err := json.MarshalIndent(u, "", "\t")
		result := fmt.Sprintf("Into table users has been added user with %v id %s email and password %s",
			u.Id,
			u.Email,
			newData)
		logrus.Infof(result)
		fmt.Printf("%s", newData)
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, result)
	}
}

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

// FindInDb looking for a user in the database by ID
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

// Adder parses json and adds user to database
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
			return c.JSON(http.StatusConflict, err)
		}

		err = app.dbConn.QueryRow(context.Background(),
			"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
			u.Email,
			u.EncPass,
		).Scan(&u.Id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		result := fmt.Sprintf(
			"Into table users has been added user with %v id %s email and password %s",
			u.Id,
			u.Email,
			u.EncPass,
		)

		logrus.Infof(result)
		out, err := json.MarshalIndent(u, "", "\t")
		fmt.Printf("%s", out)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, u)
	}
}

// Getter parses json and recieve user from database
func Getter(app App) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(User)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		u.Id = id

		if err := FindInDb(app, u); err != nil {
			return c.JSON(http.StatusConflict, err)
		}
		logrus.Infof("User's email with %v id is %s", u.Id, u.Email)

		out, err := json.MarshalIndent(u, "", "\t")
		result := fmt.Sprintf("Into table users has been added user with %v id %s email\n %s",
			u.Id,
			u.Email,
			out)

		logrus.Infof(result)
		fmt.Printf("%s", out)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, u)
	}
}

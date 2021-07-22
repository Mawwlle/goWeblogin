package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(url string) error {
	m, err := migrate.New("file:///app/sql/migrations", url)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		return err
	}

	return nil
}

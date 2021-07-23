package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

// App holds all active objects.
type App struct {
	cfg        Config
	dbConn     *pgxpool.Pool
	httpClient *http.Client
}

func NewApp(cfg *Config) (*App, error) {
	var (
		db  *pgxpool.Pool
		err error
	)

	connStr := fmt.Sprintf(
		"user=postgres password=%s sslmode=disable host='%s' application_name='go-weblogin'",
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDR"),
	)

	sleep := 10 * time.Second
	for retry := 10; ; retry-- {
		db, err = pgxpool.Connect(context.Background(), connStr)
		if err != nil {
			logrus.WithError(err).Warnf("can't connect to db=%s", connStr)
		} else {
			logrus.Infof("connected db=%s", connStr)
			break
		}

		time.Sleep(sleep)
	}

	logrus.Trace("create app")

	err = Migrate(cfg.PostgresURL)
	if err != nil {
		logrus.Info(err.Error())
	}

	app := &App{
		cfg:        *cfg,
		dbConn:     db,
		httpClient: http.DefaultClient,
	}

	return app, nil
}

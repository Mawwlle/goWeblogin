package main

import (
	"flag"
	"os"
	"strconv"
	"time"
)

var (
	postgresURL = flag.String("pgURL", "postgresql://postgres:"+os.Getenv("DB_PASSWORD")+"@"+
		os.Getenv("DB_ADDR")+"/postgres?sslmode=disable", "PostgreSQL URL")
)

const (
	shutdownHTTPServerTimeout = 2 * time.Second
	shutdownAppTimeout        = 2 * time.Second
)

type Settings struct {
	serverAddr string
	loggerMode bool
}

type Config struct {
	ServerHTTPTimeout time.Duration
	PostgresURL       string
}

func Setting() *Settings {
	b, err := strconv.ParseBool("LOGGING_MODE")
	if err != nil {
		return &Settings{
			loggerMode: true,
			serverAddr: os.Getenv("SERVER_ADDR"),
		}
	}

	return &Settings{
		loggerMode: b,
		serverAddr: os.Getenv("SERVER_ADDR"),
	}
}

func NewConfig() *Config {
	return &Config{
		PostgresURL:       *postgresURL,
		ServerHTTPTimeout: 10,
	}
}

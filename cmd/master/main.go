package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

	config := NewConfig()
	appSetting := Setting()
	application, err := NewApp(config)
	if err != nil {
		logrus.WithError(err).Panic("create app")
	}

	defer shutdownApp(application)

	server := &http.Server{
		Addr:    appSetting.serverAddr,
		Handler: e,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		logrus.Infof("start HTTP server: %v", appSetting.serverAddr)

		if err := server.ListenAndServe(); err != nil {
			logrus.WithError(err).Warn("server shutdown")
		}
	}()

	defer shutdownServer(server)
	<-quit
}

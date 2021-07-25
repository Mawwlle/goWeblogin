package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (a *App) Shutdown() error {
	logrus.Info("shutdown app")
	a.dbConn.Close()

	return nil
}

func shutdownServer(server *http.Server) {
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctxServer, cancel := context.WithTimeout(context.Background(), shutdownHTTPServerTimeout)
	defer cancel()
	logrus.Warning("shutting down server...")

	if err := server.Shutdown(ctxServer); err != nil {
		logrus.Warning("server forced to shutdown:", err)
	}
}

func shutdownApp(application *App) {
	_, cancel := context.WithTimeout(context.Background(), shutdownAppTimeout)
	defer cancel()
	logrus.Warning("shutting down app...")

	if err := application.Shutdown(); err != nil {
		logrus.Warning("app forced to shutdown:", err)
	}
}

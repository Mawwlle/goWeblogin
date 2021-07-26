package main

import "github.com/sirupsen/logrus"

func setupLog() {
	s := Setting()
	logrus.SetLevel(logrus.WarnLevel)

	if s.loggerMode {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Warning("debug mode")
	}
}

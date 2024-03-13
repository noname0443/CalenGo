package main

import (
	"backend/internal"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	initLogger(test)

	app := internal.NewApp(":3000")
	app.Run()
}

func initLogger(env Environment) {
	if env == production {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.InfoLevel)
	} else if env == test {
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		panic("unknown environment")
	}
	logrus.SetOutput(os.Stdout)
}

type Environment string

const (
	production Environment = "production"
	test       Environment = "test"
)

package main

import (
	"os"

	"github.com/noname0443/CalenGo/backend/internal"

	"github.com/sirupsen/logrus"
)

func main() {
	initLogger(test)

	app := internal.NewApp(":1516")
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

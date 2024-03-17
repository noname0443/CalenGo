package main

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/noname0443/CalenGo/backend/internal"

	"github.com/sirupsen/logrus"
)

func main() {
	initLogger(test)

	db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/sys")
	if err != nil {
		logrus.Fatal(err)
	}

	app := internal.NewApp(":1516", db)
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

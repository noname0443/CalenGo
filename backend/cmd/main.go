package main

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/noname0443/CalenGo/backend/internal"
	"github.com/noname0443/CalenGo/backend/utility"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	initLogger(test)

	var db *sqlx.DB
	var err error

	logrus.Info("Trying to connect to MySQL")
	err = utility.DoRetry(func() error {
		db, err = sqlx.Connect("mysql", "root:root@(mysql:3306)/sys")
		if err != nil {
			logrus.Error("Failed to connect to MySQL", err.Error())
		}
		return err
	})
	if err != nil {
		logrus.Fatal(err)
	}

	err = utility.InitSQL(db)
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

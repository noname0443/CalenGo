package utility

import (
	"github.com/jmoiron/sqlx"
	"math"
	"os"
	"time"

	"strings"
	"github.com/sirupsen/logrus"
)

const RETRY_COUNT = 8

func DoRetry(function func() error) error {
	var err error
	for i := 0; i < RETRY_COUNT; i++ {
		err = function()
		if err == nil {
			return nil
		}
		time.Sleep(time.Second * time.Duration(math.Pow(2.0, float64(i))))
	}
	logrus.Error("got while retry:", err.Error())
	return err
}

func InitSQL(db *sqlx.DB) error {
	filedata, err := os.ReadFile("../init-mysql.sql")
	if err != nil {
		return err
	}

	requests := strings.Split(string(filedata), ";")
	for _, request := range requests {
		if len(request) == 0 {
			continue
		}
		_, err = db.Exec(request)
		if err != nil {
			return err
		}
	}
	return nil
}
package utility

import (
	"math"
	"time"

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

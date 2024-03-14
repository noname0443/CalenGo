package tests

import (
	"backend/internal"
	"io"
	"math"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

const RETRY_COUNT = 5

func doRetry(function func() error) error {
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

func TestHelloWorld(t *testing.T) {
	app := internal.NewApp(":3000")
	go app.Run()

	err := doRetry(func() error {
		resp, err := http.Get("http://0.0.0.0:3000/")
		if err != nil {
			return err
		}

		buf := new(strings.Builder)
		_, err = io.Copy(buf, resp.Body)
		if err != nil {
			return err
		}

		logrus.Info(buf.String())
		if buf.String() != "Hello, World!" {
			t.Error("bad value for / request")
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

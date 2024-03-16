package tests

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
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

func Get(ip string) error {
	resp, err := http.Get(fmt.Sprintf("http://%s/", ip))
	if err != nil {
		return err
	}

	_, err = ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func checkServerIsWorking(ip string) error {
	return doRetry(func() error {
		resp, err := http.Get(fmt.Sprintf("http://%s/", ip))
		if err != nil {
			return err
		}

		_, err = ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return nil
	})
}

func ReadAll(reader io.ReadCloser) (string, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, reader)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

package tests

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/noname0443/CalenGo/backend/utility"
)

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
	return utility.DoRetry(func() error {
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

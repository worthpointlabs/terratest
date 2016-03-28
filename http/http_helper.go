package http_helper

import (
	"errors"
	"time"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
	"log"
)

func HttpGet(url string, logger *log.Logger) (int, string, error) {
	logger.Println("Making an HTTP GET call to URL", url)

	resp, err := http.Get(url)
	if err != nil {
		return -1, "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return -1, "", err
	}

	return resp.StatusCode, strings.TrimSpace(string(body)), nil
}

func HttpGetWithRetry(url string, expectedBody string, retries int, sleepBetweenRetries time.Duration, logger *log.Logger) error {
	for i := 0; i < retries; i++ {
		status, body, err := HttpGet(url, logger)

		if err == nil && status == 200 {
			logger.Println("Got 200 OK from URL", url)
			if body == expectedBody {
				logger.Println("Got expected body from URL", url, ":", body)
				return nil
			} else {
				logger.Println("Did not get expected body from URL", url, ". Expected:", expectedBody, ". Got:", body, ".")
			}
		}

		if err != nil {
			logger.Println("Got an error after making an HTTP get to URL", url, ":", err)
		} else if status != 200 {
			logger.Println("Got a non-200 response from URL", url, ":", status)
		}

		logger.Println("Will retry in", sleepBetweenRetries)
		time.Sleep(sleepBetweenRetries)
	}

	return errors.New("Did not get a 200 OK from URL " + url + " after " + strconv.Itoa(retries) + " retries.")
}

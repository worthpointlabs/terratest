// Package http_helper contains helpers to interact with deployed resources through HTTP.
package http_helper

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/testing"
)

type HttpDoOptions struct {
	Method    string
	Url       string
	Body      io.Reader
	Headers   map[string]string
	TlsConfig *tls.Config
	Timeout   int
}

type HttpGetOptions struct {
	Url       string
	TlsConfig *tls.Config
	Timeout   int
}

// HttpGet performs an HTTP GET, with an optional pointer to a custom TLS configuration, on the given URL and
// return the HTTP status code and body. If there's any error, fail the test.
func HttpGet(t testing.TestingT, options *HttpGetOptions) (int, string) {
	statusCode, body, err := HttpGetE(t, options)
	if err != nil {
		t.Fatal(err)
	}
	return statusCode, body
}

// HttpGetE performs an HTTP GET, with an optional pointer to a custom TLS configuration, on the given URL and
// return the HTTP status code, body, and any error.
func HttpGetE(t testing.TestingT, options *HttpGetOptions) (int, string, error) {
	logger.Logf(t, "Making an HTTP GET call to URL %s", options.Url)

	// Set HTTP client transport config
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = options.TlsConfig

	timeout := 10
	if options.Timeout > 0 {
		timeout = options.Timeout
	}

	client := http.Client{
		// By default, Go does not impose a timeout, so an HTTP connection attempt can hang for a LONG time.
		Timeout: time.Duration(timeout) * time.Second,
		// Include the previously created transport config
		Transport: tr,
	}

	resp, err := client.Get(options.Url)
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

// HttpGetWithValidation performs an HTTP GET on the given URL and verify that you get back the expected status code and body. If either
// doesn't match, fail the test.
func HttpGetWithValidation(t testing.TestingT, options *HttpGetOptions, expectedStatusCode int, expectedBody string) {
	err := HttpGetWithValidationE(t, options, expectedStatusCode, expectedBody)
	if err != nil {
		t.Fatal(err)
	}
}

// HttpGetWithValidationE performs an HTTP GET on the given URL and verify that you get back the expected status code and body. If either
// doesn't match, return an error.
func HttpGetWithValidationE(t testing.TestingT, options *HttpGetOptions, expectedStatusCode int, expectedBody string) error {
	return HttpGetWithCustomValidationE(t, options, func(statusCode int, body string) bool {
		return statusCode == expectedStatusCode && body == expectedBody
	})
}

// HttpGetWithCustomValidation performs an HTTP GET on the given URL and validate the returned status code and body using the given function.
func HttpGetWithCustomValidation(t testing.TestingT, options *HttpGetOptions, validateResponse func(int, string) bool) {
	err := HttpGetWithCustomValidationE(t, options, validateResponse)
	if err != nil {
		t.Fatal(err)
	}
}

// HttpGetWithCustomValidationE performs an HTTP GET on the given URL and validate the returned status code and body using the given function.
func HttpGetWithCustomValidationE(t testing.TestingT, options *HttpGetOptions, validateResponse func(int, string) bool) error {
	statusCode, body, err := HttpGetE(t, options)

	if err != nil {
		return err
	}

	if !validateResponse(statusCode, body) {
		return ValidationFunctionFailed{Url: options.Url, Status: statusCode, Body: body}
	}

	return nil
}

// HttpGetWithRetry repeatedly performs an HTTP GET on the given URL until the given status code and body are returned or until max
// retries has been exceeded.
func HttpGetWithRetry(t testing.TestingT, options *HttpGetOptions, expectedStatus int, expectedBody string, retries int, sleepBetweenRetries time.Duration) {
	err := HttpGetWithRetryE(t, options, expectedStatus, expectedBody, retries, sleepBetweenRetries)
	if err != nil {
		t.Fatal(err)
	}
}

// HttpGetWithRetryE repeatedly performs an HTTP GET on the given URL until the given status code and body are returned or until max
// retries has been exceeded.
func HttpGetWithRetryE(t testing.TestingT, options *HttpGetOptions, expectedStatus int, expectedBody string, retries int, sleepBetweenRetries time.Duration) error {
	_, err := retry.DoWithRetryE(t, fmt.Sprintf("HTTP GET to URL %s", options.Url), retries, sleepBetweenRetries, func() (string, error) {
		return "", HttpGetWithValidationE(t, options, expectedStatus, expectedBody)
	})

	return err
}

// HttpGetWithRetryWithCustomValidation repeatedly performs an HTTP GET on the given URL until the given validation function returns true or max retries
// has been exceeded.
func HttpGetWithRetryWithCustomValidation(t testing.TestingT, options *HttpGetOptions, retries int, sleepBetweenRetries time.Duration, validateResponse func(int, string) bool) {
	err := HttpGetWithRetryWithCustomValidationE(t, options, retries, sleepBetweenRetries, validateResponse)
	if err != nil {
		t.Fatal(err)
	}
}

// HttpGetWithRetryWithCustomValidationE repeatedly performs an HTTP GET on the given URL until the given validation function returns true or max retries
// has been exceeded.
func HttpGetWithRetryWithCustomValidationE(t testing.TestingT, options *HttpGetOptions, retries int, sleepBetweenRetries time.Duration, validateResponse func(int, string) bool) error {
	_, err := retry.DoWithRetryE(t, fmt.Sprintf("HTTP GET to URL %s", options.Url), retries, sleepBetweenRetries, func() (string, error) {
		return "", HttpGetWithCustomValidationE(t, options, validateResponse)
	})

	return err
}

// HTTPDo performs the given HTTP method on the given URL and return the HTTP status code and body.
// If there's any error, fail the test.
func HTTPDo(
	t testing.TestingT, options *HttpDoOptions,
) (int, string) {
	statusCode, respBody, err := HTTPDoE(t, options)
	if err != nil {
		t.Fatal(err)
	}
	return statusCode, respBody
}

// HTTPDoE performs the given HTTP method on the given URL and return the HTTP status code, body, and any error.
func HTTPDoE(
	t testing.TestingT, options *HttpDoOptions,
) (int, string, error) {
	logger.Logf(t, "Making an HTTP %s call to URL %s", options.Method, options.Url)

	tr := &http.Transport{
		TLSClientConfig: options.TlsConfig,
	}

	timeout := 10
	if options.Timeout > 0 {
		timeout = options.Timeout
	}

	client := http.Client{
		// By default, Go does not impose a timeout, so an HTTP connection attempt can hang for a LONG time.
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: tr,
	}

	req := newRequest(options.Method, options.Url, options.Body, options.Headers)
	resp, err := client.Do(req)
	if err != nil {
		return -1, "", err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return -1, "", err
	}

	return resp.StatusCode, strings.TrimSpace(string(respBody)), nil
}

// HTTPDoWithRetry repeatedly performs the given HTTP method on the given URL until the given status code and body are
// returned or until max retries has been exceeded.
// The function compares the expected status code against the received one and fails if they don't match.
func HTTPDoWithRetry(
	t testing.TestingT, options *HttpDoOptions, expectedStatus int,
	retries int, sleepBetweenRetries time.Duration,
) string {
	out, err := HTTPDoWithRetryE(t, options, expectedStatus, retries, sleepBetweenRetries)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// HTTPDoWithRetryE repeatedly performs the given HTTP method on the given URL until the given status code and body are
// returned or until max retries has been exceeded.
// The function compares the expected status code against the received one and fails if they don't match.
func HTTPDoWithRetryE(
	t testing.TestingT, options *HttpDoOptions, expectedStatus int,
	retries int, sleepBetweenRetries time.Duration,
) (string, error) {
	out, err := retry.DoWithRetryE(
		t, fmt.Sprintf("HTTP %s to URL %s", options.Method, options.Url), retries,
		sleepBetweenRetries, func() (string, error) {
			statusCode, out, err := HTTPDoE(t, options)
			if err != nil {
				return "", err
			}
			logger.Logf(t, "output: %v", out)
			if statusCode != expectedStatus {
				return "", ValidationFunctionFailed{Url: options.Url, Status: statusCode}
			}
			return out, nil
		})

	return out, err
}

// HTTPDoWithValidationRetry repeatedly performs the given HTTP method on the given URL until the given status code and
// body are returned or until max retries has been exceeded.
func HTTPDoWithValidationRetry(
	t testing.TestingT, options *HttpDoOptions, expectedStatus int,
	expectedBody string, retries int, sleepBetweenRetries time.Duration,
) {
	err := HTTPDoWithValidationRetryE(t, options, expectedStatus, expectedBody, retries, sleepBetweenRetries)
	if err != nil {
		t.Fatal(err)
	}
}

// HTTPDoWithValidationRetryE repeatedly performs the given HTTP method on the given URL until the given status code and
// body are returned or until max retries has been exceeded.
func HTTPDoWithValidationRetryE(
	t testing.TestingT, options *HttpDoOptions, expectedStatus int,
	expectedBody string, retries int, sleepBetweenRetries time.Duration,
) error {
	_, err := retry.DoWithRetryE(t, fmt.Sprintf("HTTP %s to URL %s", options.Method, options.Url), retries,
		sleepBetweenRetries, func() (string, error) {
			return "", HTTPDoWithValidationE(t, options, expectedStatus, expectedBody)
		})

	return err
}

// HTTPDoWithValidation performs the given HTTP method on the given URL and verify that you get back the expected status
// code and body. If either doesn't match, fail the test.
func HTTPDoWithValidation(t testing.TestingT, options *HttpDoOptions, expectedStatusCode int, expectedBody string) {
	err := HTTPDoWithValidationE(t, options, expectedStatusCode, expectedBody)
	if err != nil {
		t.Fatal(err)
	}
}

// HTTPDoWithValidationE performs the given HTTP method on the given URL and verify that you get back the expected status
// code and body. If either doesn't match, return an error.
func HTTPDoWithValidationE(t testing.TestingT, options *HttpDoOptions, expectedStatusCode int, expectedBody string) error {
	return HTTPDoWithCustomValidationE(t, options, func(statusCode int, body string) bool {
		return statusCode == expectedStatusCode && body == expectedBody
	})
}

// HTTPDoWithCustomValidation performs the given HTTP method on the given URL and validate the returned status code and
// body using the given function.
func HTTPDoWithCustomValidation(t testing.TestingT, options *HttpDoOptions, validateResponse func(int, string) bool) {
	err := HTTPDoWithCustomValidationE(t, options, validateResponse)
	if err != nil {
		t.Fatal(err)
	}
}

// HTTPDoWithCustomValidationE performs the given HTTP method on the given URL and validate the returned status code and
// body using the given function.
func HTTPDoWithCustomValidationE(t testing.TestingT, options *HttpDoOptions, validateResponse func(int, string) bool) error {
	statusCode, respBody, err := HTTPDoE(t, options)

	if err != nil {
		return err
	}

	if !validateResponse(statusCode, respBody) {
		return ValidationFunctionFailed{Url: options.Url, Status: statusCode, Body: respBody}
	}

	return nil
}

func newRequest(method string, url string, body io.Reader, headers map[string]string) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil
	}
	for k, v := range headers {
		switch k {
		case "Host":
			req.Host = v
		default:
			req.Header.Add(k, v)
		}
	}
	return req
}

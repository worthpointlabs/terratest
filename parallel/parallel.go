// This package contains utilities for running parts of a single test in parallel. Normally, you would just have
// multiple tests and let Go execute them in parallel by calling t.Parallel(), but for test cases with expensive set
// up and tear down procedures (e.g. a big integration test where you have to start up and shut down many different
// servers), you may want to have just a single test and to run a bunch of pieces within that test in parallel.
package parallel

import (
	"time"
	"log"
	"fmt"
	"errors"
)

// We use goroutines and channels to run parts things in parallel. There are two limitations to note when doing this:
//
// 1. Channels can only return a single value.
// 2. You can only call t.Fatal or t.Fail from the original goroutine used for the test.
//
// Therefore, this struct is useful for allowing the parallelized parts of the tests to return values and errors to the
// original goroutine via a channel.
type TestResult struct {
	Description 	string	// A description of the test. This should always be set.
	Value		string	// An optional value to return from the test that might be passed on to other tests.
	Err		error	// An error. This should only be set if the test failed.
}

func (result TestResult) WithValue(value string) TestResult {
	result.Value = value
	return result
}

func (result TestResult) WithError(err error) TestResult {
	result.Err = err
	return result
}

// A wrapper type for a function that returns a TestResult
type Test func() TestResult

const DEFAULT_MAX_WAIT_FOR_CHANNEL = time.Minute * 10

// Runs the given test function in a goroutine and reports its result via the given channel
func RunTestInParallel(test Test, resultChannel chan <-TestResult) {
	go func() { resultChannel <- test() }()
}

// Calls GetTestResultWithTimeout with a default time out of DEFAULT_MAX_WAIT_FOR_CHANNEL. See GetTestResultWithTimeout
// for more info.
func GetTestResult(resultChannel <-chan TestResult, channelName string, logger *log.Logger) TestResult {
	return GetTestResultWithTimeout(resultChannel, channelName, logger, DEFAULT_MAX_WAIT_FOR_CHANNEL)
}

// Wait on the given channel for up to maxTimeout and return the TestResult the channel returns. If maxTimeout is
// exceeded, return a TestResult with an error.
func GetTestResultWithTimeout(resultChannel <-chan TestResult, channelName string, logger *log.Logger, maxTimeout time.Duration) TestResult {
	logger.Printf("Waiting on a result from channel %s for up to %s", channelName, maxTimeout.String())
	select {
	case result := <-resultChannel:
		logger.Printf("Channel %s returned a result: '%s'", channelName, result)
		return result
	case <-time.After(maxTimeout):
		errorMsg := fmt.Sprintf("Exceeded max timeout of %s waiting for channel %s", maxTimeout.String(), channelName)
		logger.Println(errorMsg)
		return TestResult{Err: errors.New(errorMsg), Description: errorMsg}
	}

	errorMsg := "The code should not be able to get here. The select statement should always return or fail the test."
	return TestResult{Err: errors.New(errorMsg), Description: errorMsg}
}



package test_util

import (
	"net"
	"net/http"
	"fmt"
	"strconv"
	"sync/atomic"
)

// Run a dummy HTTP server on a unique port that will return the given text. Returns the Listener for the server, the
// port it's listening on, or an error if something went wrong while trying to start the listener. Make sure to call
// the Close() method on the Listener when you're done!
func RunDummyServer(text string) (net.Listener, int, error) {
	port := getNextPort()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, text)
	})

	listener, err := net.Listen("tcp", ":" + strconv.Itoa(port))

	if err == nil {
		go http.Serve(listener, nil)
	}

	return listener, port, err
}

// DO NOT ACCESS THIS VARIABLE DIRECTLY. See getNextPort() below.
var testServerPort int32 = 8080

// Since we run tests in parallel, we need to ensure that each test runs on a different port. This function returns a
// unique port by atomically incrementing the testServerPort variable.
func getNextPort() int {
	return int(atomic.AddInt32(&testServerPort, 1))
}

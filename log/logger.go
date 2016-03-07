// Common package so we can use a standardized logging format
package log

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "[terraform-test] ", log.LstdFlags)
}

// Common package so we can use a standardized logging format
package log

import (
	"fmt"
	"log"
	"os"
)

func NewLogger(name string) *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("%s ", name), log.LstdFlags)
}

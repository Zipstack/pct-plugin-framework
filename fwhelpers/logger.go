package fwhelpers

import (
	"log"
	"os"
)

// Disable prefix in log messages.
var logger = log.New(os.Stdout, "", 0)

func GetLogger() *log.Logger {
	return logger
}

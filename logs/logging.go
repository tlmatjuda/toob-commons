package logs

import (
	"log"
	"os"
)

// const contextLogFlags = log.Ldate | log.Ltime | log.Lshortfile
const contextLogFlags = log.Ldate | log.Ltime

var Info = getInfo()
var Error = getError()
var Warn = getWarn()

func getInfo() *log.Logger {
	return buildLogger("INFO ")
}

func getError() *log.Logger {
	return buildLogger("ERROR ")
}

func getWarn() *log.Logger {
	return buildLogger("WARN ")
}

// buildLogger
// Build a logger for us with the specified.
func buildLogger(logPrefix string) *log.Logger {
	logger := log.New(os.Stdout, logPrefix, contextLogFlags)
	return logger
}

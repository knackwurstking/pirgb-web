// Package log for debug, warn, ...
package log

import (
	"log"
	"os"
)

var (
	// Debug logger
	Debug = log.New(os.Stderr, "[DEBUG] ", log.Lshortfile)
	// Info logger
	Info = log.New(os.Stderr, "[ INFO] ", log.Lshortfile)
	// Warn logger
	Warn = log.New(os.Stderr, "[ WARN] ", log.Lshortfile)
	// Error logger
	Error = log.New(os.Stderr, "[ERROR] ", log.Lshortfile)
)

func init() {
	log.SetFlags(log.Lshortfile)
}

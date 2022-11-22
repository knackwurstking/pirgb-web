package log

import (
	"log"
	"os"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func init() {
	log.SetFlags(log.Lshortfile)

	Debug = log.New(os.Stderr, "[DEBUG] ", log.Lshortfile)
	Info = log.New(os.Stderr, "[INFO] ", log.Lshortfile)
	Warn = log.New(os.Stderr, "[WARN] ", log.Lshortfile)
	Error = log.New(os.Stderr, "[ERROR] ", log.Lshortfile)
}

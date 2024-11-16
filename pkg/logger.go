package pkg

import (
	"log"
	"os"
)

// Logger struct encapsulates the standard log.Logger
type Logger struct {
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

// NewLogger initializes and returns a Logger instance
func NewLogger() *Logger {
	return &Logger{
		Info:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Warn:  log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Global logger instance to use across the application
var Log = NewLogger()

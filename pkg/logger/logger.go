package logger

import (
	"log"
	"os"
)

type Logger struct {
	info  *log.Logger
	error *log.Logger
}

func (logger *Logger) Info(format string, v ...any) {
	logger.info.Printf(format, v...)
}

func (logger *Logger) Error(format string, v ...any) {
	logger.error.Printf(format, v...)
}

func (logger *Logger) Fatal(format string, v ...any) {
	logger.error.Fatalf(format, v...)
}

func New() *Logger {
	info := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	error := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		info:  info,
		error: error,
	}
}

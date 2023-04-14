package logger

import (
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func (logger *Logger) Debug(format string, v ...any) {
	logger.info.Printf(format, v...)
}

func (logger *Logger) Info(format string, v ...any) {
	logger.info.Printf(format, v...)
}

func (logger *Logger) Warn(format string, v ...any) {
	logger.warn.Printf(format, v...)
}

func (logger *Logger) Error(format string, v ...any) {
	logger.error.Printf(format, v...)
}

func (logger *Logger) Fatal(format string, v ...any) {
	logger.error.Fatalf(format, v...)
}

func New() *Logger {
	debug := log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime|log.Lshortfile)
	info := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	warn := log.New(os.Stdout, "WARNING\t", log.Ldate|log.Ltime|log.Lshortfile)
	error := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		debug: debug,
		info:  info,
		warn:  warn,
		error: error,
	}
}

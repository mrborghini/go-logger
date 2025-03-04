package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Logger struct {
	typeName string
}

// Initialize a new logger
func NewLogger(typeName string) *Logger {
	return &Logger{typeName: typeName}
}

// Log a message
func (l *Logger) log(loglevel logLevel, message string) {
	var shownMessage string = fmt.Sprintf("%s - %s - %s]: %s\033[m", loglevel, l.typeName, getCurrentTime(), message)
	fmt.Println(shownMessage)
}

// Log an info message
func (l *Logger) Info(message string) {
	l.log(info, message)
}

// Log a warning message
func (l *Logger) Warning(message string) {
	l.log(warning, message)
}

// Log an error message
func (l *Logger) Error(message string) {
	l.log(error, message)
}

// Log a debug message
func (l *Logger) Debug(message string) {
	if strings.ToLower(os.Getenv("DEBUG")) == "true" {
		l.log(debug, message)
	}
}

func getCurrentTime() string {
	dt := time.Now()
	return dt.Local().Format("2006-01-02 15:04:05")
}

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	writer io.Writer
}

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func NewLogger(writer io.Writer) *Logger {
	return &Logger{writer: writer}
}

func (l *Logger) log(level LogLevel, msg string, args ...interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	levelStr := []string{"DEBUG", "INFO", "WARN", "ERROR"}[level]
  logMsg := fmt.Sprintf("%.20s [%.6s] %s", timestamp, levelStr, msg)
	for i := 0; i < len(args); i += 2 {
    if i == 0 { logMsg += " {" }
    logMsg += fmt.Sprintf("%v: %v", args[i], args[i+1])
    if i == len(args) - 2 {
      logMsg += "}"
    } else {
      logMsg += ", "
    }
	}
  logMsg += "\n"
  l.writer.Write([]byte(logMsg))
}

func (l *Logger) Debug(msg string) {
	l.log(DebugLevel, msg)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.log(InfoLevel, msg, args...)
}

func (l *Logger) Warn(msg string) {
	l.log(WarnLevel, msg)
}

func (l *Logger) Error(msg string) {
	l.log(ErrorLevel, msg)
}

func main() {
	log := NewLogger(os.Stderr)
	log.Debug("this is a debug message")
	log.Info("this is a info message", "userID", 1, "name", "fang")
	log.Warn("this is a warn message")
	log.Error("this is a error message")
}

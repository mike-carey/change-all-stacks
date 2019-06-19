package change

import (
	"io"
	"os"
	"fmt"
)

//go:generate counterfeiter -o fakes/fake_logger.go Logger
type Logger interface {
	Info(msg string)
	Infof(msg string, args ...interface{})
	Debug(msg string)
	Debugf(msg string, args ...interface{})
}

// Simple logger
type logger struct {
	Verbose bool
}

func NewLogger(verbose bool) Logger {
	return &logger{
		Verbose: verbose,
	}
}

func (l *logger) log(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
}

func (l *logger) logf(w io.Writer, msg string, args ...interface{}) {
	fmt.Fprintf(w, msg + "\n", args...)
}

func (l *logger) Info(msg string) {
	l.log(os.Stdout, msg)
}

func (l *logger) Infof(msg string, args ...interface{}) {
	l.logf(os.Stdout, msg, args...)
}

func (l *logger) Debug(msg string) {
	if l.Verbose {
		l.log(os.Stderr, "[DEBUG] " + msg)
	}
}

func (l *logger) Debugf(msg string, args ...interface{}) {
	if l.Verbose {
		l.logf(os.Stderr, "[DEBUG] "+msg, args...)
	}
}

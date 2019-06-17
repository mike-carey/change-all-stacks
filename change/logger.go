package change

import (
	"io"
	"os"
	"fmt"
)

// Simple logger
type Logger struct {
	Verbose bool
}

func NewLogger(verbose bool) *Logger {
	return &Logger{
		Verbose: verbose,
	}
}

func (l *Logger) log(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
}

func (l *Logger) logf(w io.Writer, msg string, args ...interface{}) {
	fmt.Fprintf(w, msg + "\n", args...)
}

func (l *Logger) Info(msg string) {
	l.log(os.Stdout, msg)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.logf(os.Stdout, msg, args...)
}

func (l *Logger) Debug(msg string) {
	if l.Verbose {
		l.log(os.Stderr, "[DEBUG] " + msg)
	}
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	if l.Verbose {
		l.logf(os.Stderr, "[DEBUG] "+msg, args...)
	}
}

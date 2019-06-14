package logger

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/google/logger"
)

var verbose bool
var instance *logger.Logger
var once sync.Once

func getInstance() *logger.Logger {
	once.Do(func() {
		instance = logger.Init("", getVerbose(), false, ioutil.Discard)
		logger.SetFlags(log.LstdFlags)
	})

	return instance
}

// SetVerbose sets the verbose flag before the logger is initialized
func SetVerbose() {
	verbose = true
}

func getVerbose() bool {
	if verbose {
		return true
	}

	v := os.Getenv("CFQUERY_VERBOSE")
	if v == "" {
		v = "false"
	}
	verbose, err := strconv.ParseBool(v)
	if err != nil {
		panic(err)
	}

	return verbose
}

// Close ...
func Close() {
	getInstance().Close()
}

// Info logs with the Info severity.
// Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) {
	getInstance().Info(v...)
}

// InfoDepth acts as Info but uses depth to determine which call frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func InfoDepth(depth int, v ...interface{}) {
	getInstance().InfoDepth(depth, v...)
}

// Infoln logs with the Info severity.
// Arguments are handled in the manner of fmt.Println.
func Infoln(v ...interface{}) {
	getInstance().Infoln(v...)
}

// Infof logs with the Info severity.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	getInstance().Infof(format, v...)
}

// Warning logs with the Warning severity.
// Arguments are handled in the manner of fmt.Print.
func Warning(v ...interface{}) {
	getInstance().Warning(v...)
}

// WarningDepth acts as Warning but uses depth to determine which call frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func WarningDepth(depth int, v ...interface{}) {
	getInstance().WarningDepth(depth, v...)
}

// Warningln logs with the Warning severity.
// Arguments are handled in the manner of fmt.Println.
func Warningln(v ...interface{}) {
	getInstance().Warningln(v...)
}

// Warningf logs with the Warning severity.
// Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, v ...interface{}) {
	getInstance().Warningf(format, v...)
}

// Error logs with the ERROR severity.
// Arguments are handled in the manner of fmt.Print.
func Error(v ...interface{}) {
	getInstance().Error(v...)
}

// ErrorDepth acts as Error but uses depth to determine which call frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func ErrorDepth(depth int, v ...interface{}) {
	getInstance().ErrorDepth(depth, v...)
}

// Errorln logs with the ERROR severity.
// Arguments are handled in the manner of fmt.Println.
func Errorln(v ...interface{}) {
	getInstance().Errorln(v...)
}

// Errorf logs with the Error severity.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	getInstance().Errorf(format, v...)
}

// Fatal logs with the Fatal severity, and ends with os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func Fatal(v ...interface{}) {
	getInstance().Fatal(v...)
}

// FatalDepth acts as Fatal but uses depth to determine which call frame to log.
// FatalDepth(0, "msg") is the same as Fatal("msg").
func FatalDepth(depth int, v ...interface{}) {
	getInstance().FatalDepth(depth, v...)
}

// Fatalln logs with the Fatal severity, and ends with os.Exit(1).
// Arguments are handled in the manner of fmt.Println.
func Fatalln(v ...interface{}) {
	getInstance().Fatalln(v...)
}

// Fatalf logs with the Fatal severity, and ends with os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, v ...interface{}) {
	getInstance().Fatalf(format, v...)
}

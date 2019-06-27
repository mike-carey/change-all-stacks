package logger

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/google/logger"
)

var debugMode bool = false
var devMode *bool
var instance *logger.Logger
var once sync.Once

const (
	DevEnvVar = "DEVELOPMENT"
	VerboseEnvVar = "VERBOSE"
	DebugEnvVar = "DEBUG"
)

func GetInstance() *logger.Logger {
	once.Do(func() {
		if instance == nil {
			if !isDevMode() {
				panic("Cannot init unless in dev mode or Init called")
			}

			Init(nil, nil)
		}
	})

	return instance
}

func getBoolEnvVar(name string) bool {
	v := os.Getenv(name)
	if v == "" {
		v = "false"
	}

	d, err := strconv.ParseBool(v)
	if err != nil {
		panic(err)
	}

	return d

}

func isDebug() bool {
	return debugMode
}

func isDevMode() bool {
	if devMode == nil {
		d := getBoolEnvVar(DevEnvVar)
		devMode = &d
	}

	return *devMode
}

func Init(verbose *bool, debug *bool) {
	if isDevMode() {
		if verbose == nil {
			v := getBoolEnvVar(VerboseEnvVar)
			verbose = &v
		}

		if debug == nil {
			d := getBoolEnvVar(DebugEnvVar)
			debug = &d
		}
	}

	if verbose == nil || debug == nil {
		panic("Logger failed to initialize!")
	}

	debugMode = *debug

	instance = logger.Init("", *verbose, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
}

func Debug(msg string) {
	if isDebug() {
		GetInstance().Info("[DEBUG] " + msg)
	}
}

func Debugf(msg string, args ...interface{}) {
	if isDebug() {
		GetInstance().Infof("[DEBUG] " + msg, args...)
	}
}

// Close ...
func Close() {
	GetInstance().Close()
}

// Info logs with the Info severity.
// Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) {
	GetInstance().Info(v...)
}

// InfoDepth acts as Info but uses depth to determine which call frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func InfoDepth(depth int, v ...interface{}) {
	GetInstance().InfoDepth(depth, v...)
}

// Infoln logs with the Info severity.
// Arguments are handled in the manner of fmt.Println.
func Infoln(v ...interface{}) {
	GetInstance().Infoln(v...)
}

// Infof logs with the Info severity.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	GetInstance().Infof(format, v...)
}

// Warning logs with the Warning severity.
// Arguments are handled in the manner of fmt.Print.
func Warning(v ...interface{}) {
	GetInstance().Warning(v...)
}

// WarningDepth acts as Warning but uses depth to determine which call frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func WarningDepth(depth int, v ...interface{}) {
	GetInstance().WarningDepth(depth, v...)
}

// Warningln logs with the Warning severity.
// Arguments are handled in the manner of fmt.Println.
func Warningln(v ...interface{}) {
	GetInstance().Warningln(v...)
}

// Warningf logs with the Warning severity.
// Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, v ...interface{}) {
	GetInstance().Warningf(format, v...)
}

// Error logs with the ERROR severity.
// Arguments are handled in the manner of fmt.Print.
func Error(v ...interface{}) {
	GetInstance().Error(v...)
}

// ErrorDepth acts as Error but uses depth to determine which call frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func ErrorDepth(depth int, v ...interface{}) {
	GetInstance().ErrorDepth(depth, v...)
}

// Errorln logs with the ERROR severity.
// Arguments are handled in the manner of fmt.Println.
func Errorln(v ...interface{}) {
	GetInstance().Errorln(v...)
}

// Errorf logs with the Error severity.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	GetInstance().Errorf(format, v...)
}

// Fatal logs with the Fatal severity, and ends with os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func Fatal(v ...interface{}) {
	GetInstance().Fatal(v...)
}

// FatalDepth acts as Fatal but uses depth to determine which call frame to log.
// FatalDepth(0, "msg") is the same as Fatal("msg").
func FatalDepth(depth int, v ...interface{}) {
	GetInstance().FatalDepth(depth, v...)
}

// Fatalln logs with the Fatal severity, and ends with os.Exit(1).
// Arguments are handled in the manner of fmt.Println.
func Fatalln(v ...interface{}) {
	GetInstance().Fatalln(v...)
}

// Fatalf logs with the Fatal severity, and ends with os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, v ...interface{}) {
	GetInstance().Fatalf(format, v...)
}

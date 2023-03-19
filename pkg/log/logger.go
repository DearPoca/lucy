package log

import (
	"fmt"
	"os"
)

const (
	LTrace = iota
	LDebug
	LInfo
	LWarn
	LError
	LFatal
	LPanic
	logLevelCount
)

var logLevelStr = []string{
	"TRACE",
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
	"PANIC",
}

var cLogLevel int

func Trace(msg string, v ...any) {
	if cLogLevel <= LTrace {
		fmt.Println(serialize(2, LTrace, msg, v...))
	}
}

func Debug(msg string, v ...any) {
	if cLogLevel <= LDebug {
		fmt.Println(serialize(2, LDebug, msg, v...))
	}
}

func Info(msg string, v ...any) {
	if cLogLevel <= LInfo {
		fmt.Println(serialize(2, LInfo, msg, v...))
	}
}

func Warn(msg string, v ...any) {
	if cLogLevel <= LWarn {
		fmt.Println(serialize(2, LWarn, msg, v...))
	}
}

func Error(msg string, v ...any) {
	if cLogLevel <= LError {
		fmt.Println(serialize(2, LError, msg, v...))
	}
}

func Fatal(msg string, v ...any) {
	if cLogLevel <= LFatal {
		fmt.Println(serialize(2, LFatal, msg, v...))
		os.Exit(1)
	}
}

func Panic(msg string, v ...any) {
	if cLogLevel <= LPanic {
		s := serialize(2, LPanic, msg, v...)
		fmt.Println(s)
		panic(s)
	}
}

func SetLevel(l int) {
	cLogLevel = l
}

func SetLevelStr(l string) {
	for i := 0; i < logLevelCount; i++ {
		if l == logLevelStr[i] {
			cLogLevel = i
		}
	}
}

func init() {
	cLogLevel = LInfo
}

package log

import (
	"fmt"
	"log"
	"os"

	"lucy/pkg/setting"
)

const (
	trace = iota
	debug
	info
	warn
	error
	fatal
	panic
	logEnumSize
)

var logLevelStr = []string{
	"Trace",
	"Debug",
	"Info",
	"Warn",
	"Error",
	"Fatal",
	"Panic",
}

var Logger *log.Logger
var loggers []*log.Logger
var cLogLevel int

func Trace(msg string, v ...any) {
	if cLogLevel <= trace {
		loggers[trace].Println(serialize(msg, v...))
	}
}

func Debug(msg string, v ...any) {
	if cLogLevel <= debug {
		loggers[debug].Println(serialize(msg, v...))
	}
}

func Info(msg string, v ...any) {
	if cLogLevel <= info {
		loggers[info].Println(serialize(msg, v...))
	}
}

func Warn(msg string, v ...any) {
	if cLogLevel <= warn {
		loggers[warn].Println(serialize(msg, v...))
	}
}

func Error(msg string, v ...any) {
	if cLogLevel <= error {
		loggers[error].Println(serialize(msg, v...))
	}
}

func Fatal(msg string, v ...any) {
	if cLogLevel <= fatal {
		loggers[fatal].Fatalln(serialize(msg, v...))
	}
}

func Panic(msg string, v ...any) {
	if cLogLevel <= trace {
		loggers[trace].Panicln(serialize(msg, v...))
	}
}

func serialize(msg string, v ...any) string {
	format := msg
	var a []any
	for i := 0; i < len(v); i += 2 {
		if i+1 < len(v) {
			format += ", " + v[i].(string) + ": %+v"
			a = append(a, v[i+1])
		} else {
			format += ", " + v[i].(string)
		}
	}
	return fmt.Sprintf(format, a)
}

func init() {
	cLogLevel = info
	for i := 0; i < logEnumSize; i++ {
		loggers = append(loggers, log.New(os.Stdout, "["+logLevelStr[i]+"]", log.Lshortfile|log.Ldate|log.Ltime))
		if setting.AppSetting.LogLevel == logLevelStr[i] {
			cLogLevel = i
		}
	}
	Logger = loggers[cLogLevel]
}

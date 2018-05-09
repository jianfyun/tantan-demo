package log

import (
	"os"

	logging "github.com/op/go-logging"
)

var log *logging.Logger

func Init(module string) {
	log = logging.MustGetLogger(module)
	format := logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05.000} [%{level:.4s}] [%{shortfile} %{shortfunc}] %{color:reset} %{message}`,
	)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)
	leveled := logging.AddModuleLevel(formatter)
	leveled.SetLevel(logging.DEBUG, "")
	logging.SetBackend(leveled)
}

// Debug write one line of debug level log with the providing arguments.
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Debugf write one line of debug level log with the providing format and arguments.
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Info write one line of info level log with the providing arguments.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Infof write one line of info level log with the providing format and arguments.
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Notice write one line of notice level log with the providing arguments.
func Notice(args ...interface{}) {
	log.Notice(args...)
}

// Noticef write one line of notice level log with the providing format and arguments.
func Noticef(format string, args ...interface{}) {
	log.Noticef(format, args...)
}

// Warning write one line of warning level log with the providing arguments.
func Warning(args ...interface{}) {
	log.Warning(args...)
}

// Warningf write one line of warning level log with the providing format and arguments.
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

// Error write one line of error level log with the providing arguments.
func Error(args ...interface{}) {
	log.Error(args...)
}

// Errorf write one line of error level log with the providing format and arguments.
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Critical write one line of critical level log with the providing arguments.
func Critical(args ...interface{}) {
	log.Critical(args...)
}

// Criticalf write one line of critical level log with the providing format and arguments.
func Criticalf(format string, args ...interface{}) {
	log.Criticalf(format, args...)
}

// Fatal write one line of critical level log with the providing arguments, and then exit the program.
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf write one line of critical level log with the providing format and arguments, and then exit the program.
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

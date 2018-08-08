// Package log is an alternative to log package in standard library.
package log

import "os"

type Level int

// Logging levels.
const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

func (l Level) String() string {
	names := map[Level]string{
		CRITICAL: "CRITICAL",
		ERROR:    "ERROR",
		WARNING:  "WARNING",
		NOTICE:   "NOTICE",
		INFO:     "INFO",
		DEBUG:    "DEBUG",
	}
	s, ok := names[l]
	if !ok {
		return "UNKNOWN"
	}
	return s
}

var (
	DefaultLogger    Logger    = NewLogger(procName)
	DefaultLevel     Level     = INFO
	DefaultHandler   Handler   = NewFileHandler(os.Stderr)
	DefaultFormatter Formatter = defaultFormatter{}
)

func init() {
	DefaultLogger.SetCallDepth(1)
}

///////////////////
//               //
// DefaultLogger //
//               //
///////////////////

// SetLevel changes the level of DefaultLogger and DefaultHandler.
func SetLevel(l Level) {
	DefaultLogger.SetLevel(l)
	DefaultHandler.SetLevel(l)
}

func Fatal(args ...interface{})                    { DefaultLogger.Fatal(args...) }
func Fatalf(format string, args ...interface{})    { DefaultLogger.Fatalf(format, args...) }
func Fatalln(args ...interface{})                  { DefaultLogger.Fatalln(args...) }
func Panic(args ...interface{})                    { DefaultLogger.Panic(args...) }
func Panicf(format string, args ...interface{})    { DefaultLogger.Panicf(format, args...) }
func Panicln(args ...interface{})                  { DefaultLogger.Panicln(args...) }
func Critical(args ...interface{})                 { DefaultLogger.Critical(args...) }
func Criticalf(format string, args ...interface{}) { DefaultLogger.Criticalf(format, args...) }
func Criticalln(args ...interface{})               { DefaultLogger.Criticalln(args...) }
func Error(args ...interface{})                    { DefaultLogger.Error(args...) }
func Errorf(format string, args ...interface{})    { DefaultLogger.Errorf(format, args...) }
func Errorln(args ...interface{})                  { DefaultLogger.Errorln(args...) }
func Warning(args ...interface{})                  { DefaultLogger.Warning(args...) }
func Warningf(format string, args ...interface{})  { DefaultLogger.Warningf(format, args...) }
func Warningln(args ...interface{})                { DefaultLogger.Warningln(args...) }
func Notice(args ...interface{})                   { DefaultLogger.Notice(args...) }
func Noticef(format string, args ...interface{})   { DefaultLogger.Noticef(format, args...) }
func Noticeln(args ...interface{})                 { DefaultLogger.Noticeln(args...) }
func Info(args ...interface{})                     { DefaultLogger.Info(args...) }
func Infof(format string, args ...interface{})     { DefaultLogger.Infof(format, args...) }
func Infoln(args ...interface{})                   { DefaultLogger.Infoln(args...) }
func Debug(args ...interface{})                    { DefaultLogger.Debug(args...) }
func Debugf(format string, args ...interface{})    { DefaultLogger.Debugf(format, args...) }
func Debugln(args ...interface{})                  { DefaultLogger.Debugln(args...) }

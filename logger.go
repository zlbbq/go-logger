package logger

import (
	"github.com/daviddengcn/go-colortext"
	"fmt"
	"time"
	"io"
	"os"
)

//LogLevel
type Level int

//Logger
type Logger struct {
	Name string
	Level Level
	Colorful bool
	Output io.Writer
}

//LogLevel constants
const (
	LevelDebug = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

//global logger
//
// The global logger is named "root" and it is a colorful logger with level DEBUG and log to os.Stdout
var gLogger = NewLogger(LevelDebug, true, "", nil)
var gLoggers = make(map[string]*Logger)

//debug log
func Debug(fmt string, v ...interface{}) {
	gLogger.Debug(fmt, v...)
}

//info log
func Info(fmt string, v ...interface{}) {
	gLogger.Info(fmt, v...)
}

//warning log
func Warn(fmt string, v ...interface{}) {
	gLogger.Warn(fmt, v...)
}

//error log
func Error(fmt string, v ...interface{}) {
	gLogger.Error(fmt, v...)
}

//fatal error log
func Fatal(fmt string, v ...interface{}) {
	gLogger.Fatal(fmt, v...)
}

//set log level of the global logger
func SetLevel(level Level) {
	gLogger.Level = level
}

//set colorful of the global logger
func SetColorful(b bool) {
	gLogger.Colorful = b
}

//set where to write the log text of the global logger
func SetOutput(output io.Writer) {
	gLogger.Output = output
}

//get a logger from logger pool, if cooresponding logger is not found, a simple logger is created and registered then return
//
//it is recommended that libraries call this function to initialize a library-inner-logger and pass library full name
//to the "name" argument
func Get(name string) *Logger {
	l := gLoggers[name]
	if(l == nil) {
		l = NewSimpleLogger(name)
		Register(l)
	}
	return l
}

func Register(l * Logger)  {
	gLoggers[l.Name] = l
}

//create a logger
func NewLogger(level Level, colorful bool, name string, output io.Writer) *Logger {
	if(output == nil) {
		output = os.Stdout
	}
	if(name == "") {
		name = "root"
	}
	return &Logger{name, level, colorful, output}
}

//create a named logger
//
//simple logger is a colorful logger with level DEBUG and log to os.Stdout
func NewSimpleLogger(name string) *Logger{
	return NewLogger(LevelDebug, true, name, nil)
}

//debug log
func (logger *Logger) Debug(fmt string, v ...interface{}) {
	if(logger.Level > LevelDebug) {
		return
	}
	if(logger.Colorful == true) {
		ct.ChangeColor(ct.Cyan, false, ct.None, false)
	}
	logger.logText(logger.Output, "TRACE:", fmt, v...)
}

//info log
func (logger *Logger) Info(fmt string, v ...interface{}) {
	if(logger.Level > LevelInfo) {
		return
	}
	if(logger.Colorful == true) {
		ct.ChangeColor(ct.Green, false, ct.None, false)
	}
	logger.logText(logger.Output, "INFO:", fmt, v...)
}

//warning log
func (logger *Logger) Warn(fmt string, v ...interface{}) {
	if(logger.Level > LevelWarning) {
		return
	}
	if(logger.Colorful == true) {
		ct.ChangeColor(ct.Yellow, false, ct.None, false)
	}
	logger.logText(logger.Output, "***WARN***:", fmt, v...)
}

//error log
func (logger *Logger) Error(fmt string, v ...interface{}) {
	if(logger.Level > LevelError) {
		return
	}
	if(logger.Colorful == true) {
		ct.ChangeColor(ct.Red, false, ct.None, false)
	}
	logger.logText(logger.Output, "***ERROR***:", fmt, v...)
}

//fatal error log
func (logger *Logger) Fatal(fmt string, v ...interface{}) {
	if(logger.Level > LevelFatal) {
		return
	}
	if(logger.Colorful == true) {
		ct.ChangeColor(ct.Red, false, ct.None, false)
	}
	logger.logText(logger.Output, "***FATAL***:", fmt, v...)
}

//******************************************************************************************//
//******************************************************************************************//
func (logger *Logger) logText(output io.Writer, levelFlagString, formatString string, v ...interface{}) {
	t := time.Now()
	fmt.Fprintf(output, "%d-%02d-%02d %02d:%02d:%02d.%03d - [%s]%s ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond() / 1000000, logger.Name, levelFlagString)
	fmt.Fprintf(output, formatString, v...)
	fmt.Fprintln(output)
}
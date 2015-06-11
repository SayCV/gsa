// Wrapper for log.Logger. The purpose behind this is to change where
// log messages go to by only changing this file.

package log

import (
  //"github.com/golang/glog"
	stdLog "log"
	"os"
)

var printLogger *stdLog.Logger
var errorLogger *stdLog.Logger

func getErrorLogger() *stdLog.Logger {
	if errorLogger == nil {
		// TODO allow configuring an error file here?
		errorLogger = stdLog.New(os.Stderr, "", stdLog.Ldate|stdLog.Ltime|stdLog.Lshortfile) 
	}

	return errorLogger
}

func getPrintLogger() *stdLog.Logger {
	if printLogger == nil {
		printLogger = stdLog.New(os.Stdout, "", stdLog.Ldate|stdLog.Ltime)
	}

	return printLogger
}

func Println(v ...interface{}) {
	getPrintLogger().Println(v...)
}

func Fatal(v ...interface{}) {
	getErrorLogger().Fatal(v...)
}

func Error(v ...interface{}) {
	getErrorLogger().Println(v...)
}

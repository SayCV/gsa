// A wrapper around glog and uses the syslog severity levels
// http://en.wikipedia.org/wiki/Syslog#Severity_levels

package log

import (
	"flag"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

const (
	logLevelEnvVar        = "LOG_LEVEL"
	logToStdErrEnvVar     = "LOG_TO_STDERR"
	logFileLocationEnvVar = "LOG_FILE_LOCATION"
)

var levelName = []string{"[emergency]", "[alert]", "[critical]", "[error]", "[warn]", "[notice]", "[info]", "[debug]"}

func Init() {
	var lL = os.Getenv(logLevelEnvVar)
	if lL == "" {
		lL = strconv.FormatInt(int64(len(levelName)-1), 10)
	}
	var lStd = os.Getenv(logToStdErrEnvVar)
	if lStd == "true" {
		flag.Lookup("alsologtostderr").Value.Set("true")
	}
	flag.Lookup("v").Value.Set(lL)

	var lfL = os.Getenv(logFileLocationEnvVar)
	if lfL != "" {
		flag.Lookup("log_dir").Value.Set(lfL)
	}
}

func callerAndLevelInfo(level glog.Level, args []interface{}) []interface{} {
	if int32(level) >= int32(len(levelName)) {
		level = glog.Level(len(levelName) - 1)
	}
	_, file, line, _ := runtime.Caller(3)
	slash := strings.LastIndex(file, "/")
	if slash >= 0 {
		file = file[slash+1:]
	}

	return append([]interface{}{levelName[level], file + ":" + strconv.Itoa(line) + "]", "\n"}, args...)
}

func logLevel(level glog.Level, args ...interface{}) {
	if glog.V(level) {
		args = callerAndLevelInfo(level, args)
		glog.Info(args...)
	}
}

func logLevelln(level glog.Level, args ...interface{}) {
	if glog.V(level) {
		args = callerAndLevelInfo(level, args)
		glog.Infoln(args...)
	}
}

func logLevelf(level glog.Level, format string, args ...interface{}) {
	if glog.V(level) {
		args = callerAndLevelInfo(level, args)
		format = "%s %s " + format
		glog.Infof(format, args...)
	}
}

func Emergency(args ...interface{}) {
	logLevel(0, args...)
}

func Emergencyln(args ...interface{}) {
	logLevelln(0, args...)
}

func Emergencyf(format string, args ...interface{}) {
	logLevelf(0, format, args...)
}

func Alert(args ...interface{}) {
	logLevel(1, args...)
}

func Alertln(args ...interface{}) {
	logLevelln(1, args...)
}

func Alertf(format string, args ...interface{}) {
	logLevelf(1, format, args...)
}

func Critical(args ...interface{}) {
	logLevel(2, args...)
}

func Criticalln(args ...interface{}) {
	logLevelln(2, args...)
}

func Criticalf(format string, args ...interface{}) {
	logLevelf(2, format, args...)
}

func Error(args ...interface{}) {
	logLevel(3, args...)
	glog.Error(args...)
}

func Errorln(args ...interface{}) {
	logLevelln(3, args...)
	glog.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	logLevelf(3, format, args...)
	glog.Errorf(format, args...)
}

func Warn(args ...interface{}) {
	logLevel(4, args...)
	glog.Warning(args...)
}

func Warnln(args ...interface{}) {
	logLevelln(4, args...)
	glog.Warningln(args...)
}

func Warnf(format string, args ...interface{}) {
	logLevelf(4, format, args...)
	glog.Warningf(format, args...)
}

func Notice(args ...interface{}) {
	logLevel(5, args...)
}

func Noticeln(args ...interface{}) {
	logLevelln(5, args...)
}

func Noticef(format string, args ...interface{}) {
	logLevelf(5, format, args...)
}

func Info(args ...interface{}) {
	logLevel(6, args...)
}

func Infoln(args ...interface{}) {
	logLevelln(6, args...)
}

func Infof(format string, args ...interface{}) {
	logLevelf(6, format, args...)
}

func Debug(args ...interface{}) {
	logLevel(7, args...)
}

func Debugln(args ...interface{}) {
	logLevelln(7, args...)
}

func Debugf(format string, args ...interface{}) {
	logLevelf(7, format, args...)
}
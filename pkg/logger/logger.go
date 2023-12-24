package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logLevelMap = map[string]logrus.Level{
	//"Trace":           logrus.TraceLevel,
	"Debug": logrus.DebugLevel,
	"Info":  logrus.InfoLevel,
	"Warn":  logrus.WarnLevel,
	"Error": logrus.ErrorLevel,
	"Fatal": logrus.FatalLevel,
	//"Panic":           logrus.PanicLevel,
}

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	level := "Debug"
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logLevelMap[level])
}

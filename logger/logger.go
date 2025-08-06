package logger

import "github.com/sirupsen/logrus"

var Log = logrus.New()

func Init() {
	Log.SetFormatter(&logrus.JSONFormatter{}) // or use TextFormatter
	Log.SetLevel(logrus.DebugLevel)
}

package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init() {
	// Open or create the log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// Fall back to stderr if file can't be opened
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	// Set output to the file
	Log.SetOutput(file)

	// Set formatter and level
	Log.SetFormatter(&logrus.JSONFormatter{}) // or use TextFormatter
	Log.SetLevel(logrus.DebugLevel)
}

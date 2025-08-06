package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init() {
	// Open or create the log file
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	// Set output to both file and console
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	Log.SetOutput(multiWriter)

	// Optional: JSON or Text formatter
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set the log level
	Log.SetLevel(logrus.DebugLevel)
}

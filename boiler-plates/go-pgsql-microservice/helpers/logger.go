package helpers

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Example usage:
// import "github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/helpers"
// helpers.log.WithFields(helpers.LogFields{ "errors": false  }).Info("Finished Task")

var Log = logrus.New()

type LogFields map[string]interface{}

func LogAsJSON() {
	Log.SetFormatter(&logrus.JSONFormatter{})
}

func SetupConsoleLogger() {
	// Show logs on the console
	Log.Out = os.Stdout
}

func SetupFileLogger(filePath string) {
	// Save logs to a file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("Failed to log to file, using default stderr")
	}
}

func SetLogLevel(level logrus.Level) {
	// Set Loglevel
	// Example: log.SetLevel(logrus.InfoLevel)
	Log.SetLevel(level)
}

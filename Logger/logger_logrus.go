package Logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Create a new instance of the logger. You can have any number of instances.
var logger = logrus.New()

func LoggerService() {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	logger.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }

	logger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}

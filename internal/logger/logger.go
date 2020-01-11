package logger

import (
	"github.com/sirupsen/logrus"
)

// Config represent the logger configuration
type Config struct {
	Environment string
	Name        string
}

// NewLogger create a new logger
func NewLogger(config *Config) *logrus.Entry {
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	log := logger.WithFields(logrus.Fields{"name": config.Name, "environment": config.Environment})

	return log
}

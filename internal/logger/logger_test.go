package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	tables := []struct {
		name        string
		environment string
		level       logrus.Level
	}{
		{"cisco-api", "develop", logrus.DebugLevel},
	}

	for _, table := range tables {
		logger := NewLogger(&Config{
			Name:        table.name,
			Environment: table.environment,
		})

		assert.NotNil(t, logger)
		// Check logging level
		assert.Equal(t, logger.Logger.Level, table.level, "they should be equal")

		// Check default variable
		assert.Equal(t, logger.Data["environment"], table.environment, "they should be equal")
		assert.Equal(t, logger.Data["name"], table.name, "they should be equal")
	}
}

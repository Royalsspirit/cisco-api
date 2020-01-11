package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test NewServer
func TestNewServer(t *testing.T) {
	port := "4242"
	env := "production"
	server := NewServer(&ServerConfig{Environment: env, Port: port})

	assert := assert.New(t)

	assert.Equal(server.Port, port, "they should be equal")
	assert.Equal(server.Environment, env, "they should be equal")
}

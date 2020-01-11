package api

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	server := Server{Port: "8080", Environment: "develop"}
	r := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	server.health(w, r)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, 200)
}

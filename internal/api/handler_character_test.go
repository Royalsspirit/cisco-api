package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type integrationHandlerSuite struct {
	APISuite
}

// TestAPISuiteIntegration represents the integration suite tests
func TestAPISuiteIntegration(t *testing.T) {
	x := &integrationHandlerSuite{APISuite{}}

	suite.Run(t, x)
}

type user struct {
	ID        uuid.UUID
	Pseudo    string
	Email     string
	Password  string
	Firstname string
	Lastname  string
}

// Testselect test get endpoints
func (p *integrationHandlerSuite) Testselect() {
	server := NewServer(&ServerConfig{
		DB:          p.DB,
		Port:        "5001",
		Environment: "develop",
		Validator:   p.Validator,
		Logger:      p.Logger,
	})

	_, err := http.NewRequest("GET", "/character", nil)

	if err != nil {
		p.T().Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.list())

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr)

	assert.Equal(p.T(), rr.Code, 200)
}

// TestUpdate test update endpoint
func (p *integrationHandlerSuite) TestUpdate() {
	server := NewServer(&ServerConfig{
		DB:          p.DB,
		Port:        "5001",
		Environment: "develop",
		Validator:   p.Validator,
		Logger:      p.Logger,
	})

	body := []byte(`{"pseudo": "pseudo", "email": "dreamseat@dreamseat.com", "password": "Test12345@"}`)
	req, err := http.NewRequest("PUT", "/character", bytes.NewBuffer(body))

	if err != nil {
		p.T().Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.update())

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr)

	assert.Equal(p.T(), rr.Code, 200)
}

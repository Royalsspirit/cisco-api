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

	req, err := http.NewRequest("GET", "/character", nil)

	if err != nil {
		p.T().Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.list)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

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

	body := []byte(`{"Name":"Luke Skywalker","Height":"172","Mass":"77","HairColor":"blond","SkinColor":"fair","EyeColor":"blue","BirthYear":"19BBY","Gender":"male","Homeworld":"1","Films":"","Vehicles":[{"ID":"30","Name":"Imperial Speeder Bike","Model":"74-Z speeder bike"},{"ID":"42","Name":"Imperial Speeder Bike","Model":"74-Z speeder bike"}],"Species":[{"ID":"1","Name":"Human","Classification":"mammal"},{"ID":"1","Name":"Human","Classification":"mammal"},{"ID":"4","Name":"Rodian","Classification":"sentient"},{"ID":"5","Name":"Hutt","Classification":"gastropod"},{"ID":"6","Name":"Yoda's ","Classification":"mammal"},{"ID":"7","Name":"Trandoshan","Classification":"reptile"},{"ID":"9","Name":"Ewok","Classification":"mammal"},{"ID":"10","Name":"Sullustan","Classification":"mammal"},{"ID":"11","Name":"Neimodian","Classification":"unknown"},{"ID":"12","Name":"Gungan","Classification":"amphibian"},{"ID":"14","Name":"Dug","Classification":"mammal"},{"ID":"18","Name":"Xexto","Classification":"unknown"},{"ID":"19","Name":"Toong","Classification":"unknown"},{"ID":"21","Name":"Nautolan","Classification":"amphibian"},{"ID":"22","Name":"Zabrak","Classification":"mammal"},{"ID":"25","Name":"Quermian","Classification":"mammal"},{"ID":"26","Name":"Kel Dor","Classification":"unknown"},{"ID":"28","Name":"Geonosian","Classification":"insectoid"},{"ID":"29","Name":"Mirialan","Classification":"mammal"},{"ID":"32","Name":"Kaminoan","Classification":"amphibian"},{"ID":"34","Name":"Muun","Classification":"mammal"},{"ID":"35","Name":"Togruta","Classification":"mammal"}],"Starships":"","Created":"2014-12-09T13:50:51.644000Z","URL":"1"}`)
	req, err := http.NewRequest("PUT", "/character/1", bytes.NewBuffer(body))

	if err != nil {
		p.T().Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.update)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	assert.Equal(p.T(), rr.Code, http.StatusNoContent)
}

//TestInsert TestInsert
func (p *integrationHandlerSuite) TestCreate() {
	server := NewServer(&ServerConfig{
		DB:          p.DB,
		Port:        "5001",
		Environment: "develop",
		Validator:   p.Validator,
		Logger:      p.Logger,
	})

	body := []byte(`{"Name":"Luke Skywalker","Height":"172","Mass":"77","HairColor":"blond","SkinColor":"fair","EyeColor":"blue","BirthYear":"19BBY","Gender":"male","Homeworld":"1","Films":"","Vehicles":[{"ID":"30","Name":"Imperial Speeder Bike","Model":"74-Z speeder bike"},{"ID":"42","Name":"Imperial Speeder Bike","Model":"74-Z speeder bike"}],"Species":[{"ID":"1","Name":"Human","Classification":"mammal"},{"ID":"1","Name":"Human","Classification":"mammal"},{"ID":"4","Name":"Rodian","Classification":"sentient"},{"ID":"5","Name":"Hutt","Classification":"gastropod"},{"ID":"6","Name":"Yoda's ","Classification":"mammal"},{"ID":"7","Name":"Trandoshan","Classification":"reptile"},{"ID":"9","Name":"Ewok","Classification":"mammal"},{"ID":"10","Name":"Sullustan","Classification":"mammal"},{"ID":"11","Name":"Neimodian","Classification":"unknown"},{"ID":"12","Name":"Gungan","Classification":"amphibian"},{"ID":"14","Name":"Dug","Classification":"mammal"},{"ID":"18","Name":"Xexto","Classification":"unknown"},{"ID":"19","Name":"Toong","Classification":"unknown"},{"ID":"21","Name":"Nautolan","Classification":"amphibian"},{"ID":"22","Name":"Zabrak","Classification":"mammal"},{"ID":"25","Name":"Quermian","Classification":"mammal"},{"ID":"26","Name":"Kel Dor","Classification":"unknown"},{"ID":"28","Name":"Geonosian","Classification":"insectoid"},{"ID":"29","Name":"Mirialan","Classification":"mammal"},{"ID":"32","Name":"Kaminoan","Classification":"amphibian"},{"ID":"34","Name":"Muun","Classification":"mammal"},{"ID":"35","Name":"Togruta","Classification":"mammal"}],"Starships":"","Created":"2014-12-09T13:50:51.644000Z","URL":"1"}`)
	req, err := http.NewRequest("POST", "/character", bytes.NewBuffer(body))

	if err != nil {
		p.T().Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.create)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	assert.Equal(p.T(), rr.Code, http.StatusCreated)
}

//TestDelete TestDelete
func (p *integrationHandlerSuite) TestDelete() {
	server := NewServer(&ServerConfig{
		DB:          p.DB,
		Port:        "5001",
		Environment: "develop",
		Validator:   p.Validator,
		Logger:      p.Logger,
	})

	req, err := http.NewRequest("DELETE", "/character/1", nil)

	if err != nil {
		p.T().Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.delete)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	assert.Equal(p.T(), rr.Code, http.StatusNoContent)
}

package util

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseError(t *testing.T) {
	var err Error
	errorExpected := Error{Status: 403, Message: "My Message"}

	w := httptest.NewRecorder()

	ResponseError(w, errorExpected.Message, errorExpected.Status)

	resp := w.Result()
	json.NewDecoder(resp.Body).Decode(&err)

	assert.Equal(t, err, errorExpected)
	assert.Equal(t, resp.StatusCode, errorExpected.Status)
}

func TestResponseErrors(t *testing.T) {
	var err Errors

	errors := []map[string]string{{"field": "pseudo"}}
	errorExpected := Errors{Status: 403, Message: "Validation error", Errors: errors}

	w := httptest.NewRecorder()

	ResponseErrors(w, errors, errorExpected.Status)

	resp := w.Result()
	json.NewDecoder(resp.Body).Decode(&err)

	assert.Equal(t, err, errorExpected)
	assert.Equal(t, resp.StatusCode, errorExpected.Status)
}

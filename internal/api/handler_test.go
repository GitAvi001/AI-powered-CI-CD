package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/GitAvi001/AI-powered-CI-CD/internal/model"
)

func TestPredictHandler(t *testing.T) {
	// Mock model
	m := &model.Model{} // Simplified, no actual classifier needed for this test
	handler := NewHandler(m)

	// Create a test request
	req, err := http.NewRequest("GET", "/predict", nil)
	assert.NoError(t, err)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler.PredictHandler(rr, req)

	// Check status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check response body
	expected := `{"prediction":"Iris-setosa"}`
	assert.JSONEq(t, expected, rr.Body.String())
}
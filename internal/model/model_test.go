package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//Unit testing for the model
func TestNewModel(t *testing.T) {
	m, err := NewModel()
	assert.NoError(t, err)
	assert.NotNil(t, m.Classifier)
}
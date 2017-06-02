package utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/swatlabs/GoDataberus/driver"
	"testing"
)

func TestGetDriver(t *testing.T) {
	input := "mongo"
	expected := driver.MongoDB{}
	result := GetDriver(input)
	assert.EqualValues(t, &expected, result)
}

func TestNewResourceID(t *testing.T) {
	result := NewResourceID()
	assert.NotEmpty(t, result)
}

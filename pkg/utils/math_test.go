package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeDiv(t *testing.T) {
	assert.Equal(t, 0.0, SafeDiv(5, 0))
	assert.Equal(t, 2.5, SafeDiv(5, 2))
}

func TestRound(t *testing.T) {
	assert.Equal(t, 42, Round(42.0))
	assert.Equal(t, 42, Round(42.2))
	assert.Equal(t, 43, Round(42.6))
}

func TestToFixed(t *testing.T) {
	assert.Equal(t, 42.42, ToFixed(42.42, 3))
	assert.Equal(t, 42.42, ToFixed(42.42, 2))
	assert.Equal(t, 42.4, ToFixed(42.42, 1))
	assert.Equal(t, 42.0, ToFixed(42.42, 0))
}

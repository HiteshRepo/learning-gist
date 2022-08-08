package Tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecipher(t *testing.T) {
	inputStr := "VTAOG"
	expected := "TRYME"
	actual := Decipher(inputStr)
	assert.Equal(t, expected, actual)
}

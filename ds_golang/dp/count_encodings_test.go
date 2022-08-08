package dp_test

import (
	"fmt"
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/dp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountEncodings(t *testing.T) {
	mappings := make(map[string]string)
	for i:=1; i<=26; i++ {
		ascii := 97 + (i-1)
		char := rune(ascii)
		mappings[fmt.Sprintf("%d", i)] = string(char)
	}

	code := "231011"
	expected := 4

	actual := dp.CountEncodings(code, mappings)
	assert.Equal(t, expected, actual)
}

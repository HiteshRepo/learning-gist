package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalancedParantheses(t *testing.T) {
	testcases := map[string]map[string]interface{} {
		"tc1": {
			"brackets": "{()}",
			"expected": true,
		},
		"tc2": {
			"brackets": "{)(}",
			"expected": false,
		},
	}

	for _, tc := range testcases {
		brackets := tc["brackets"].(string)
		expected := tc["expected"].(bool)
		actual := amazon.BalancedParantheses(brackets)
		assert.Equal(t, expected, actual)
	}
}

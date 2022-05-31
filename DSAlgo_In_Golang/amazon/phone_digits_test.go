package amazon_test

import (
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/amazon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPhoneNumberWords(t *testing.T) {
	num := "234"
	want := []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg",
		"bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}
	got := amazon.GetPhoneNumberWords(num)
	assert.Equal(t, want, got)
}

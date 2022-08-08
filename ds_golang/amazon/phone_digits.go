package amazon

import "fmt"

var phone = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

// Given phone digits, print all possible words that can be formed from them.
func GetPhoneNumberWords(num string) []string {
	ans := make([]string, 0)

	currStr := ""
	addWords(num, 0, &ans, &currStr, len(num))

	return ans
}

func addWords(phoneNumber string, currPos int, ans *[]string, currentString *string, n int) {
	if currPos == n {
		*ans = append(*ans, *currentString)
		return
	}

	currNumber := phoneNumber[currPos] - 48
	currWord := phone[currNumber]
	for _, c := range currWord {
		char := string(c)
		*currentString = fmt.Sprintf("%s%s", *currentString, char)
		addWords(phoneNumber, currPos+1, ans, currentString, n)
		*currentString = (*currentString)[0:len(*currentString)-1]
		if currNumber == 0 || currNumber == 1 {
			return
		}
	}
}

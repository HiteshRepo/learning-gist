package amazon

import (
	"fmt"
	"github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/stack"
)

// Given a string ‘str’ of digits and an integer ‘n’, build the lowest possible number by removing ‘n’ digits from the string and not changing the order of input digits.
func SmallestNumberByRemovingN(digits string, n int) string {
	//return recurseSolution(digits, n)
	return optimalSolution(digits, n)
}

func optimalSolution(digits string, n int) string {
	st := stack.NewStack()

	for _,d := range digits {
		for st.Len() > 0 && n > 0 && (st.Peek() > (int(d) - 48)) {
			st.Pop()
			n -= 1
		}

		if st.Len() > 0 || int(d) != 48 {
			st.Push(int(d) - 48)
		}
	}

	for st.Len() > 0 && n > 0 {
		st.Pop()
		n -= 1
	}

	if st.Len() == 0 {
		return ""
	}

	res := ""
	for st.Len() > 0 {
		d := st.Pop()
		res = fmt.Sprintf("%c%s", rune(d+48), res)
	}

	return res
}

func recurseSolution(digits string, n int) string {
	res := ""

	if n == 0 {
		return digits
	}

	if len(digits) < n {
		return res
	}

	removeNumbers(digits, n, &res)

	return res
}

func removeNumbers(currentDigits string, left int, ans *string) {
	if left <= 0 {
		if len(currentDigits) > 0 {
			*ans = fmt.Sprintf("%s%s", *ans, currentDigits)
		}
		return
	}

	pos := findLowestDigitIndex(currentDigits[0: left+1])

	*ans = fmt.Sprintf("%s%c", *ans, currentDigits[pos])
	currentDigits = fmt.Sprintf("%s", currentDigits[pos+1:])
	removeNumbers(currentDigits, left-pos, ans)
}

func findLowestDigitIndex(digits string) int {
	pos := 0
	lowest := int(digits[0]) - 48

	for i,d := range digits {
		num := int(d) - 48
		if num < lowest {
			lowest = num
			pos = i
		}
	}

	return pos
}

package Tests

import "fmt"

func DecodeStringFreq(input string) string {
	ans := ""

	prevAscii := -1
	currCount := 0
	for i:=0; i<len(input); i++ {
		ascii := int(input[i]) - 97
		if ascii == prevAscii || prevAscii == -1 {
			currCount += 1
			prevAscii = ascii
			continue
		}

		ans = fmt.Sprintf("%s%d", ans, prevAscii+1)
		if int(prevAscii + 97) > 105 {
			ans = fmt.Sprintf("%s#", ans)
		}

		if currCount > 1 {
			ans = fmt.Sprintf("%s(%d)", ans, currCount)
		}

		prevAscii = ascii
		currCount = 0
	}

	ans = fmt.Sprintf("%s%d", ans, prevAscii+1)
	if int(prevAscii + 97) > 105 {
		ans = fmt.Sprintf("%s#", ans)
	}

	if currCount+1 > 1 {
		ans = fmt.Sprintf("%s(%d)", ans, currCount+1)
	}

	return ans
}

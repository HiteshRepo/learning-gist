package Tests

import (
	"fmt"
	"strings"
)

func Decipher(inputStr string) string {
	copyStr := strings.ToLower(inputStr)

	alphabets := make([]rune, 0)

	for i:=97; i<=122; i++ {
		alphabets = append(alphabets, rune(i))
	}

	ans := ""
	for i:=0; i<len(copyStr); i++ {
		ch := copyStr[i]
		pos := int(ch) - 97
		var newCh rune
		if pos - 2 >= 0 {
			newCh = alphabets[pos-2]
		} else {
			pos = 26 + (pos - 2)
			newCh = alphabets[pos]
		}
		ans = fmt.Sprintf("%s%c", ans, newCh)
	}

	return strings.ToUpper(ans)
}

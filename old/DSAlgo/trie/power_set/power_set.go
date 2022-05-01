package power_set

import (
	"fmt"
	"math"
)

func GetAllPossibleStrings(str string) []string {
	n := len(str)

	possibleSubstrings := math.Pow(float64(2), float64(n))
	ans := make([]string, 0)

	for i:=0; i<int(possibleSubstrings); i++ {
		sub := ""
		for j:=0; j<n; j++ {
			bit := i & (1<<j)
			if bit != 0 {
				sub = fmt.Sprintf("%s%s", sub, string(str[j]))
			}
		}
		ans = append(ans, sub)
	}

	return ans
}

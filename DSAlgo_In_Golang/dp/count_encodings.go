package dp

import "fmt"

func CountEncodings(code string, mappings map[string]string) int {
	dpArr := make([]int, len(code))

	if _, ok := mappings[string(code[0])]; ok {
		dpArr[0] = 1
	}

	i := 1
	for i < len(code) {
		count := 0
		currCode := string(code[i])

		if currCode != "0" {
			count += dpArr[i-1]
		}

		currCode = fmt.Sprintf("%s%s", string(code[i-1]), currCode)
		_, ok := mappings[currCode]

		if ok && i==1 {
			count += 1
		}

		if ok && i>1 {
			count += dpArr[i-2]
		}

		dpArr[i] = count

		i += 1
	}

	return dpArr[len(dpArr)-1]
}

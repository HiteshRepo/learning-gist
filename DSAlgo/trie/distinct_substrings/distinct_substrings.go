package distinct_substrings

import "fmt"

func PrintDistinctSubstringsCount(word string) {
	allSubstrings := make([]string, 0)
	allSubstringsMap := make(map[string]struct{})
	allSubstrings = append(allSubstrings, "")

	for i:=0; i<len(word); i++ {
		str := ""
		currentSubstrings := make([]string, 0)
		for j:=i; j<len(word); j++ {
			str = fmt.Sprintf("%s%s", str, string(word[j]))
			currentSubstrings = append(currentSubstrings, str)
			if _, ok := allSubstringsMap[str]; ok {
				continue
			} else {
				allSubstringsMap[str] = struct{}{}
			}
			allSubstrings = append(allSubstrings, str)
		}
		fmt.Printf("substrings of %s are: %v\n", string(word[i]), currentSubstrings)
	}
	fmt.Printf("all substrings of are: %v\n", allSubstrings)
}

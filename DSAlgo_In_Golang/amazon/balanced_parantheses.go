package amazon

import "github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/stack"

func BalancedParantheses(brackets string) bool {
	closings := map[string]int{ ")": 0, "}": 1, "]": 2 }
	openings := map[string]int{ "(": 0, "{": 1, "[": 2 }

	st := stack.NewStack()

	first := string(brackets[0])
	last := string(brackets[len(brackets)-1])

	if _, ok := closings[first]; ok {
		return false
	}

	if _, ok := openings[last]; ok {
		return false
	}

	for _,b := range brackets {
		br := string(b)

		if n1, ok := closings[br]; ok && st.Len() > 0 {
			n2 := st.Pop()
			if n1 != n2 {
				return false
			}
			continue
		}

		n, _ := openings[br];
		st.Push(n)

	}

	return st.Len() == 0
}



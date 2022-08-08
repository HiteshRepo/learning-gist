package amazon

import "github.com/hiteshrepo/LearningGist/DSAlgo_In_Golang/stack"

func WellFormedParantheses(brackets string) int {
	st := stack.NewStack()
	st.Push(-1)

	res := 0

	for i,b := range brackets {
		br := string(b)

		if br == "(" {
			st.Push(i)
			continue
		}

		if st.Len() > 0 {
			st.Pop()
		}

		if st.Len() > 0 {
			res = max(res, i - st.Peek())
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

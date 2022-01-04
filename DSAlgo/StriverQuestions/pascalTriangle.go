func generate(numRows int) [][]int {
    count := 2
    ans := make([][]int, 0)
    
    ans = append(ans, []int{1})
    if numRows == 1 {
        return ans
    }
    
    ans = append(ans, []int{1,1})
    if numRows == 2 {
        return ans
    }
    
    for count < numRows {
        last := ans[len(ans)-1]
        prev := 0
        curr := make([]int, 0)
        for _,n := range last {
            curr = append(curr, n+prev)
            prev = n 
        }
        curr = append(curr, last[len(last)-1])
        ans = append(ans, curr)
        count = count+1
    }
    
    return ans
}

testCases := map[string]map[string]interface{}{
		"tc1": {
			"input":    5,
      "expected": [][]int{{1},{1,1},{1,2,1},{1,3,3,1},{1,4,6,4,1}},
		},
		"tc2": {
			"input":    1,
      "expected": [][]int{{1}},
		},
	}

	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
    input := tc["input"].(int)
    expected := tc["expected"].([][]int)
    actual := generate(input)
		fmt.Println("is solution correct: ", utils.IntMatricesEquals(actual, expected))
	}
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
Memory Usage: 2.1 MB, less than 98.95% of Go online submissions for Pascal's Triangle.
*/

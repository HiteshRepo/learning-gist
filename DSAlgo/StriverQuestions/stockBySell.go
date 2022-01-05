func maxProfit(prices []int) int {
    minPrice := int(^uint(0) >> 1)
    maxProfit := 0
    for _,p := range prices {
        if p < minPrice {
            minPrice = p
            continue
        }
        profit := p - minPrice
        if profit > maxProfit {
            maxProfit = profit
        }
    }
    return maxProfit
}

func runTestsMaxProfit() {
  testCases := map[string]map[string]interface{}{
		"tc1": {
			"input":    []int{7,1,5,3,6,4},
			"expected": 5,
		},
		"tc2": {
			"input":    {7,6,4,3,1},
			"expected": 0,
		},
	}

	for name, tc := range testCases {
		fmt.Printf("running test-case: %s\n", name)
    input := tc["input"].([]int)
    expected := tc["expected"].(int)
    actual := maxProfit(input)
		fmt.Println("is solution correct: ", expected == actual)
	}
}

/*
2nd:
Runtime: 120 ms, faster than 79.25% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 8.3 MB, less than 96.46% of Go online submissions for Best Time to Buy and Sell Stock.

1st:
Runtime: 280 ms, faster than 5.14% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 10.1 MB, less than 5.57% of Go online submissions for Best Time to Buy and Sell Stock.
*/

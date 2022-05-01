package array

import "fmt"

func maxProfitOptimized(prices []int) int {
	minPrice := int(^uint(0) >> 1)
	maxProfit := 0
	for _, p := range prices {
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

func maxProfitBruteForce(prices []int) int {
	maxProfit := 0
	for i:=0; i< len(prices); i++ {
		for j:=i+1; j<len(prices); j++ {
			if prices[j] > prices[i] {
				if maxProfit < (prices[j] - prices[i]) {
					maxProfit = prices[j] - prices[i]
				}
			}
		}
	}

	return maxProfit
}

func RunTestsMaxProfitBruteForce() {
	tcs := getTestCasesForMaxProfit()

	for name, tc := range tcs {
		fmt.Printf("running test-case: %s\n", name)
		input := tc["input"].([]int)
		expected := tc["expected"].(int)
		actual := maxProfitBruteForce(input)
		fmt.Println("is solution correct: ", expected == actual)
	}
}

func RunTestsMaxProfitOptimized() {
	tcs := getTestCasesForMaxProfit()

	for name, tc := range tcs {
		fmt.Printf("running test-case: %s\n", name)
		input := tc["input"].([]int)
		expected := tc["expected"].(int)
		actual := maxProfitOptimized(input)
		fmt.Println("is solution correct: ", expected == actual)
	}
}

func getTestCasesForMaxProfit() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"input":    []int{7, 1, 5, 3, 6, 4},
			"expected": 5,
		},
		"tc2": {
			"input":    []int{7, 6, 4, 3, 1},
			"expected": 0,
		},
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

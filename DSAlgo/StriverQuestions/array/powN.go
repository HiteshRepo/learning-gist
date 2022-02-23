package array

import (
	"fmt"
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/utils"
	"math"
)

func myPow_O_N(x float64, n int) float64 {
	count := n
	if n < 0 {
		count = -1 * n
	}
	ans := float64(1)
	pow := x

	if n == 0 {
		return 1
	}

	for count > 0 {
		if count % 2 == 0 {
			pow *= pow
			count = count/2
		} else {
			ans *= pow
			count -= 1
		}
	}

	if n < 0 {
		ans = 1/ans
	}

	return ans
}

func myPowBruteForce(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	ans := x
	count := 0
	for count < int(math.Abs(float64(n))) - 1 {
		ans *= x
		count += 1
	}

	if n < 0 {
		ans = 1/ans
	}

	return ans
}

func RunTestsForMyPowBruteForce()  {
	tcs := getTestCasesForMyPow()

	defer utils.Duration(utils.Track("MyPowBruteForce"))
	for name, tc := range tcs {
		x := tc["x"].(float64)
		n := tc["n"].(int)
		expected := tc["expected"].(float64)

		actual := myPowBruteForce(x, n)
		fmt.Printf("is solution correct for %s: %v\n", name, utils.CheckFloatEquals(expected, actual))
	}
}

func RunTestsForMyPow_O_N()  {
	tcs := getTestCasesForMyPow()

	defer utils.Duration(utils.Track("MyPow_O_N"))
	for name, tc := range tcs {
		x := tc["x"].(float64)
		n := tc["n"].(int)
		expected := tc["expected"].(float64)

		actual := myPow_O_N(x, n)
		fmt.Printf("is solution correct for %s: %v\n", name, utils.CheckFloatEquals(expected, actual))
	}
}

func getTestCasesForMyPow() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"tc1": {
			"x":        2.00000,
			"n":        10,
			"expected": 1024.0000,
		},
		"tc2": {
			"x":   2.10000,
			"n":   3,
			"expected": 9.26100,
		},
		"tc3": {
			"x":   2.00000,
			"n":   -2,
			"expected": 0.25000,
		},
		"tc4": {
			"x":   0.44528,
			"n":   0,
			"expected": 1.0000,
		},
	}
}

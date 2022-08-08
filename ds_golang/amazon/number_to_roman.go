package amazon

import "fmt"

// Given a number, find its corresponding Roman numeral.
func NumberToRoman(num int) string {
	ans := ""

	nums := []int{1,4,5,9,10,40,50,90,100,400,500,900,1000}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}


	for i:=len(nums)-1; i > 0; i-- {
		n := nums[i]
		for num / n > 0 {
			num = num - n
			ans = fmt.Sprintf("%s%s", ans, romans[i])
		}
	}

	return ans
}

func RomanToNumber(roman string) int {
	ans := 0

	romanMap := map[string]int {
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var prev int

	for i:=len(roman)-1; i>=0; i-- {
		c := roman[i]
		num := romanMap[string(c)]

		ans = ans + num

		if prev > num {
			ans = ans - (2 * num)
			prev = 0
			continue
		}

		prev = num
	}

	return ans
}

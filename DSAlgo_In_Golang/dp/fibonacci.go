package dp

func GetNthFibonacciNumber(n int) int {
	return fibonacci(n)
}


func fibonacci(n int) int {
	if n == 0 || n==1 {
		return n
	}

	fib1 := fibonacci(n-1)
	fib2 := fibonacci(n-2)
	fibn := fib1 + fib2

	return fibn
}

func GetNthFibonacciNumberOptimized(n int) int {
	fibMap := make(map[int]int)
	return fibonacciOptimized(n, fibMap)
}


func fibonacciOptimized(n int, fibMap map[int]int) int {
	if n == 0 || n==1 {
		return n
	}

	if fibn, ok := fibMap[n]; ok {
		return fibn
	}

	fib1 := fibonacciOptimized(n-1, fibMap)
	fib2 := fibonacciOptimized(n-2, fibMap)
	fibn := fib1 + fib2

	fibMap[n] = fibn

	return fibn
}
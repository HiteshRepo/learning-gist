package minimumCoins

func MinCoins(coins []int, value int) int {
	sortCoins(coins)
	copyVal := value
	coinMap := make(map[int]int)
	count := 0
	for i:=len(coins)-1; i>=0; i-- {
		if copyVal == 0 {
			break
		}

		if coins[i] > copyVal {
			continue
		}
		count += 1
		if _,ok := coinMap[coins[i]]; !ok {
			coinMap[coins[i]] = 1
		} else {
			coinMap[coins[i]] += 1
		}
		copyVal -= coins[i]
	}

	return count
}

func sortCoins(coins []int) {
	for i:=0; i<len(coins); i++ {
		minDenom := i
		for j:=i+1; j<len(coins); j++ {
			if coins[minDenom] > coins[j] {
				minDenom = j
			}
		}

		tnp := coins[minDenom]
		coins[minDenom] = coins[i]
		coins[i] = tnp
	}
}

package dp

func ArrangeBuildings(blocks int, roads int) int {
	noOfSides := roads + 1

	count := countCombinations(blocks)
	ans := 1

	i := 0
	for i < noOfSides {
		ans = ans*count
		i+=1
	}

	return ans
}

func countCombinations(length int) int {

	previousSpaces := 1
	previousBuildings := 1

	currentSpaces := 0
	currentBuildings := 0

	i := 0
	for i < length {
		currentSpaces = previousSpaces + previousBuildings
		currentBuildings = previousSpaces

		previousSpaces = currentSpaces
		previousBuildings = currentBuildings

		i += 1
	}

	return currentBuildings + currentSpaces
}
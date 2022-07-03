package dp

func CountBinaryStrings(length int) int {

	previousOnes := 1
	previousZeroes := 1

	currentOnes := 0
	currentZeroes := 0

	i := 0
	for i < length {
		currentOnes = previousOnes + previousZeroes
		currentZeroes = previousOnes

		previousOnes = currentOnes
		previousZeroes = currentZeroes

		i += 1
	}

	return currentZeroes + currentOnes
}

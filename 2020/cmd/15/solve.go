package main

func solve(input []int, resultIndex int) int {
	lastFoundIndex := make(map[int]int)

	for i, x := range input[:len(input)-1] {
		lastFoundIndex[x] = i
	}

	previousNum := input[len(input)-1]
	for i := len(input); i < resultIndex; i++ {
		previousNumIndex, found := lastFoundIndex[previousNum]
		currentNum := 0
		if found {
			currentNum = i - 1 - previousNumIndex
		}

		lastFoundIndex[previousNum] = i - 1
		previousNum = currentNum
	}

	return previousNum
}

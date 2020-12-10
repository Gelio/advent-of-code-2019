package main

import (
	"errors"
)

func solveA(ns []int, preambleLen int) (int, error) {
	sums := make(map[int]int)

	for i := 0; i < preambleLen; i++ {
		for j := 0; j < i; j++ {
			sums[ns[i]+ns[j]]++
		}
	}

	for i := preambleLen; i < len(ns); i++ {
		if sums[ns[i]] == 0 {
			return ns[i], nil
		}

		oldestNum := ns[i-preambleLen]
		for j := i - preambleLen + 1; j < i; j++ {
			sums[oldestNum+ns[j]]--
			sums[ns[i]+ns[j]]++
		}
	}

	return 0, errors.New("solution not found")
}

func solveB(ns []int, target int) (min, max int, err error) {
	startIdx, endIdx, sum := 0, 0, ns[0]

	for startIdx < len(ns) && sum != target {
		if sum < target {
			// Add next number
			endIdx++
			if endIdx == len(ns) {
				break
			}

			sum += ns[endIdx]
		} else {
			// Remove oldest number
			sum -= ns[startIdx]
			startIdx++

			if startIdx == len(ns) {
				break
			}

			if startIdx > endIdx {
				endIdx = startIdx
				sum += ns[startIdx]
			}
		}
	}

	if sum == target {
		n := ns[startIdx:(endIdx + 1)]
		return getMin(n), getMax(n), nil
	}

	return 0, 0, errors.New("cannot find solution")
}

func getMin(ns []int) int {
	m := ns[0]
	for _, x := range ns[1:] {
		if x < m {
			m = x
		}
	}

	return m
}

func getMax(ns []int) int {
	m := ns[0]
	for _, x := range ns[1:] {
		if x > m {
			m = x
		}
	}

	return m
}

package main

import (
	"errors"
	"strconv"
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

	return 0, errors.New("Solution not found")
}

func parseNums(lines []string) ([]int, error) {
	nums := make([]int, 0, len(lines))

	for _, l := range lines {
		num, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		nums = append(nums, num)
	}

	return nums, nil
}

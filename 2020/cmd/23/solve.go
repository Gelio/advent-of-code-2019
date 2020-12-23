package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveA(input string, moves int) (string, error) {
	nums, err := parseNums(input)
	if err != nil {
		return "", fmt.Errorf("cannot parse input: %w", err)
	}

	cups := getCupsFromNums(nums)

	simulate(cups, nums[0], moves)

	resultCupNums, err := getCupNumbers(cups)
	if err != nil {
		return "", fmt.Errorf("cannot get final cup numbers: %w", err)
	}

	var resultCupRawNums []string
	for _, num := range resultCupNums {
		resultCupRawNums = append(resultCupRawNums, fmt.Sprintf("%d", num))
	}

	// Remember to omit the initial 1
	return strings.Join(resultCupRawNums[1:], ""), nil
}

func parseNums(input string) ([]int, error) {
	rawNums := strings.Split(input, "")
	var nums []int
	for _, rawNum := range rawNums {
		num, err := strconv.Atoi(rawNum)
		if err != nil {
			return nil, fmt.Errorf("cannot parse number %q: %w", rawNum, err)
		}

		nums = append(nums, num)
	}

	return nums, nil
}

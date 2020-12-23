package main

import (
	"errors"
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

func solveB(input string) (int, error) {
	nums, err := parseNums(input)
	if err != nil {
		return 0, fmt.Errorf("cannot parse input: %w", err)
	}

	nums = appendNumsUntil(nums, 1000000)

	cups := getCupsFromNums(nums)

	const moves = 10000000
	simulate(cups, nums[0], moves)

	// Find 2 values immediately after 1
	cupWith1 := cups[1]
	if cupWith1 == nil {
		return 0, errors.New("cannot find cup with number 1")
	}

	nextCup := cupWith1.Next
	result := nextCup.Value * nextCup.Next.Value

	return result, nil
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

func appendNumsUntil(nums []int, target int) []int {
	res := make([]int, 0, target)
	res = append(res, nums...)

	i := nums[0]
	for _, num := range nums {
		if num > i {
			i = num
		}
	}

	for i = i + 1; i <= target; i++ {
		res = append(res, i)
	}

	return res
}

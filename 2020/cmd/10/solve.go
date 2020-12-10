package main

import "sort"

func solveA(nums []int) int {
	sort.Ints(nums)

	oneJumps, threeJumps := 0, 1

	v := 0

	for i := 0; i < len(nums); i++ {
		diff := nums[i] - v
		if diff == 1 {
			oneJumps++
		} else if diff == 3 {
			threeJumps++
		}

		if diff > 0 && diff <= 3 {
			v = nums[i]
		}
	}

	return oneJumps * threeJumps
}

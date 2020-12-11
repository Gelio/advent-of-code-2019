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

func solveB(originalNums []int) int {
	nums := originalNums
	sort.Ints(nums)

	nums = append([]int{0}, nums...)
	sink := nums[len(nums)-1] + 3
	nums = append(nums, sink)

	g := newAdapterGraph(nums)

	return g.CountPathsToSink(0, sink)
}

type adapterGraph map[int][]int

func newAdapterGraph(nums []int) adapterGraph {
	g := make(map[int][]int)

	for i, v := range nums {
		for _, w := range nums[i+1:] {
			if w == v {
				continue
			}

			if w-v > 3 {
				break
			}

			g[v] = append(g[v], w)
		}
	}

	return g
}

func (g adapterGraph) CountPathsToSink(start, sink int) int {
	results := make(map[int]int)
	results[sink] = 1

	g.countPathsToSinkRecursive(start, results)

	return results[start]
}

func (g adapterGraph) countPathsToSinkRecursive(v int, results map[int]int) {
	pathsForV := 0

	for _, n := range g[v] {
		if results[n] == 0 {
			g.countPathsToSinkRecursive(n, results)
		}

		pathsForV += results[n]
	}

	results[v] = pathsForV
}

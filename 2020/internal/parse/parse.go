package parse

import "strconv"

// Ints parses an array of strings into integers
func Ints(lines []string) ([]int, error) {
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

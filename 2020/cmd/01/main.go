package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var nums []int

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("Error while reading input", err)
			return
		}

		if len(line) == 0 {
			break
		}

		num, err := strconv.Atoi(strings.TrimSpace(line))

		if err != nil {
			fmt.Println("Cannot parse line as number", line)
			return
		}

		nums = append(nums, num)
	}

	const targetSum = 2020
	for i, x := range nums {
		for j, y := range nums[i:] {
			for _, z := range nums[j:] {
				if x+y+z == targetSum {
					fmt.Println("Result: ", x*y*z)
				}
			}
		}
	}
}

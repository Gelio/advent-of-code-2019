package main

import "fmt"

func main() {
	input := []int{6, 4, 12, 1, 20, 0, 16}

	fmt.Println("Result A:", solve(input, 2020))
	fmt.Println("Result B:", solve(input, 30000000))
}

package main

import "fmt"

func main() {
	input := "583976241"

	res, err := solveA(input, 100)
	if err != nil {
		fmt.Println("Error while solving A:", err)
		return
	}

	fmt.Println("Result A:", res)
}

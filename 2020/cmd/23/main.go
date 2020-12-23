package main

import "fmt"

func main() {
	input := "583976241"

	resA, err := solveA(input, 100)
	if err != nil {
		fmt.Println("Error while solving A:", err)
		return
	}

	fmt.Println("Result A:", resA)

	resB, err := solveB(input)
	if err != nil {
		fmt.Println("Error while solving B:", err)
		return
	}

	fmt.Println("Result B:", resB)
}

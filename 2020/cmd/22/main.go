package main

import (
	"fmt"
)

func main() {
	d1, d2, err := getPlayerDecksFromInput()
	if err != nil {
		fmt.Println("Error when getting input:", err)
	}

	fmt.Println("Result A:", solveA(d1, d2))
	fmt.Println("Result B:", solveB(d1, d2))
}

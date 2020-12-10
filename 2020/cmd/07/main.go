package main

import (
	"aoc-2020/internal/stdin"
	"fmt"
)

type bagWithQuantity struct {
	color    string
	quantity int
}

type rule struct {
	bagColor string
	contents []*bagWithQuantity
}

func main() {
	lines, err := stdin.ReadAllLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	// QUESTION: is this a good way to synchronize wait for subroutines to finish? Should I use `sync.WaitGroup`?
	c := make(chan bool)

	// QUESTION: does it make sense to run those 2 parts concurrently?
	go func() {
		resultA, err := SolveA(lines)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Result A:", resultA)
		c <- true
	}()

	go func() {
		resultB, err := SolveB(lines)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Result B:", resultB)
		c <- true
	}()

	<-c
	<-c
}

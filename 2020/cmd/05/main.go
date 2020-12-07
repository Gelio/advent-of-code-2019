package main

import (
	"aoc-2020/internal/stdin"
	"fmt"
	"sort"
)

func main() {
	lines, err := stdin.ReadAllLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	seats := make([]int, 0, len(lines))

	for _, line := range lines {
		seat, err := getPassengerSeat(line)
		if err != nil {
			fmt.Println("Cannot get passenger seat for", line, ". Error: ", err)
			continue
		}
		seatID := getSeatID(&seat)
		seats = append(seats, seatID)
	}

	sort.Ints(seats)

	for i := 1; i < len(seats)-2; i++ {
		if seats[i]-seats[i-1] == 2 {
			fmt.Println("Result:", seats[i]-1)
			return
		}
	}
}

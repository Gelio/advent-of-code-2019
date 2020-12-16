package main

import (
	"aoc-2020/internal/stdin"
	"fmt"
	"strings"
)

func main() {
	specs, err := parseSpecs(strings.Split(`departure location: 40-152 or 161-969
departure station: 39-838 or 845-971
departure platform: 39-209 or 217-970
departure track: 47-76 or 82-955
departure date: 41-167 or 178-949
departure time: 25-652 or 660-953
arrival location: 36-798 or 810-964
arrival station: 30-688 or 702-973
arrival platform: 44-248 or 268-969
arrival track: 45-536 or 552-956
class: 29-751 or 760-951
duration: 40-912 or 934-971
price: 44-896 or 911-965
route: 32-582 or 590-953
row: 46-269 or 282-971
seat: 49-114 or 134-971
train: 37-395 or 401-969
type: 43-180 or 206-960
wagon: 41-462 or 480-953
zone: 35-411 or 427-960`, "\n"))

	if err != nil {
		fmt.Println("Error when parsing specs", err)
		return
	}

	rawNearbyTickets, err := stdin.ReadLinesFromFile("nearby-tickets.txt")
	if err != nil {
		fmt.Println("Error when reading input", err)
		return
	}
	nearbyTickets, err := parseTickets(rawNearbyTickets)
	if err != nil {
		fmt.Println("Error when parsing nearby tickets", err)
		return
	}

	res := solveA(specs, nearbyTickets)
	fmt.Println("Result A:", res)

	myTicket := []int{139, 109, 61, 149, 101, 89, 103, 53, 107, 59, 73, 151, 71, 67, 97, 113, 83, 163, 137, 167}
	res = solveB(specs, myTicket, nearbyTickets)
	fmt.Println("Result B:", res)
}

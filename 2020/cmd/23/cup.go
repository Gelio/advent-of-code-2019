package main

import "fmt"

type cup struct {
	Value int
	// Next cup clockwise
	Next *cup
}

func getCupsFromNums(nums []int) map[int]*cup {
	m := make(map[int]*cup)

	var prevCup *cup
	for _, v := range nums {
		c := &cup{Value: v}
		if prevCup != nil {
			prevCup.Next = c
		}
		prevCup = c
		m[c.Value] = c
	}

	prevCup.Next = m[nums[0]]

	return m
}

func simulate(cups map[int]*cup, startingNum, moves int) error {
	c := cups[startingNum]

	if c == nil {
		return fmt.Errorf("cannot found cup for starting number %d", startingNum)
	}

	highestNum := 1
	for num := range cups {
		if num > highestNum {
			highestNum = num
		}
	}

	for i := 0; i < moves; i++ {
		// Pick up 3 next cups
		nextCup := c.Next
		middlePickedUpCup := nextCup.Next
		lastPickedUpCup := middlePickedUpCup.Next

		c.Next = lastPickedUpCup.Next
		lastPickedUpCup.Next = nil

		// Find destination cup
		destinationNum := c.Value - 1
		for {
			if destinationNum < 1 {
				destinationNum = highestNum
			}

			destinationCup := cups[destinationNum]
			if destinationCup != nextCup && destinationCup != middlePickedUpCup && destinationCup != lastPickedUpCup {
				break
			}

			destinationNum--
		}

		// Place 3 picked up cups after the destination cup
		destinationCup := cups[destinationNum]
		lastPickedUpCup.Next = destinationCup.Next
		destinationCup.Next = nextCup

		// Move the current cup
		c = c.Next
	}

	return nil
}

func getCupNumbers(cups map[int]*cup) ([]int, error) {
	firstCup := cups[1]

	if firstCup == nil {
		return nil, fmt.Errorf("cup with number 1 was not found")
	}

	nums := []int{firstCup.Value}
	for c := firstCup.Next; c != firstCup; c = c.Next {
		nums = append(nums, c.Value)
	}

	return nums, nil
}

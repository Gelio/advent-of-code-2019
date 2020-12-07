package main

import "testing"

func TestGetPassengerSeat(t *testing.T) {
	cases := []struct {
		input       string
		row, column int
	}{
		{input: "FBFBBFFRLR", row: 44, column: 5},
		{input: "BFFFBBFRRR", row: 70, column: 7},
		{input: "FFFBBBFRRR", row: 14, column: 7},
		{input: "BBFFBBFRLL", row: 102, column: 4},
	}

	for i, c := range cases {
		res, err := getPassengerSeat(c.input)
		if err != nil {
			t.Errorf("Case %d failed: %v", i+1, err)
			continue
		}

		if res.row != c.row {
			t.Errorf("Case %d failed: Invalid row. Got %d, expected %d", i+1, res.row, c.row)
		}

		if res.column != c.column {
			t.Errorf("Case %d failed: Invalid column. Got %d, expected %d", i+1, res.column, c.column)
		}
	}
}

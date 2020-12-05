package main

import "testing"

func TestGetSeatIDFromPass(t *testing.T) {
	example := "FBFBBFFRLR"
	result := getSeatIDFromPass(example)
	if result != 357 {
		t.Errorf("Got wrong seat ID. Got %d; want %d", result, 357)
	}
	example = "BFFFBBFRRR"
	result = getSeatIDFromPass(example)
	if result != 567 {
		t.Errorf("Got wrong seat ID. Got %d; want %d", result, 567)
	}

	example = "FFFBBBFRRR"
	result = getSeatIDFromPass(example)
	if result != 119 {
		t.Errorf("Got wrong seat ID. Got %d; want %d", result, 119)
	}

	example = "BBFFBBFRLL"
	result = getSeatIDFromPass(example)
	if result != 820 {
		t.Errorf("Got wrong seat ID. Got %d; want %d", result, 820)
	}
}

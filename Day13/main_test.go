package main

import (
	"testing"
)

func TestMinutesUntilArrival(t *testing.T) {
	res := minutesUntilArrival(939, 59)
	if res != 5 {
		t.Errorf("Wrong minutes until arrival. Want %d; got %d", 5, res)
	}
}

func TestTakeEarliestBus(t *testing.T) {
	res := takeEarliestBus("./example.txt")
	if res != 295 {
		t.Errorf("Wrong bus taken. Want %d; got %d", 295, res)
	}
}

func TestFindConvergence(t *testing.T) {
	res := findConvergence("./example.txt")
	if res != 1068781 {
		t.Errorf("Wrong convergence. Want %d; got %d", 1068781, res)
	}
}

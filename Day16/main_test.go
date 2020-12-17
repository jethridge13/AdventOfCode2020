package main

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	ranges, ticket, others := parseInput("./example.txt")
	if len(ranges) != 6 ||
		len(ticket) != 3 ||
		len(others) != 12 {
		t.Errorf("Error in parsing input. Wrong slice size")
	}
}

func TestFindInvalidTickets(t *testing.T) {
	res := findInvalidTickets("./example.txt")
	if res != 71 {
		t.Errorf("Error in invalid ticket number. Want %d; got %d", 71, res)
	}
}

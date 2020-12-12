package main

import "testing"

func TestBuildGrid(t *testing.T) {
	grid := buildGrid("./example.txt")
	if len(grid) != 10 || len(grid[0]) != 10 {
		t.Errorf("Grid built wrong")
	}
	if grid[0][0] != 'L' {
		t.Errorf("Grid built wrong")
	}
}

func TestApplyRules(t *testing.T) {
	grid := buildGrid("./example.txt")
	count := applyRules(grid, 1)
	if count != 71 {
		t.Errorf("Grid rules applied wrong. %d changed; want %d", count, 71)
	}
}

func TestCountOccupiedSeats(t *testing.T) {
	grid := buildGrid("./example.txt")
	applyUntilStable(grid)
	count := countOccupiedSeats(grid)
	if count != 37 {
		t.Errorf("Final seats wrong. Got %d; want %d", count, 37)
	}
}

func TestCountOccupiedSeatsPart2(t *testing.T) {
	grid := buildGrid("./example2.txt")
	applyUntilStablePart2(grid)
	count := countOccupiedSeats(grid)
	if count != 26 {
		t.Errorf("Final seals wrong. Got %d; want %d", count, 26)
	}
}

package main

import "testing"

func TestCountTiles(t *testing.T) {
	inp := map[[2]int]bool{
		[2]int{0, 0}: true,
		[2]int{0, 1}: false,
		[2]int{1, 0}: true,
	}
	res := countTiles(inp)
	exp := 2
	if res != exp {
		t.Errorf("Wrong tiles counted. Want %d; got %d", exp, res)
	}
}

func TestPlaceTiles(t *testing.T) {
	res := placeTiles("./example.txt")
	count := countTiles(res)
	expCount := 10
	if count != expCount {
		t.Errorf("Wrong tile placement. %d tiles placed, should be %d", count, expCount)
	}
}

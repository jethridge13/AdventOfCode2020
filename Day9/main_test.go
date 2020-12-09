package main

import (
	"testing"
)

func TestAddToXmas(t *testing.T) {
	src := xmas{sumCount: 5}
	newX := addToXmas(src, 1)
	if len(newX.recentN) != 1 {
		t.Errorf("Did not add to xmas")
	}
	if len(src.recentN) != 0 {
		t.Errorf("Original xmas changed")
	}

	src = xmas{sumCount: 5, recentN: []int{1, 2, 3, 4, 5}}
	newX = addToXmas(src, 6)
	if len(newX.recentN) != 5 {
		t.Errorf("Added more than max recent size: %+v", newX)
	}
	if newX.recentN[0] != 2 || newX.recentN[4] != 6 {
		t.Errorf("N added incorrectly: %+v", newX)
	}
}

func TestVerifyN(t *testing.T) {
	recentN := make([]int, 25)
	for i := range recentN {
		recentN[i] = i + 1
	}
	src := xmas{sumCount: 25, recentN: recentN}
	res := verifyN(src, 26)
	if !res {
		t.Errorf("Could not verify %d for %+v", 26, src)
	}
	res = verifyN(src, 49)
	if !res {
		t.Errorf("Could not verify %d for %+v", 49, src)
	}
	res = verifyN(src, 50)
	if res {
		t.Errorf("Could not verify %d for %+v", 50, src)
	}
	res = verifyN(src, 100)
	if res {
		t.Errorf("Could not verify %d for %+v", 100, src)
	}
}

func TestFindFirstInvalidN(t *testing.T) {
	res, err := findFirstInvalidN("./example.txt", 5)
	if err != nil {
		t.Errorf("Got error when checking first invalid: %s", err)
	}
	if res != 127 {
		t.Errorf("First invalid check failed. Want %d; got %d", 127, res)
	}
}

func TestFindContigiousSet(t *testing.T) {
	res, err := findContiguousSet("./example.txt", 5)
	if err != nil {
		t.Errorf("Got error when finding set: %s", err)
	}
	if res != 62 {
		t.Errorf("Wrong set. Want %d; got %d", 62, res)
	}
}

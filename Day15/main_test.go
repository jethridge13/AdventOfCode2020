package main

import (
	"testing"
)

func TestFillStartMap(t *testing.T) {
	res, nextTurn, count := fillStartMap(EXAMPLE1)
	if count != 4 {
		t.Errorf("Wrong start count. Want %d; got %d", 4, count)
	}
	if nextTurn != 0 {
		t.Errorf("Wrong next turn. Want %d; got %d", 0, nextTurn)
	}
	if len(res) != 3 {
		t.Errorf("Wrong map returned")
	}
}

func TestGetNthNumber(t *testing.T) {
	res := getNthNumber(EXAMPLE1, 4)
	if res != 0 {
		t.Errorf("%s: Turn %d should be %d but got %d", EXAMPLE1, 4, 0, res)
	}

	res = getNthNumber(EXAMPLE1, 5)
	if res != 3 {
		t.Errorf("%s: Turn %d should be %d but got %d", EXAMPLE1, 5, 3, res)
	}

	res = getNthNumber(EXAMPLE1, 6)
	if res != 3 {
		t.Errorf("%s: Turn %d should be %d but got %d", EXAMPLE1, 6, 3, res)
	}

	res = getNthNumber(EXAMPLE1, 7)
	if res != 1 {
		t.Errorf("%s: Turn %d should be %d but got %d", EXAMPLE1, 7, 1, res)
	}

	res = getNthNumber(EXAMPLE1, 8)
	if res != 0 {
		t.Errorf("%s: Turn %d should be %d but got %d", EXAMPLE1, 8, 0, res)
	}

	res = getNthNumber(EXAMPLE1, 9)
	if res != 4 {
		t.Errorf("%s: Turn %d should be %d but got %d", EXAMPLE1, 9, 4, res)
	}

	res = getNthNumber(EXAMPLE1, 10)
	if res != 0 {
		t.Errorf("%s: Turn %d should be %d but got %d", EXAMPLE1, 10, 0, res)
	}

	res = getNthNumber(EXAMPLE1, 2020)
	if res != 436 {
		t.Errorf("%s: Turn %d should be %d but got %d", EXAMPLE1, 2020, 436, res)
	}

	res = getNthNumber(EXAMPLE2, 2020)
	if res != 1 {
		t.Errorf("%s: Turn %d should be %d but got %d", EXAMPLE2, 2020, 1, res)
	}

	res = getNthNumber(EXAMPLE3, 2020)
	if res != 10 {
		t.Errorf("%s: Turn %d should be %d but got %d", EXAMPLE3, 2020, 10, res)
	}

	res = getNthNumber(EXAMPLE4, 2020)
	if res != 27 {
		t.Errorf("%s: Turn %d should be %d but got %d", EXAMPLE4, 2020, 27, res)
	}

}

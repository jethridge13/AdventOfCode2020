package main

import (
	"testing"
)

func TestGetDiffs(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	res := getDiffs(list)
	if len(res) != 2 {
		t.Errorf("Wrong number of diffs returned")
	}
	if res[1] != 5 {
		t.Errorf("Diff check wrong: %+v", res)
	}

	list, err := buildList("./example.txt")
	if err != nil {
		t.Errorf("buildList failed")
	}
	res = getDiffs(list)
	if res[1] != 7 || res[3] != 5 {
		t.Errorf("Diff check wrong: %+v", res)
	}

	list, err = buildList("./example2.txt")
	if err != nil {
		t.Errorf("buildList failed")
	}
	res = getDiffs(list)
	if res[1] != 22 || res[3] != 10 {
		t.Errorf("Diff check wrong: %+v", res)
	}
}

func TestGetListOfOneDiffs(t *testing.T) {
	list, err := buildList("./example.txt")
	if err != nil {
		t.Errorf("buildList failed")
	}
	res := getListOfOneDiffs(list)
	if len(res) != 5 {
		t.Errorf("Wrong list of ones")
	}
	prod := getProduct(res)
	if prod != 8 {
		t.Errorf("Wrong get product. Want %d; got %d", 8, prod)
	}

	list, err = buildList("./example2.txt")
	if err != nil {
		t.Errorf("buildList failed")
	}
	res = getListOfOneDiffs(list)
	prod = getProduct(res)
	if prod != 19208 {
		t.Errorf("Wrong product. Want %d; got %d", 19208, prod)
	}
}

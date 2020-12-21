package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	ingredients, allergens := parseLine("mxmxvkd kfcds sqjhc nhms (contains dairy, fish)")
	if len(ingredients) != 4 || len(allergens) != 2 {
		t.Errorf("Error in line parsing: %v, %v", ingredients, allergens)
	}

	ingredients, allergens = parseLine("trh fvjkl sbzzf mxmxvkd (contains dairy)")
	if len(ingredients) != 4 || len(allergens) != 1 {
		t.Errorf("Error in line parsing: %v, %v", ingredients, allergens)
	}
}

func TestSetUnion(t *testing.T) {
	u1 := map[string]bool{"a": true, "c": true}
	u2 := map[string]bool{"b": true, "c": true}
	union := setUnion(u1, u2)
	if !union["c"] || union["a"] || union["b"] {
		t.Errorf("Wrong union: %v & %v -> %v", u1, u2, union)
	}
}

func TestSetDifference(t *testing.T) {
	u1 := map[string]bool{"a": true, "c": true}
	u2 := map[string]bool{"b": true, "c": true}
	diff := setDifference(u1, u2)
	if !diff["a"] || diff["b"] || diff["c"] {
		t.Errorf("Wrong diff: %v & %v -> %v", u1, u2, diff)
	}
}

func TestFindNonAllergens(t *testing.T) {
	res, _ := findNonAllergens("./example.txt")
	if res != 5 {
		t.Errorf("Wrong list of nonallergens. Want %d; got %d", 4, res)
	}
}

func TestFindDangerousList(t *testing.T) {
	_, set := findNonAllergens("./example.txt")
	list := findDangerousList(set)
	exp := "mxmxvkd,sqjhc,fvjkl"
	if list != exp {
		t.Errorf("Wrong list. Want %s; got %s", exp, list)
	}

	_, set = findNonAllergens("./input.txt")
	list = findDangerousList(set)
	exp = "pgnpx,srmsh,ksdgk,dskjpq,nvbrx,khqsk,zbkbgp,xzb"
	if list != exp {
		t.Errorf("Wrong list. Want %s; got %s", exp, list)
	}
}

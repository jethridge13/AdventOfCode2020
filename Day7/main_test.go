package main

import (
	"log"
	"os"
	"testing"
)

func TestGetBagWithCount(t *testing.T) {
	input := "1 bright white bag"
	bag, err := getBagWithCount(input)
	if err != nil {
		t.Errorf("Incorrectly reported no bag")
	}
	if bag.bag != "bright white" || bag.count != 1 {
		t.Errorf("Got wrong bag: %+v", bag)
	}
	input = "2 muted yellow bags."
	bag, err = getBagWithCount(input)
	if err != nil {
		t.Errorf("Incorrectly reported no bag")
	}
	if bag.bag != "muted yellow" || bag.count != 2 {
		t.Errorf("Got wrong bag: %v", bag)
	}
	input = "no other bags."
	bag, err = getBagWithCount(input)
	if err == nil {
		t.Errorf("Did not detect no other bags")
	}
}

func TestParseLine(t *testing.T) {
	input := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	bag, list := parseLine(input)
	if bag != "light red" {
		t.Errorf("Got wrong bag name. Want light red; got %s", bag)
	}
	if len(list) != 2 {
		t.Errorf("Got wrong bag list")
	}
}

func TestBuildBagMap(t *testing.T) {
	input, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	bagMap := buildBagMap(input)
	input.Close()
	if len(bagMap) != 9 {
		t.Errorf("Wrong number of bags in map")
	}
	if len(bagMap["light red"]) != 2 {
		t.Errorf("Bag map built wrong")
	}
	if len(bagMap["dotted black"]) != 0 {
		t.Errorf("Bag map built wrong")
	}
}

func TestBuildReverseBagMap(t *testing.T) {
	input, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	bagMap := buildBagMap(input)
	revMap := buildReverseBagMap(bagMap)
	if len(revMap["bright white"]) != 2 {
		t.Errorf("Reverse map built wrong")
	}
	if len(revMap["vibrant plum"]) != 1 {
		t.Errorf("Reverse map built wrong")
	}
}

func TestGetNestedBagCount(t *testing.T) {
	input, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	bagMap := buildBagMap(input)
	revMap := buildReverseBagMap(bagMap)
	seenMap := make(map[string]bool)
	count := getNestedBagCount("shiny gold", revMap, seenMap)
	if count != 4 {
		t.Errorf("Wrong nested count. Want %d; got %d", 4, count)
	}
}

func TestGetInteriorBagCount(t *testing.T) {
	input, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	bagMap := buildBagMap(input)
	count := getInteriorBagCount("shiny gold", bagMap, 1)
	if count != 32 {
		t.Errorf("Wrong number of bags counted. Want %d; got %d", 32, count)
	}
	input, err = os.Open("./example2.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	bagMap = buildBagMap(input)
	count = getInteriorBagCount("shiny gold", bagMap, 1)
	if count != 126 {
		t.Errorf("Wrong number of bags counted. Want %d; got %d", 126, count)
	}
}

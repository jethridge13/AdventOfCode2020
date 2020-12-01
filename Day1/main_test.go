package main

import (
	"log"
	"os"
	"testing"
)

func TestGetProduct(t *testing.T) {
	input := [2]int{1, 2}
	product := getProduct(input[:])
	if product != 2 {
		t.Errorf("1 * 2 = %d; want 2", product)
	}
}

func TestGetProductNDigits(t *testing.T) {
	input := [4]int{1, 2, 3, 4}
	product := getProduct(input[:])
	if product != 24 {
		t.Errorf("1 * 2 * 3 * 4 = %d; want 24", product)
	}
}

func TestFindSum(t *testing.T) {
	input, err := os.Open("./example.txt")
	defer input.Close()
	if err != nil {
		log.Fatal("Error opening input")
	}
	digits, err := findSum(2020, input)
	if err != nil {
		t.Error("Error finding sum for example input")
	}
	if !(digits[0] == 1721 || digits[1] == 1721) && (digits[0] == 299 || digits[1] == 299) {
		t.Errorf("Error in returned sum. Expected 1721 and 299, got %d and %d", digits[0], digits[1])
	}
}

func TestDay1Answer(t *testing.T) {
	input, err := os.Open("./input.txt")
	defer input.Close()
	if err != nil {
		log.Fatal("Error opening input")
	}
	// Part 1: 805731
	digits, err := findSum(2020, input)
	if err != nil {
		log.Fatal("Error finding sum")
	}
	product := getProduct(digits[:])
	if product != 805731 {
		t.Errorf("Error in Day 1 answer. Expected 805731, got %d", product)
	}
}

func TestPart2Example(t *testing.T) {
	path := "./example.txt"
	digits, err := findSumPartTwo(2020, path)
	if err != nil {
		t.Error("Error in getting Part Two Sum")
	}
	product := getProduct(digits[:])
	if product != 241861950 {
		t.Errorf("Error in Day 1 Part 2 example. Got %d; want 241861950", product)
	}
}

func TestPart2Answer(t *testing.T) {
	// Part 2: 192684960
	path := "./input.txt"
	digits, err := findSumPartTwo(2020, path)
	if err != nil {
		t.Error("Error in getting Part Two Sum")
	}
	product := getProduct(digits[:])
	if product != 192684960 {
		t.Errorf("Error in Day 1 Part 2 example. Got %d; want 241861950", product)
	}
}

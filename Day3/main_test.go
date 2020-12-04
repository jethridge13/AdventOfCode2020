package main

import (
	"log"
	"os"
	"testing"
)

func TestIsTreeHit(t *testing.T) {
	input := "..##......."
	result := isTreeHit(input, 0)
	if result {
		t.Errorf("TestIsTreeHit: False positive: %s %d", input, 0)
	}

	result = isTreeHit(input, 2)
	if !result {
		t.Errorf("TestIsTreeHit: False negative: %s %d", input, 2)
	}

	result = isTreeHit(input, 11)
	if result {
		t.Errorf("TestIsTreeHit: False positive: %s %d", input, 11)
	}

	result = isTreeHit(input, 13)
	if !result {
		t.Errorf("TestIsTreeHit: False negative: %s %d", input, 13)
	}
}

func TestCheckLines(t *testing.T) {
	input, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count := checkLines(input, 3, 1)
	input.Close()
	if count != 7 {
		t.Errorf("Error counting tree hits. Got %d; want %d", count, 7)
	}
}

func TestCheckLinesPart2(t *testing.T) {
	input, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count := checkLines(input, 1, 2)
	input.Close()
	if count != 2 {
		t.Errorf("Error counting tree hits. Got %d; want %d", count, 2)
	}

	input, err = os.Open("./example.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count = checkLines(input, 7, 1)
	input.Close()
	if count != 4 {
		t.Errorf("Error counting tree hits. Got %d; want %d", count, 4)
	}
}

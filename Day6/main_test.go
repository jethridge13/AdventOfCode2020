package main

import (
	"log"
	"os"
	"testing"
)

func TestCountValidYes(t *testing.T) {
	input := make(map[rune]int)
	input['a'] = 1
	input['b'] = 1
	input['c'] = 1
	// abc
	count := countValidYes(input, 0)
	if count != 3 {
		t.Errorf("Got wrong valid yes. Want %d; got %d", 3, count)
	}
	// abc
	count = countValidYes(input, 1)
	if count != 3 {
		t.Errorf("Got wrong valid yes. Want %d; got %d", 3, count)
	}
	// a
	// b
	// c
	count = countValidYes(input, 3)
	if count != 0 {
		t.Errorf("Got wrong valid yes. Want %d; got %d", 0, count)
	}
	// ab
	// ac
	input = make(map[rune]int)
	input['a'] = 2
	input['b'] = 1
	input['c'] = 1
	count = countValidYes(input, 0)
	if count != 3 {
		t.Errorf("Got wrong valid yes. Want %d; got %d", 3, count)
	}
	count = countValidYes(input, 2)
	if count != 1 {
		t.Errorf("Got wrong valid yes. Want %d; got %d", 1, count)
	}
}

func TestCountYesQuestions(t *testing.T) {
	input, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count := countYesQuestions(input, 1)
	if count != 11 {
		t.Errorf("Got wrong yes count. Want %d; got %d", 11, count)
	}
	input.Close()
	input, err = os.Open("./example.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count = countYesQuestions(input, 2)
	if count != 6 {
		t.Errorf("Got wrong yes count. Want %d; got %d", 6, count)
	}
}

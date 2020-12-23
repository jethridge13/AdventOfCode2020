package main

import (
	"testing"
)

func TestParseDecks(t *testing.T) {
	decks := parseDecks("./example.txt")
	if len(decks) != 2 || len(decks[0]) != 5 || len(decks[1]) != 5 {
		t.Errorf("Decks made wrong: %v", decks)
	}
}

func TestCalcScore(t *testing.T) {
	input := []int{3, 2, 10, 6, 8, 5, 9, 4, 7, 1}
	score := calcScore(input)
	exp := 306
	if score != exp {
		t.Errorf("Score wrong. Want %d; got %d", exp, score)
	}

	input = []int{1, 42, 9, 36, 30, 32, 14, 24, 7, 34, 19, 38, 29, 48, 15, 23, 21, 50, 5, 35, 25, 43, 16, 41, 39, 40, 22, 45, 12, 37, 13, 18, 6, 47, 4, 33, 26, 46, 2}
	score = calcScore(input)
	exp = 20389
	if score != exp {
		t.Errorf("Score wrong. Want %d; got %d", exp, score)
	}
}

func TestBattle(t *testing.T) {
	decks := parseDecks("./example.txt")
	winner, score := battle(decks[0], decks[1])
	if winner != 2 || score != 306 {
		t.Errorf("Battle wrong. Winner: Want %d; got %d. Score: Want %d; got %d", 2, winner, 306, score)
	}
}

func TestGetDeckHash(t *testing.T) {
	decks := parseDecks("./example.txt")
	hash := getDeckHash(decks[0], decks[1])
	exp := "P0:92631P1:584710"
	if hash != exp {
		t.Errorf("Wrong hash. Want %s; got %s", exp, hash)
	}
}

func TestRecursiveBattle(t *testing.T) {
	decks := parseDecks("./example.txt")
	winner, score := recursiveBattle(decks[0], decks[1], 1)
	if winner != 2 || score != 291 {
		t.Errorf("Battle wrong. Winner: Want %d; got %d. Score: Want %d; got %d", 2, winner, 291, score)
	}
}

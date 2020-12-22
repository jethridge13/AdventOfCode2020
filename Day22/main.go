package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseDecks(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file")
	}
	scanner := bufio.NewScanner(file)
	decks := make([][]int, 0)
	currDeck := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "Player") >= 0 {
			currDeck = make([]int, 0)
		} else if len(line) == 0 {
			decks = append(decks, currDeck)
		} else {
			digit, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("Could not atoi %s", line)
			}
			currDeck = append(currDeck, digit)
		}
	}
	decks = append(decks, currDeck)
	return decks
}

func calcScore(cards []int) int {
	score := 0
	for index, value := range cards {
		score += (len(cards) - index) * value
	}
	return score
}

func battle(p1 []int, p2 []int) (int, int) {
	for len(p1) > 0 && len(p2) > 0 {
		p1Card := p1[0]
		p1 = p1[1:]
		p2Card := p2[0]
		p2 = p2[1:]
		if p1Card > p2Card {
			p1 = append(p1, p1Card, p2Card)
		} else {
			p2 = append(p2, p2Card, p1Card)
		}
	}
	if len(p1) == 0 {
		return 2, calcScore(p2)
	}
	return 1, calcScore(p1)
}

func getDeckHash(decks ...[]int) string {
	hash := ""
	for _, deck := range decks {
		for _, card := range deck {
			hash += strconv.Itoa(card)
		}
	}
	return hash
}

func recursiveBattle(p1 []int, p2 []int, game int) (int, int) {
	//log.Printf("=== Game %d ===\n", game)
	handHashMap := make(map[string]bool)
	round := 1
	for len(p1) > 0 && len(p2) > 0 {
		// Prevents infinite games of Recusrive Combat
		handHash := getDeckHash(p1, p2)
		if handHashMap[handHash] {
			return 1, calcScore(p1)
		}
		handHashMap[handHash] = true
		//log.Printf("-- Round %d (Game %d) --", round, game)
		//log.Printf("Player 1's deck: %v", p1)
		//log.Printf("Player 2's deck: %v", p2)
		p1Card := p1[0]
		p1 = p1[1:]
		p2Card := p2[0]
		p2 = p2[1:]
		//log.Printf("Player 1 plays: %d", p1Card)
		//log.Printf("Player 2 plays: %d", p2Card)
		winner := 0
		// If both players have at least as many cards remaining in their deck
		// as the value of the card they just drew, the winner of the round is
		// determined by playing a new game of Recursive Combat
		if p1Card <= len(p1) && p2Card <= len(p2) {
			//log.Println("Playing recursive battle")
			p1Copy := make([]int, p1Card)
			copy(p1Copy, p1)
			p2Copy := make([]int, p2Card)
			copy(p2Copy, p2)
			winner, _ = recursiveBattle(p1Copy, p2Copy, game+1)
		}
		if winner == 1 || (winner == 0 && p1Card > p2Card) {
			//log.Printf("Player 1 wins round %d of game %d!", round, game)
			p1 = append(p1, p1Card, p2Card)
		} else {
			//log.Printf("Player 2 wins round %d of game %d!", round, game)
			p2 = append(p2, p2Card, p1Card)
		}
		//log.Println()
		round++
	}
	if len(p1) == 0 {
		return 2, calcScore(p2)
	}
	return 1, calcScore(p1)
}

func main() {
	// Part 1: 31781
	decks := parseDecks("./input.txt")
	_, score := battle(decks[0], decks[1])
	fmt.Println(score)
	// Part 2:
	_, score = recursiveBattle(decks[0], decks[1], 1)
	fmt.Println(score)
}

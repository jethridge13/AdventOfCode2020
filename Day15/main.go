package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// INPUT Day 15 input
var INPUT = "6,19,0,5,7,13,1"

// EXAMPLE1 Day 15 example. Expected outcome: 436
var EXAMPLE1 = "0,3,6"

// EXAMPLE2 Day 15 example. Expected outcome: 1
var EXAMPLE2 = "1,3,2"

// EXAMPLE3 Day 15 example. Expected outcome: 10
var EXAMPLE3 = "2,1,3"

// EXAMPLE4 Day 15 example. Expected outcome: 27
var EXAMPLE4 = "1,2,3"

func fillStartMap(input string) (map[int][]int, int, int) {
	digits := strings.Split(input, ",")
	count := 1
	lastTurnMap := make(map[int][]int)
	for _, digit := range digits {
		n, err := strconv.Atoi(digit)
		if err != nil {
			log.Fatalf("Could not atoi %s", digit)
		}
		lastTurnMap[n] = append(lastTurnMap[n], count)
		count++
	}
	return lastTurnMap, 0, count
}

func getNthNumber(input string, n int) int {
	lastTurnMap, lastTurn, turn := fillStartMap(input)
	for turn < n {
		lastTurnMap[lastTurn] = append(lastTurnMap[lastTurn], turn)
		switch len(lastTurnMap[lastTurn]) {
		case 1:
			lastTurn = 0
		case 2:
			lastTurn = lastTurnMap[lastTurn][1] - lastTurnMap[lastTurn][0]
		case 3:
			lastTurnMap[lastTurn] = lastTurnMap[lastTurn][1:]
			lastTurn = lastTurnMap[lastTurn][1] - lastTurnMap[lastTurn][0]
		}
		turn++
	}
	return lastTurn
}

func main() {
	// Part 1: 468
	res := getNthNumber(INPUT, 2020)
	fmt.Println(res)
	// Part 2: 1801753
	res = getNthNumber(INPUT, 30000000)
	fmt.Println(res)
}

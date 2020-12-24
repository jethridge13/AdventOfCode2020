package main

import "log"

// INPUT Day 23 input
var INPUT = []int{1, 5, 6, 7, 9, 4, 8, 2, 3}

// EXAMPLE Day 23 example
var EXAMPLE = []int{3, 8, 9, 1, 2, 5, 4, 6, 7}

func runRounds(cups []int, rounds int) []int {
	currCup := 0
	for i := 0; i < rounds; i++ {
		log.Printf("-- move %d --", i+1)
		log.Printf("cups: %v", cups)
		log.Printf("currCup: %d at position %d", cups[currCup], currCup)
		// The crab picks up three cups immediately clockwise of the current cup
		removedCups := make([]int, 0)
		indices := []int{currCup + 1, currCup + 2, currCup + 3}
		for index, value := range indices {
			indices[index] = value % len(cups)
		}
		removedCount := 0
		for _, value := range indices {
			index := value - removedCount
			removedCups = append(removedCups, cups[index])
			cups = append(cups[:index], cups[index+1:]...)
			removedCount++
		}
		log.Printf("pick up: %v", removedCups)
		// The crab selects a destination cup: the cup with a label equal to the current cup's label minus one
		currCupValue := cups[currCup]
		currCupValue--
		destCup := -1
		for destCup == -1 {
			for index, value := range cups {
				if value == currCupValue {
					destCup = index
				}
			}
			currCupValue--
			if currCupValue <= 0 {
				currCupValue = 9
			}
		}
		log.Printf("destination: %d\n\n", cups[destCup])
		// The crab places the cups it just picked up clockwise of the destination cup
		tail := make([]int, len(cups[destCup+1:]))
		copy(tail, cups[destCup+1:])
		cups = append(cups[:destCup+1], removedCups...)
		cups = append(cups, tail...)
		if destCup < currCup {
			log.Printf("currCup - destCup :: %d - %d = %d", currCup, destCup, currCup-destCup)
			cups = append(cups[currCup-destCup+1:], cups[:currCup-destCup+1]...)
		}
		// The crab selects a new current cup, which is immediately clockwise of the current cup
		currCup++
		currCup %= len(cups)
	}
	return cups
}

func main() {

}

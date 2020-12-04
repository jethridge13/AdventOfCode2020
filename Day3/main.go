package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isTreeHit(line string, pos int) bool {
	if pos >= len(line) {
		pos = pos % len(line)
	}
	return line[pos] == '#'
}

func checkLines(file *os.File, hStep int, vStep int) int {
	count := 0
	pos := 0
	vPos := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if vPos%vStep != 0 {
			vPos++
			continue
		}
		if isTreeHit(line, pos) {
			count++
		}
		pos += hStep
		vPos++
	}
	return count
}

func main() {
	// Part 1: 200
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count := checkLines(input, 3, 1)
	input.Close()
	fmt.Println(count)

	// Part 2: 3737923200
	// Right 1 Down 1
	input, err = os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count = count * checkLines(input, 1, 1)
	input.Close()
	// Right 5 Down 1
	input, err = os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count = count * checkLines(input, 5, 1)
	input.Close()
	// Right 7 Down 1
	input, err = os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count = count * checkLines(input, 7, 1)
	input.Close()
	// Right 1 Down 2
	input, err = os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count = count * checkLines(input, 1, 2)
	input.Close()
	fmt.Println(count)
}

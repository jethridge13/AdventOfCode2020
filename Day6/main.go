package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countValidYes(groupMap map[rune]int, people int) int {
	count := 0
	if people != 0 {
		for _, value := range groupMap {
			if value == people {
				count++
			}
		}
	} else {
		count += len(groupMap)
	}
	return count
}

func countYesQuestions(file *os.File, part int) int {
	count := 0
	people := 0
	groupMap := make(map[rune]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if part == 2 {
				count += countValidYes(groupMap, people)
			} else {
				count += countValidYes(groupMap, 0)
			}
			people = 0
			groupMap = make(map[rune]int)
			continue
		}
		people++
		for _, char := range line {
			groupMap[char] = groupMap[char] + 1
		}
	}
	count += countValidYes(groupMap, people)
	return count
}

func main() {
	// Part 1: 6551
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count := countYesQuestions(input, 1)
	input.Close()
	fmt.Println(count)
	// Part 2: 3358
	input, err = os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	count = countYesQuestions(input, 2)
	input.Close()
	fmt.Println(count)
}

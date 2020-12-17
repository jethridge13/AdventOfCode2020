package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(path string) ([][2]int, []int, []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file")
	}
	scanner := bufio.NewScanner(file)
	readRanges := true
	readYourTicket := false
	readOtherTickets := false
	ranges := make([][2]int, 0)
	yourTicket := make([]int, 0)
	nearbyTickets := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if readRanges {
			if len(line) == 0 {
				readRanges = false
				readYourTicket = true
				continue
			}
			runes := []rune(line)
			rangeLine := string(runes[strings.Index(line, ":")+1:])
			rangeGroups := strings.Split(rangeLine, " or ")
			for _, grp := range rangeGroups {
				digits := strings.Split(grp, "-")
				min, err := strconv.Atoi(strings.TrimSpace(digits[0]))
				if err != nil {
					log.Fatalf("Could not atoi %s", digits[0])
				}
				max, err := strconv.Atoi(strings.TrimSpace(digits[1]))
				if err != nil {
					log.Fatalf("Could not atoi %s", digits[1])
				}
				newRange := [2]int{min, max}
				ranges = append(ranges, newRange)
			}
		} else if readYourTicket {
			if len(line) == 0 {
				readYourTicket = false
				readOtherTickets = true
				continue
			}
			if line == "your ticket:" {
				continue
			}
			digits := strings.Split(line, ",")
			for _, digit := range digits {
				n, err := strconv.Atoi(digit)
				if err != nil {
					log.Fatalf("Could not atoi %s", digit)
				}
				yourTicket = append(yourTicket, n)
			}
		} else if readOtherTickets {
			if line == "nearby tickets:" {
				continue
			}
			digits := strings.Split(line, ",")
			for _, digit := range digits {
				n, err := strconv.Atoi(digit)
				if err != nil {
					log.Fatalf("Could not atoi %s", digit)
				}
				nearbyTickets = append(nearbyTickets, n)
			}
		}
	}
	file.Close()
	return ranges, yourTicket, nearbyTickets
}

func findInvalidTickets(path string) int {
	ranges, _, otherTickets := parseInput(path)
	invalidSum := 0
	for _, ticket := range otherTickets {
		valid := false
		for _, r := range ranges {
			if ticket >= r[0] && ticket <= r[1] {
				valid = true
				break
			}
		}
		if !valid {
			invalidSum += ticket
		}
	}
	return invalidSum
}

func main() {
	// Part 1: 21996
	res := findInvalidTickets("./input.txt")
	fmt.Println(res)
}

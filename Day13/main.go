package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func minutesUntilArrival(timestamp int, busID int) int {
	return busID - (timestamp % busID)
}

func takeEarliestBus(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file")
	}
	reader := bufio.NewReader(file)
	line, _, _ := reader.ReadLine()
	timestamp, err := strconv.Atoi(string(line))
	if err != nil {
		log.Fatalf("Could not atoi timestamp: %s, %s", line, err)
	}
	busLine, _, _ := reader.ReadLine()
	buses := strings.Split(string(busLine), ",")
	file.Close()
	soonest := math.MaxInt64
	soonestID := -1
	for _, bus := range buses {
		if bus == "x" {
			continue
		}
		busID, err := strconv.Atoi(bus)
		if err != nil {
			log.Fatal("Tried to parse wrong bus")
		}
		minutes := minutesUntilArrival(timestamp, busID)
		if minutes < soonest {
			soonest = minutes
			soonestID = busID
		}
	}
	return soonest * soonestID
}

func findConvergence(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file")
	}
	reader := bufio.NewReader(file)
	_, _, _ = reader.ReadLine()
	line, _, _ := reader.ReadLine()
	file.Close()
	buses := strings.Split(string(line), ",")
	busWithPos := make([][2]int, 0)
	for index, value := range buses {
		if value == "x" {
			continue
		}
		bus, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Could not atoi")
		}
		busWithPos = append(busWithPos, [2]int{bus, index})
	}
	// T such that Ti % Ni == 0
	// For example, T % 7 == 0, T+1 % 19 == 0, T+4 % 59 == 0, etc.
	lcm := 1
	ts := 0
	count := 0
	for true {
		for _, value := range busWithPos {
			check := (ts + value[1]) % value[0]
			// log.Printf("Checking (%d + %d) %d = %d.", ts, value[1], value[0], check)
			if check == 0 {
				lcm *= value[0]
				count++
			} else {
				ts += lcm
				lcm = 1
				count = 0
				break
			}
		}
		if count == len(busWithPos) {
			return ts
		}
	}
	return -1
}

func main() {
	// Part 1: 4808
	result := takeEarliestBus("./input.txt")
	fmt.Println(result)
	// Part 2: 741745043105674
	result = findConvergence("./input.txt")
	fmt.Println(result)
}

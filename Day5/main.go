package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func getSeatIDFromPass(pass string) int {
	rowMin := 0
	rowMax := 127
	columnMin := 0
	columnMax := 7
	for _, value := range pass {
		switch value {
		case 'F':
			rowMax = rowMax - ((rowMax - rowMin) / 2) - 1
		case 'B':
			rowMin = rowMin + ((rowMax - rowMin) / 2) + 1
		case 'L':
			columnMax = columnMax - ((columnMax - columnMin) / 2) - 1
		case 'R':
			columnMin = columnMin + ((columnMax - columnMin) / 2) + 1
		}
	}
	row := rowMin
	column := columnMin
	return row*8 + column
}

func main() {
	// Part 1: 944
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	max := 0
	seatIDs := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		seatID := getSeatIDFromPass(scanner.Text())
		seatIDs = append(seatIDs, seatID)
		if seatID > max {
			max = seatID
		}
	}
	fmt.Println(max)
	// Part 2: 554
	sort.Ints(seatIDs)
	for index, ID := range seatIDs {
		if seatIDs[index+1] == ID+2 {
			fmt.Println(ID + 1)
			break
		}
	}
}

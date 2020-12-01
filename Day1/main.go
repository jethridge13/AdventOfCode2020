package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findSum(target int, file *os.File) ([2]int, error) {
	seenValues := make(map[int]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Error parsing input")
		}
		if seenValues[target-n] {
			return [2]int{n, target - n}, nil
		}
		seenValues[n] = true
	}

	return [2]int{}, errors.New("Could not find sum")
}

func findSumPartTwo(target int, path string) ([3]int, error) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening input.txt")
	}
	// Read all lines into slice
	scanner := bufio.NewScanner(file)
	var lines []int
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Error reading from input")
		}
		lines = append(lines, n)
	}
	log.Print(lines)
	file.Close()
	// Test each number
	for _, value := range lines {
		input, err := os.Open("./input.txt")
		digits, err := findSum(target-value, input)
		input.Close()
		if err == nil {
			return [3]int{value, digits[0], digits[1]}, nil
		}
	}
	return [3]int{}, errors.New("Could not find sum")
}

func getProduct(numbers []int) int {
	product := 1
	for _, value := range numbers {
		product = product * value
	}
	return product
}

func main() {
	input, err := os.Open("./input.txt")
	defer input.Close()
	if err != nil {
		log.Fatal("Error opening input")
	}
	// Part 1: 805731
	digits, err := findSum(2020, input)
	if err != nil {
		log.Fatal("Error finding sum")
	}
	fmt.Println(getProduct(digits[:]))

	// Part 2: 192684960
	path := "./input.txt"
	digitsPartTwo, err := findSumPartTwo(2020, path)
	if err != nil {
		log.Fatal("Error in getting Part Two Sum")
	}
	log.Print(digitsPartTwo)
	fmt.Println(getProduct(digitsPartTwo[:]))
}

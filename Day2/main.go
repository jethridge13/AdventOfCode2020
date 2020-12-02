package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pwdRules struct {
	minCount int
	maxCount int
	letter   rune
	pwd      string
}

func parseLine(line string) (pwdRules, error) {
	parts := strings.Fields(line)
	minAndMax := strings.Split(parts[0], "-")
	min, err := strconv.Atoi(minAndMax[0])
	if err != nil {
		return pwdRules{}, errors.New("Could not parse line")
	}
	max, err := strconv.Atoi(minAndMax[1])
	if err != nil {
		return pwdRules{}, errors.New("Could not parse line")
	}
	letter := []rune(parts[1])[0]
	pwd := parts[2]
	return pwdRules{minCount: min, maxCount: max, letter: letter, pwd: pwd}, nil
}

func validatePwdRules(pwd pwdRules) bool {
	count := 0
	for _, char := range pwd.pwd {
		if char == pwd.letter {
			count++
		}
	}
	if count <= pwd.maxCount && count >= pwd.minCount {
		return true
	}
	return false
}

func validatePwdRulesPart2(pwd pwdRules) bool {
	letterCount := 0
	runeList := []rune(pwd.pwd)
	if runeList[pwd.minCount-1] == pwd.letter {
		letterCount++
	}
	if runeList[pwd.maxCount-1] == pwd.letter {
		letterCount++
	}
	return letterCount == 1
}

func getValidPasswordCount(file *os.File) (int, error) {
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		line, err := parseLine(inputLine)
		if err != nil {
			log.Printf("Could not parse input line: %s", inputLine)
		}
		if validatePwdRules(line) {
			count++
		}
	}
	return count, nil
}

func getValidPasswordCountPart2(file *os.File) (int, error) {
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		line, err := parseLine(inputLine)
		if err != nil {
			log.Printf("Could not parse input line: %s", inputLine)
		}
		if validatePwdRulesPart2(line) {
			count++
		}
	}
	return count, nil
}

func main() {
	// Part 1: 416
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	pwdCount, err := getValidPasswordCount(input)
	input.Close()
	fmt.Println(pwdCount)

	// Part 2: 688
	input, err = os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	pwdCount, err = getValidPasswordCountPart2(input)
	input.Close()
	fmt.Println(pwdCount)
}

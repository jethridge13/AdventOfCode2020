package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func evaluateLine(line string) int {
	stack := []int{0}
	line = strings.TrimSpace(line)
	tokens := strings.ReplaceAll(line, " ", "")
	opStack := []rune{'+'}
	for _, token := range tokens {
		switch token {
		case '+', '*':
			opStack[len(opStack)-1] = token
		case '(':
			stack = append(stack, 0)
			opStack = append(opStack, '+')
		case ')':
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			opStack = opStack[:len(opStack)-1]
			op := opStack[len(opStack)-1]
			switch op {
			case '+':
				stack[len(stack)-1] += n
			case '*':
				stack[len(stack)-1] *= n
			}
		default:
			n, err := strconv.Atoi(string(token))
			if err != nil {
				log.Fatalf("Could not atoi %s", string(token))
			}
			switch opStack[len(opStack)-1] {
			case '+':
				stack[len(stack)-1] += n
			case '*':
				stack[len(stack)-1] *= n
			}
		}
	}
	return stack[0]
}

func evaluateLinePart2(line string) int {
	postfix := infixToPostfix(line)
	infix := postfixToInfix(postfix)
	return evaluateLine(infix)
}

func infixToPostfix(input string) string {
	input = strings.ReplaceAll(input, " ", "")
	outputRunes := make([]rune, 0)
	stack := make([]rune, 0)
	for _, char := range input {
		switch char {
		case '+':
			for len(stack) > 0 && stack[len(stack)-1] == '+' {
				token := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				outputRunes = append(outputRunes, token)
			}
			stack = append(stack, '+')
		case '*':
			if len(stack) == 0 {
				stack = append(stack, '*')
			} else {
				for len(stack) > 0 && stack[len(stack)-1] != '(' {
					token := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					outputRunes = append(outputRunes, token)
				}
				stack = append(stack, '*')
			}
		case '(':
			stack = append(stack, '(')
		case ')':
			token := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for token != '(' && len(stack) > 0 {
				outputRunes = append(outputRunes, token)
				token = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		default:
			outputRunes = append(outputRunes, rune(char))
		}
	}
	for len(stack) > 0 {
		token := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		outputRunes = append(outputRunes, token)
	}
	return string(outputRunes)
}

func postfixToInfix(input string) string {
	input = strings.ReplaceAll(input, " ", "")
	stack := make([]string, 0)
	for _, char := range input {
		switch char {
		case '+', '*':
			token := "(" + stack[len(stack)-2] + string(char) + stack[len(stack)-1] + ")"
			stack = stack[:len(stack)-2]
			stack = append(stack, token)
		default:
			stack = append(stack, string(char))
		}
	}
	return stack[0]
}

func main() {
	// Part 1: 202553439706
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		res := evaluateLine(scanner.Text())
		sum += res
	}
	file.Close()
	fmt.Println(sum)
	// Part 2: 88534268715686
	file, err = os.Open("./input.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	scanner = bufio.NewScanner(file)
	sum = 0
	for scanner.Scan() {
		res := evaluateLinePart2(scanner.Text())
		sum += res
	}
	file.Close()
	fmt.Println(sum)
}

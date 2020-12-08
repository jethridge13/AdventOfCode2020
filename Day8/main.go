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

type ins struct {
	op  string
	val int
}

func parseIns(line string) ins {
	var instruction ins
	parts := strings.Split(line, " ")
	instruction.op = parts[0]
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("Could not atoi")
	}
	instruction.val = value
	return instruction
}

func parseInsSet(path string) []ins {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file %s", path)
	}
	instructions := make([]ins, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inst := parseIns(line)
		instructions = append(instructions, inst)
	}
	file.Close()
	return instructions
}

func switchNopAndJmp(inst ins) ins {
	var newInst ins
	newInst.val = inst.val
	if inst.op == "nop" {
		newInst.op = "jmp"
	} else {
		newInst.op = "nop"
	}
	return newInst
}

func runWithInfiniteLoopFix(insts []ins) (int, error) {
	instsToCheck := make([]int, 0)
	for index, inst := range insts {
		if inst.op == "nop" || inst.op == "jmp" {
			instsToCheck = append(instsToCheck, index)
		}
	}
	for _, value := range instsToCheck {
		newInst := switchNopAndJmp(insts[value])
		insts[value] = newInst
		res, err := run(insts)
		if err == nil {
			// log.Println(value)
			return res, nil
		}
		revertInst := switchNopAndJmp(newInst)
		insts[value] = revertInst
	}
	return -1, errors.New("Could not fix")
}

func run(insts []ins) (int, error) {
	acc := 0
	ptr := 0
	foundIndices := make(map[int]bool)
	for true {
		if foundIndices[ptr] {
			// Infinite loop detected
			return acc, errors.New("Infinite loop detected")
		}
		if ptr == len(insts) {
			// Normal execution ends
			return acc, nil
		}
		foundIndices[ptr] = true
		inst := insts[ptr]
		switch inst.op {
		case "acc":
			acc += inst.val
			ptr++
		case "jmp":
			ptr += inst.val
		default:
			ptr++
		}
	}
	return acc, nil
}

func main() {
	// Part 1: 1600
	insts := parseInsSet("./input.txt")
	res, _ := run(insts)
	fmt.Println(res)
	// Part 2: 1543
	res, err := runWithInfiniteLoopFix(insts)
	if err != nil {
		log.Fatal("Infinite loop detected")
	}
	fmt.Println(res)
}

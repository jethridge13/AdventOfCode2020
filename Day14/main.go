package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func combineMaskPart2(mask string, memAddr string) string {
	maskedRunes := make([]byte, 0)
	memAddrInt, err := strconv.Atoi(memAddr)
	if err != nil {
		log.Fatalf("Could not atoi %s", memAddr)
	}
	binValue := fmt.Sprintf("%036b", memAddrInt)
	for index := range binValue {
		digit := binValue[index]
		if mask[index] == '1' {
			digit = '1'
		} else if mask[index] == 'X' {
			digit = 'X'
		}
		maskedRunes = append(maskedRunes, digit)
	}
	return string(maskedRunes)
}

func getAllPossibleAddresses(combined string) []int64 {
	index := strings.Index(combined, "X")
	if index == -1 {
		combinedInt, err := strconv.ParseInt(combined, 2, 64)
		if err != nil {
			log.Fatal("Could not convert from binary")
		}
		list := make([]int64, 1)
		list[0] = combinedInt
		return list
	}
	// Replace X with 0
	bin0Runes := []rune(combined)
	bin0Runes[index] = '0'
	bin0 := string(bin0Runes)
	list0 := getAllPossibleAddresses(bin0)
	// Replace X with 1
	bin1Runes := []rune(combined)
	bin1Runes[index] = '1'
	bin1 := string(bin1Runes)
	list1 := getAllPossibleAddresses(bin1)
	return append(append([]int64{}, list0...), list1...)
}

func runProgram(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file")
	}
	scanner := bufio.NewScanner(file)
	mask := ""
	mem := make(map[int]int64)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal("Could not atoi value")
			}
			runes := []rune(parts[0])
			memAddr, err := strconv.Atoi(string(runes[4 : len(runes)-1]))
			if err != nil {
				log.Fatal("Could not atoi mem")
			}
			binValue := fmt.Sprintf("%036b", value)
			maskedRunes := make([]byte, 0)
			for index := range binValue {
				digit := binValue[index]
				if mask[index] != 'X' {
					digit = mask[index]
				}
				maskedRunes = append(maskedRunes, digit)
			}
			maskedValue := string(maskedRunes)
			maskedRes, err := strconv.ParseInt(maskedValue, 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			mem[memAddr] = maskedRes
		}
	}
	file.Close()
	count := int64(0)
	for _, value := range mem {
		count += value
	}
	return count
}

func runProgramPart2(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file")
	}
	scanner := bufio.NewScanner(file)
	mask := ""
	mem := make(map[int64]int)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal("Could not atoi value")
			}
			runes := []rune(parts[0])
			memAddr := string(runes[4 : len(runes)-1])
			if err != nil {
				log.Fatal("Could not atoi mem")
			}
			combined := combineMaskPart2(mask, memAddr)
			addresses := getAllPossibleAddresses(combined)
			for _, addr := range addresses {
				mem[addr] = value
			}
		}
	}
	file.Close()
	count := 0
	for _, value := range mem {
		count += value
	}
	return count
}

func main() {
	// Part 1: 15514035145260
	res := runProgram("./input.txt")
	fmt.Println(res)
	// Part 2: 3926790061594
	res2 := runProgramPart2("./input.txt")
	fmt.Println(res2)
}

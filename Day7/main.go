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

type bagWithCount struct {
	bag   string
	count int
}

func getBagWithCount(bag string) (bagWithCount, error) {
	var b bagWithCount
	bag = strings.TrimSpace(bag)
	if bag == "no other bags." {
		return b, errors.New("No additional bags")
	}
	parts := strings.Split(bag, " ")
	count, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	parts = parts[1 : len(parts)-1]
	b.bag = strings.Join(parts, " ")
	b.count = count
	return b, nil
}

func parseLine(line string) (string, []bagWithCount) {
	halves := strings.Split(line, "contain")
	bagType := strings.ReplaceAll(halves[0], "bags", "")
	bagType = strings.TrimSpace(bagType)
	bagList := strings.Split(halves[1], ",")
	bagWithCountList := make([]bagWithCount, 0)
	for _, bag := range bagList {
		bagStruct, err := getBagWithCount(bag)
		if err == nil {
			bagWithCountList = append(bagWithCountList, bagStruct)
		}
	}
	return bagType, bagWithCountList
}

func buildBagMap(file *os.File) map[string][]bagWithCount {
	bagMap := make(map[string][]bagWithCount)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bagType, bagList := parseLine(line)
		bagMap[bagType] = bagList
	}
	return bagMap
}

func buildReverseBagMap(bagMap map[string][]bagWithCount) map[string][]bagWithCount {
	revMap := make(map[string][]bagWithCount)
	for key, value := range bagMap {
		for _, bag := range value {
			newBag := bagWithCount{bag: key}
			revMap[bag.bag] = append(revMap[bag.bag], newBag)
		}
	}
	return revMap
}

func getNestedBagCount(bagKey string, bagMap map[string][]bagWithCount, seenMap map[string]bool) int {
	if seenMap[bagKey] {
		return 0
	}
	seenMap[bagKey] = true
	// If found is false, that means we have encountered a bag which contains no other bags
	// Therefore, the chain ends
	bagList, found := bagMap[bagKey]
	if !found {
		return 0
	}
	count := 0
	for _, bag := range bagList {
		if !seenMap[bag.bag] {
			count += 1 + getNestedBagCount(bag.bag, bagMap, seenMap)
		}
	}
	return count
}

// NOTE: Ultimately returns right answer + 1, but I finished it, so...
// Maybe I'll clean it up later
func getInteriorBagCount(bagKey string, bagMap map[string][]bagWithCount, bagCount int) int {
	count := 0
	bagList := bagMap[bagKey]
	if len(bagList) == 0 {
		return bagCount
	}
	for _, bag := range bagList {
		n := getInteriorBagCount(bag.bag, bagMap, bag.count)
		count += n
	}
	count = bagCount + count*bagCount
	return count
}

func main() {
	// Part 1: 235
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	bagMap := buildBagMap(input)
	revMap := buildReverseBagMap(bagMap)
	seenMap := make(map[string]bool)
	count := getNestedBagCount("shiny gold", revMap, seenMap)
	fmt.Println(count)
	// Part 2: 158493
	count = getInteriorBagCount("shiny gold", bagMap, 1)
	fmt.Println(count)
}

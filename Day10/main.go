package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func buildList(path string) ([]int, error) {
	list := make([]int, 0)
	file, err := os.Open(path)
	if err != nil {
		return list, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return list, err
		}
		list = append(list, n)
	}
	sort.Ints(list)
	return list, nil
}

func getDiffs(list []int) map[int]int {
	diffs := make(map[int]int)
	for index, value := range list {
		if index == 0 {
			diffs[value]++
		} else {
			diff := value - list[index-1]
			diffs[diff]++
		}
	}
	// The end is always 3
	diffs[3]++
	return diffs
}

func getListOfOneDiffs(list []int) []int {
	diffs := make([]int, 0)
	count := 0
	for index, value := range list {
		if index == 0 {
			if value == 1 {
				count++
			} else {
				diffs = append(diffs, count)
				count = 0
			}
		} else {
			if value-list[index-1] == 1 {
				count++
			} else {
				diffs = append(diffs, count)
				count = 0
			}
		}
	}
	diffs = append(diffs, count)
	return diffs
}

func getTribonacci(n int) int {
	list := []int{1, 1, 2, 4, 7, 13, 24, 44}
	return list[n]
}

func getProduct(list []int) int {
	n := 1
	for _, value := range list {
		n *= getTribonacci(value)
	}
	return n
}

func main() {
	// Part 1: 2048
	list, err := buildList("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	diffs := getDiffs(list)
	fmt.Println(diffs[1] * diffs[3])
	// Part 2: 1322306994176
	res := getListOfOneDiffs(list)
	prod := getProduct(res)
	fmt.Println(prod)
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type xmas struct {
	sumCount int
	recentN  []int
	nMap     map[int]bool
}

func copyXmas(src xmas) xmas {
	newX := xmas{sumCount: src.sumCount}
	copy(newX.recentN, src.recentN)
	newX = initXmasNmap(newX)
	return newX
}

func initXmasNmap(x xmas) xmas {
	if len(x.nMap) == 0 {
		x.nMap = make(map[int]bool)
		for _, value := range x.recentN {
			x.nMap[value] = true
		}
	}
	return x
}

func addToXmas(x xmas, n int) xmas {
	newX := copyXmas(x)
	newX.recentN = append(x.recentN, n)
	newX = initXmasNmap(newX)
	newX.nMap[n] = true
	if len(newX.recentN) > newX.sumCount {
		poppedN := newX.recentN[0]
		delete(newX.nMap, poppedN)
		newX.recentN = newX.recentN[1:]
	}
	return newX
}

func verifyN(x xmas, check int) bool {
	x = initXmasNmap(x)
	for _, n := range x.recentN {
		dif := check - n
		if x.nMap[dif] && dif != check-dif {
			// log.Printf("%d validates with %d and %d", check, dif, n)
			return true
		}
	}
	return false
}

func findFirstInvalidN(path string, sumCheckCount int) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	x := xmas{sumCount: sumCheckCount}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		digit, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Could not atoi")
		}
		if len(x.recentN) == sumCheckCount {
			res := verifyN(x, digit)
			if !res {
				return digit, nil
			}
		}
		x = addToXmas(x, digit)
	}
	return -1, errors.New("All numbers valid")
}

func findContiguousSet(path string, sumCheckCount int) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	digits := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		digit, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Could not atoi")
		}
		digits = append(digits, digit)
	}
	file.Close()
	target, err := findFirstInvalidN(path, sumCheckCount)
	if err != nil {
		return -1, err
	}
	strPtr := 0
	endPtr := 1
	count := digits[strPtr] + digits[endPtr]
	for strPtr != endPtr {
		if count == target {
			// Find min and max, then sum
			set := digits[strPtr:endPtr]
			min := math.MaxInt64
			max := 0
			for _, i := range set {
				if i < min {
					min = i
				}
				if i > max {
					max = i
				}
			}
			return min + max, nil
		} else if count > target {
			count -= digits[strPtr]
			strPtr++
		} else {
			endPtr++
			count += digits[endPtr]
		}
	}
	return -1, errors.New("Could not find set")
}

func main() {
	// Part 1: 25918798
	res, _ := findFirstInvalidN("./input.txt", 25)
	fmt.Println(res)
	// Part 2: 3340942
	res, _ = findContiguousSet("./input.txt", 25)
	fmt.Println(res)
}

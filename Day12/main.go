package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ins struct {
	op  rune
	val int
}

func parseIns(line string) ins {
	var instruction ins
	runes := []rune(line)
	instruction.op = runes[0]
	val, err := strconv.Atoi(string(runes[1:]))
	if err != nil {
		log.Fatal("Could not parse instruction")
	}
	instruction.val = val
	return instruction
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func navigate(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open %s", path)
	}
	scanner := bufio.NewScanner(file)
	x := 0
	y := 0
	deg := 90
	for scanner.Scan() {
		nav := parseIns(scanner.Text())
		switch nav.op {
		case 'N':
			y += nav.val
		case 'S':
			y -= nav.val
		case 'E':
			x += nav.val
		case 'W':
			x -= nav.val
		case 'L':
			deg -= nav.val
			deg %= 360
		case 'R':
			deg += nav.val
			deg %= 360
		case 'F':
			switch deg {
			case 0:
				y += nav.val
			case 90, -270:
				x += nav.val
			case 180, -180:
				y -= nav.val
			case 270, -90:
				x -= nav.val
			default:
				log.Fatalf("Incorrect degree: %d", deg)
			}
		}
	}
	file.Close()
	return abs(x) + abs(y)
}

func navigateWaypoint(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open %s", path)
	}
	scanner := bufio.NewScanner(file)
	shipX := 0
	shipY := 0
	wayX := 10
	wayY := 1
	for scanner.Scan() {
		nav := parseIns(scanner.Text())
		switch nav.op {
		case 'N':
			wayY += nav.val
		case 'S':
			wayY -= nav.val
		case 'E':
			wayX += nav.val
		case 'W':
			wayX -= nav.val
		case 'L':
			deg := nav.val % 360
			switch deg {
			case 90, -270:
				temp := wayX
				wayX = -wayY
				wayY = temp
			case 180, -180:
				wayX = -wayX
				wayY = -wayY
			case 270, -90:
				temp := -wayX
				wayX = wayY
				wayY = temp
			}
		case 'R':
			deg := nav.val % 360
			switch deg {
			case 90, -270:
				temp := -wayX
				wayX = wayY
				wayY = temp
			case 180, -180:
				wayX = -wayX
				wayY = -wayY
			case 270, -90:
				temp := wayX
				wayX = -wayY
				wayY = temp
			}
		case 'F':
			shipX += wayX * nav.val
			shipY += wayY * nav.val
		}
	}
	file.Close()
	return abs(shipX) + abs(shipY)
}

func main() {
	// Part 1: 1133
	res := navigate("./input.txt")
	fmt.Println(res)
	// Part 2: 61053
	res = navigateWaypoint("./input.txt")
	fmt.Println(res)
}

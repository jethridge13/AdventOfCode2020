package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func buildGrid(path string) [][]rune {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open %s", path)
	}
	scanner := bufio.NewScanner(file)
	grid := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		grid = append(grid, runes)
	}
	file.Close()
	return grid
}

func getNeighborSeatCount(grid [][]rune, i int, j int) int {
	seats := 0
	// Check North neighbors
	if i-1 >= 0 {
		// Check NW
		if j-1 >= 0 && grid[i-1][j-1] == '#' {
			seats++
		}
		// Check N
		if grid[i-1][j] == '#' {
			seats++
		}
		// Check NE
		if j+1 < len(grid[i]) && grid[i-1][j+1] == '#' {
			seats++
		}
	}
	// Check direct West neighbor
	if j-1 >= 0 {
		if grid[i][j-1] == '#' {
			seats++
		}
	}
	// Check direct East neighbor
	if j+1 < len(grid[i]) {
		if grid[i][j+1] == '#' {
			seats++
		}
	}
	// Check South neighbors
	if i+1 < len(grid) {
		// Check SW
		if j-1 >= 0 && grid[i+1][j-1] == '#' {
			seats++
		}
		// Check S
		if grid[i+1][j] == '#' {
			seats++
		}
		// Check SE
		if j+1 < len(grid[i]) && grid[i+1][j+1] == '#' {
			seats++
		}
	}
	return seats
}

func getNeighborSeatCountPart2(grid [][]rune, i int, j int) int {
	seats := 0
	// Check NW
	iC := i - 1
	jC := j - 1
	for iC >= 0 && jC >= 0 {
		if grid[iC][jC] == '#' {
			seats++
			iC = -1
			jC = -1
		} else if grid[iC][jC] == 'L' {
			iC = -1
			jC = -1
		} else {
			iC--
			jC--
		}
	}
	// Check N
	iC = i - 1
	jC = j
	for iC >= 0 {
		if grid[iC][jC] == '#' {
			seats++
			iC = -1
			jC = -1
		} else if grid[iC][jC] == 'L' {
			iC = -1
			jC = -1
		} else {
			iC--
		}
	}
	// Check NE
	iC = i - 1
	jC = j + 1
	for iC >= 0 && jC < len(grid[iC]) {
		if grid[iC][jC] == '#' {
			seats++
			iC = -1
			jC = -1
		} else if grid[iC][jC] == 'L' {
			iC = -1
			jC = -1
		} else {
			iC--
			jC++
		}
	}
	// Check W
	iC = i
	jC = j - 1
	for jC >= 0 {
		if grid[iC][jC] == '#' {
			seats++
			iC = -1
			jC = -1
		} else if grid[iC][jC] == 'L' {
			iC = -1
			jC = -1
		} else {
			jC--
		}
	}
	// Check E
	iC = i
	jC = j + 1
	for iC >= 0 && jC < len(grid[i]) {
		if grid[iC][jC] == '#' {
			seats++
			iC = -1
			jC = -1
		} else if grid[iC][jC] == 'L' {
			iC = -1
			jC = -1
		} else {
			jC++
		}
	}
	// Check SW
	iC = i + 1
	jC = j - 1
	for iC < len(grid) && jC >= 0 {
		if grid[iC][jC] == '#' {
			seats++
			iC = -1
			jC = -1
		} else if grid[iC][jC] == 'L' {
			iC = -1
			jC = -1
		} else {
			iC++
			jC--
		}
	}
	// Check S
	iC = i + 1
	jC = j
	for iC < len(grid) && jC >= 0 {
		if grid[iC][jC] == '#' {
			seats++
			iC = -1
			jC = -1
		} else if grid[iC][jC] == 'L' {
			iC = -1
			jC = -1
		} else {
			iC++
		}
	}
	// Check SW
	iC = i + 1
	jC = j + 1
	for iC < len(grid) && jC < len(grid[i]) && iC >= 0 && jC >= 0 {
		if grid[iC][jC] == '#' {
			seats++
			iC = -1
			jC = -1
		} else if grid[iC][jC] == 'L' {
			iC = -1
			jC = -1
		} else {
			iC++
			jC++
		}
	}
	return seats
}

func applyRules(grid [][]rune, part int) int {
	changed := 0
	gridCopy := make([][]rune, len(grid))
	for i := range grid {
		gridCopy[i] = make([]rune, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}
	for i, iV := range grid {
		for j, jV := range iV {
			if jV == '.' {
				continue
			}
			var seats int
			if part == 2 {
				seats = getNeighborSeatCountPart2(gridCopy, i, j)
			} else {
				seats = getNeighborSeatCount(gridCopy, i, j)
			}
			if jV == 'L' {
				if seats == 0 {
					grid[i][j] = '#'
					changed++
				}
			} else if jV == '#' {
				if (part == 1 && seats >= 4) || (part == 2 && seats >= 5) {
					grid[i][j] = 'L'
					changed++
				}
			}
		}
	}
	return changed
}

func applyUntilStable(grid [][]rune) {
	for true {
		count := applyRules(grid, 1)
		if count == 0 {
			return
		}
	}
}

func applyUntilStablePart2(grid [][]rune) {
	for true {
		count := applyRules(grid, 2)
		if count == 0 {
			return
		}
	}
}

func countOccupiedSeats(grid [][]rune) int {
	count := 0
	for _, line := range grid {
		for _, value := range line {
			if value == '#' {
				count++
			}
		}
	}
	return count
}

func main() {
	// Part 1: 2270
	grid := buildGrid("./input.txt")
	applyUntilStable(grid)
	count := countOccupiedSeats(grid)
	fmt.Println(count)
	// Part 2: 2042
	grid = buildGrid("./input.txt")
	applyUntilStablePart2(grid)
	count = countOccupiedSeats(grid)
	fmt.Println(count)
}

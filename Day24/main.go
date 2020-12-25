package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countTiles(m map[[2]int]bool) int {
	count := 0
	for _, value := range m {
		if value {
			count++
		}
	}
	return count
}

func placeTiles(path string) map[[2]int]bool {
	// https://www.reddit.com/r/gamedev/comments/5fh3wr/hexagon_grid_what_data_structure_to_use/
	s := ""
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file")
	}
	reader := bufio.NewReader(file)
	coords := [2]int{0, 0}
	tiles := make(map[[2]int]bool)
	for true {
		r, _, err := reader.ReadRune()
		if err != nil {
			log.Println("Reached end")
			log.Println(err)
			break
		}
		switch r {
		case 'e':
			// East
			coords[0]++
			s += "e"
		case 'w':
			// West
			coords[0]--
			s += "w"
		case 'n', 's':
			if r == 'n' {
				coords[1]--
				s += "n"
			} else {
				coords[1]++
				s += "s"
			}
			r2, _, err := reader.ReadRune()
			if err != nil {
				log.Println(err)
				log.Fatal("Encountered bad rune")
			}
			switch r2 {
			case 'e':
				if (coords[1]-1)%2 != 0 {
					coords[0]++
				}
				s += "e"
			case 'w':
				if (coords[1]-1)%2 == 0 {
					coords[0]--
				}
				s += "w"
			}
		default:
			tiles[coords] = !tiles[coords]
			coords[0] = 0
			coords[1] = 0
		}
	}
	tiles[coords] = !tiles[coords]
	file.Close()
	return tiles
}

func main() {
	// Part 1: 277
	fmt.Println(countTiles(placeTiles("./input.txt")))
}

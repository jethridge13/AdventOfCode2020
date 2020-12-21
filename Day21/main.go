package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func parseLine(line string) ([]string, []string) {
	allergenPart := line[strings.Index(line, "(contains ")+10 : strings.Index(line, ")")]
	ingredientPart := line[:strings.Index(line, " (")]
	ingredients := strings.Split(ingredientPart, " ")
	allergens := strings.Split(allergenPart, ", ")
	return ingredients, allergens
}

func makeSet(list []string) map[string]bool {
	set := make(map[string]bool)
	for _, s := range list {
		set[s] = true
	}
	return set
}

func setUnion(u1 map[string]bool, u2 map[string]bool) map[string]bool {
	union := make(map[string]bool)
	for key := range u1 {
		if u2[key] {
			union[key] = true
		}
	}
	return union
}

func setDifference(u1 map[string]bool, u2 map[string]bool) map[string]bool {
	diff := make(map[string]bool)
	for key := range u1 {
		if !u2[key] {
			diff[key] = true
		}
	}
	return diff
}

func findNonAllergens(path string) (int, map[string]map[string]bool) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file")
	}
	scanner := bufio.NewScanner(file)
	allIngredients := make(map[string]bool)
	allergenSets := make(map[string]map[string]bool)
	allergenSeen := make(map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()
		ingredients, allergens := parseLine(line)
		for _, i := range ingredients {
			allIngredients[i] = true
		}
		ingredientSet := makeSet(ingredients)
		for _, allergen := range allergens {
			if !allergenSeen[allergen] {
				allergenSets[allergen] = ingredientSet
				allergenSeen[allergen] = true
			} else {
				allergenSets[allergen] = setUnion(ingredientSet, allergenSets[allergen])
			}
		}
	}
	file.Close()
	nonAllergens := allIngredients
	for _, set := range allergenSets {
		nonAllergens = setDifference(nonAllergens, set)
	}
	count := 0
	file, err = os.Open(path)
	if err != nil {
		log.Fatal("Could not open file")
	}
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		incredients, _ := parseLine(line)
		for _, i := range incredients {
			if nonAllergens[i] {
				count++
			}
		}
	}
	file.Close()
	return count, allergenSets
}

func isolateAllergens(m map[string]map[string]bool) map[string]map[string]bool {
	singles := make([]string, 0)
	for key, value := range m {
		if len(value) == 1 {
			singles = append(singles, key)
		}
	}
	if len(singles) == len(m) {
		return m
	}
	for key := range m {
		for _, single := range singles {
			if key == single {
				continue
			}
			m[key] = setDifference(m[key], m[single])
		}
	}
	return isolateAllergens(m)
}

func findDangerousList(m map[string]map[string]bool) string {
	list := ""
	m = isolateAllergens(m)
	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	seenMap := make(map[string]bool)
	for _, key := range keys {
		for ingredient := range m[key] {
			if !seenMap[ingredient] {
				if len(list) == 0 {
					list = ingredient
				} else {
					list = list + "," + ingredient
				}
				seenMap[ingredient] = true
			}
		}
	}
	return list
}

func main() {
	// Part 1: 1945
	count, allergenSets := findNonAllergens("./input.txt")
	fmt.Println(count)
	// Part 2: pgnpx,srmsh,ksdgk,dskjpq,nvbrx,khqsk,zbkbgp,xzb
	list := findDangerousList(allergenSets)
	fmt.Println(list)
}

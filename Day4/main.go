package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func parseInput(file *os.File) [][]string {
	passports := make([][]string, 0)
	pp := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			passports = append(passports, pp)
			pp = make([]string, 0)
		}
		pp = append(pp, line)
	}
	passports = append(passports, pp)
	return passports
}

func parsePassport(lines []string) (passport, error) {
	pp := passport{}
	for _, line := range lines {
		parts := strings.Fields(line)
		for _, part := range parts {
			field := strings.Split(part, ":")
			var err error
			switch field[0] {
			case "byr":
				pp.byr = field[1]
			case "iyr":
				pp.iyr = field[1]
			case "eyr":
				pp.eyr = field[1]
			case "hgt":
				pp.hgt = field[1]
			case "hcl":
				pp.hcl = field[1]
			case "ecl":
				pp.ecl = field[1]
			case "pid":
				pp.pid = field[1]
			case "cid":
				pp.cid = field[1]
			}
			if err != nil {
				return pp, err
			}
		}
	}
	return pp, nil
}

func parsePassportFromInput(passports [][]string) ([]passport, error) {
	pps := make([]passport, 0)
	for _, pp := range passports {
		newPP, err := parsePassport(pp)
		if err != nil {
			return pps, err
		}
		pps = append(pps, newPP)
	}
	return pps, nil
}

func looseValidatePassport(pp passport) bool {
	if pp.byr == "" || pp.iyr == "" || pp.eyr == "" || pp.hgt == "" || pp.hcl == "" || pp.ecl == "" || pp.pid == "" {
		return false
	}
	return true
}

func strictValidatePassport(pp passport) bool {
	if !looseValidatePassport(pp) {
		return false
	}
	// Validate byr
	byr, err := strconv.Atoi(pp.byr)
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}
	// Validate iyr
	iyr, err := strconv.Atoi(pp.iyr)
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}
	// Validate eyr
	eyr, err := strconv.Atoi(pp.eyr)
	if err != nil {
		return false
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}
	// Validate hgt
	units := pp.hgt[len(pp.hgt)-2:]
	length, err := strconv.Atoi(pp.hgt[:len(pp.hgt)-2])
	if err != nil {
		return false
	}
	if units == "cm" {
		if length < 150 || length > 193 {
			return false
		}
	} else if units == "in" {
		if length < 59 || length > 76 {
			return false
		}
	} else {
		return false
	}
	// Validate hcl
	if pp.hcl[0] != '#' {
		return false
	}
	_, err = hex.DecodeString(pp.hcl[1:])
	if err != nil {
		return false
	}
	// Validate ecl
	switch pp.ecl {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false
	}
	// Validate pid
	if len(pp.pid) != 9 {
		return false
	}
	// Validate cid
	return true
}

func countValidPassports(passports []passport, strict bool) int {
	count := 0
	if strict {
		for _, pp := range passports {
			if looseValidatePassport(pp) {
				count++
			}
		}
	} else {
		for _, pp := range passports {
			if strictValidatePassport(pp) {
				count++
			}
		}
	}
	return count
}

func main() {
	// Part 1: 245
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	passports := parseInput(input)
	parsedPassports, err := parsePassportFromInput(passports)
	if err != nil {
		log.Fatal("Encountered error when parsing passports: %s", err)
	}
	count := countValidPassports(parsedPassports, false)
	fmt.Println(count)
	// Part 2: 133
	count = countValidPassports(parsedPassports, true)
	fmt.Println(count)
}

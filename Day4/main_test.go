package main

import (
	"log"
	"os"
	"testing"
)

func TestParsePassport(t *testing.T) {
	testLine := []string{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm"}
	expectedOutput := passport{ecl: "gry", pid: "860033327", eyr: "2020", hcl: "#fffffd", byr: "1937", iyr: "2017", cid: "147", hgt: "183cm"}
	output, err := parsePassport(testLine)
	if err != nil {
		t.Errorf("Encountered error when parsing passport")
	}
	if output != expectedOutput {
		t.Errorf("Passport parsing failed: got %+v; expected %+v", output, expectedOutput)
	}
}

func TestParseInput(t *testing.T) {
	input, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	passports := parseInput(input)
	if len(passports) != 4 {
		t.Errorf("Incorrect number of passports parsed. Got %d; want %d", len(passports), 4)
	}
}

func TestParsePassportFromInputs(t *testing.T) {
	input, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	passports := parseInput(input)
	parsedPassports, err := parsePassportFromInput(passports)
	if err != nil {
		t.Errorf("Encountered error when parsing passports")
	}
	if len(parsedPassports) != 4 {
		t.Errorf("Incorrect number of passports parsed. Got %d; want %d", len(passports), 4)
	}
}

func TestCountValidPassports(t *testing.T) {
	input, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	passports := parseInput(input)
	parsedPassports, err := parsePassportFromInput(passports)
	if err != nil {
		t.Errorf("Encountered error when parsing passports")
	}
	count := countValidPassports(parsedPassports, false)
	if count != 2 {
		t.Errorf("Incorrect number of valid passports counted. Got %d; want %d", count, 2)
	}
}

package main

import "testing"

func TestParseLine(t *testing.T) {
	input := "1-3 a: abcde"
	parsedLine, err := parseLine(input)
	if err != nil {
		t.Errorf("Could not parse line")
	}
	if parsedLine.minCount != 1 ||
		parsedLine.maxCount != 3 ||
		parsedLine.letter != 'a' ||
		parsedLine.pwd != "abcde" {
		t.Errorf("Incorrect pwd struct returned: %v", parsedLine)
	}
}

func TestValidatePwdRules(t *testing.T) {
	input := pwdRules{minCount: 1, maxCount: 3, letter: 'a', pwd: "abcde"}
	valid := validatePwdRules(input)
	if !valid {
		t.Errorf("Error validating pwd: %v", input)
	}
}

func TestValidatePwdRulesPart2(t *testing.T) {
	input := pwdRules{minCount: 1, maxCount: 3, letter: 'a', pwd: "abcde"}
	valid := validatePwdRulesPart2(input)
	if !valid {
		t.Errorf("Error validating pwd: %v", input)
	}

	input = pwdRules{minCount: 1, maxCount: 3, letter: 'b', pwd: "cdefg"}
	valid = validatePwdRulesPart2(input)
	if valid {
		t.Errorf("Error validating pwd: %v", input)
	}
}

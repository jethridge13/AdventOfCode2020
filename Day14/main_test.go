package main

import (
	"testing"
)

func TestRunProgram(t *testing.T) {
	res := runProgram("./example.txt")
	if res != 165 {
		t.Errorf("Program memory wrong. Want %d; got %d", 165, res)
	}
}

func TestCombineMaskPart2(t *testing.T) {
	mask := "000000000000000000000000000000X1001X"
	addr := "42"
	res := combineMaskPart2(mask, addr)
	expected := "000000000000000000000000000000X1101X"
	if res != expected {
		t.Errorf("Mask combined wrong. Want %s, got %s", expected, res)
	}

	mask = "00000000000000000000000000000000X0XX"
	addr = "26"
	res = combineMaskPart2(mask, addr)
	expected = "00000000000000000000000000000001X0XX"
	if res != expected {
		t.Errorf("Mask combined wrong. Want %s, got %s", expected, res)
	}
}

func TestGetAllPossibleAddresses(t *testing.T) {
	mask := "000000000000000000000000000000X1001X"
	addr := "42"
	combined := combineMaskPart2(mask, addr)
	res := getAllPossibleAddresses(combined)
	if len(res) != 4 {
		t.Errorf("Length of list wrong")
	}

	mask = "00000000000000000000000000000000X0XX"
	addr = "26"
	combined = combineMaskPart2(mask, addr)
	res = getAllPossibleAddresses(combined)
	if len(res) != 8 {
		t.Errorf("Length of list wrong")
	}
}

func TestRunProgramPart2(t *testing.T) {
	res := runProgramPart2("./example2.txt")
	if res != 208 {
		t.Errorf("Program wrong. Want %d; got %d", 208, res)
	}
}

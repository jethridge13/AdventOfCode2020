package main

import "testing"

func TestParseIns(t *testing.T) {
	input := "nop +0"
	resIns := parseIns(input)
	if resIns.op != "nop" || resIns.val != 0 {
		t.Errorf("Wrong ins parsed: %+v", resIns)
	}

	input = "acc -99"
	resIns = parseIns(input)
	if resIns.op != "acc" || resIns.val != -99 {
		t.Errorf("Wrong ins parsed: %+v", resIns)
	}
}

func TestParseInsSet(t *testing.T) {
	insts := parseInsSet("./example.txt")
	if len(insts) != 9 {
		t.Errorf("Wrong # of instructions parsed")
	}
	if insts[0].op != "nop" || insts[0].val != 0 {
		t.Errorf("First instruction wrong: %+v", insts[0])
	}
	if insts[8].op != "acc" || insts[8].val != 6 {
		t.Errorf("Last instruction wrong: %+v", insts[8])
	}
}

func TestRun(t *testing.T) {
	insts := parseInsSet("./example.txt")
	res, err := run(insts)
	if err == nil || res != 5 {
		t.Errorf("Inf loop detection wrong. Want %d; got %d", 5, res)
	}

	insts = parseInsSet("./example2.txt")
	res, err = run(insts)
	if err != nil || res != 8 {
		t.Errorf("Normal run wrong. Want %d, got %d", 8, res)
	}
}

func TestRunWithInfiniteLoopFix(t *testing.T) {
	insts := parseInsSet("./example.txt")
	res, err := runWithInfiniteLoopFix(insts)
	if err != nil {
		t.Errorf("Could not fix inf loop")
	}
	if res != 8 {
		t.Errorf("Run with fix wrong. Want %d; got %d", 8, res)
	}
}

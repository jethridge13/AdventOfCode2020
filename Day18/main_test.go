package main

import (
	"testing"
)

func TestEvaluateLine(t *testing.T) {
	input := "1 + 2 * 3 + 4 * 5 + 6"
	res := evaluateLine(input)
	exp := 71
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}

	input = "1 + (2 * 3) + (4 * (5 + 6))"
	res = evaluateLine(input)
	exp = 51
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}

	input = "2 * 3 + (4 * 5)"
	res = evaluateLine(input)
	exp = 26
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}

	input = "5 + (8 * 3 + 9 + 3 * 4 * 3)"
	res = evaluateLine(input)
	exp = 437
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}

	input = "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	res = evaluateLine(input)
	exp = 12240
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}

	input = "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	res = evaluateLine(input)
	exp = 13632
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}
}

func TestInfixToPostfix(t *testing.T) {
	input := "1 + 2 * 3 + 4 * 5 + 6"
	res := infixToPostfix(input)
	exp := "12+34+*56+*"
	if res != exp {
		t.Errorf("Infix to postfix wrong. Want %s; got %s", exp, res)
	}

	input = "1 + (2 * 3) + (4 * (5 + 6))"
	res = infixToPostfix(input)
	exp = "123*+456+*+"
	if res != exp {
		t.Errorf("Infix to postfix wrong. Want %s; got %s", exp, res)
	}
}

func TestPostfixToInfex(t *testing.T) {
	input := "12+34+*56+*"
	res := postfixToInfix(input)
	exp := "(((1+2)*(3+4))*(5+6))"
	if res != exp {
		t.Errorf("Infix to postfix wrong. Want %s; got %s", exp, res)
	}
}

func TestEvaluateLinePart2(t *testing.T) {
	input := "1 + 2 * 3 + 4 * 5 + 6"
	res := evaluateLinePart2(input)
	exp := 231
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}

	input = "1 + (2 * 3) + (4 * (5 + 6))"
	res = evaluateLinePart2(input)
	exp = 51
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}

	input = "2 * 3 + (4 * 5)"
	res = evaluateLinePart2(input)
	exp = 46
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}

	input = "5 + (8 * 3 + 9 + 3 * 4 * 3)"
	res = evaluateLinePart2(input)
	exp = 1445
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}

	input = "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	res = evaluateLinePart2(input)
	exp = 669060
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}

	input = "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	res = evaluateLinePart2(input)
	exp = 23340
	if res != exp {
		t.Errorf("Evaluation wrong. Want %d; got %d", exp, res)
	}
}

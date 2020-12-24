package main

import (
	"log"
	"testing"
)

func TestRunRounds(t *testing.T) {
	res := runRounds(EXAMPLE, 10)
	log.Println(res)
}

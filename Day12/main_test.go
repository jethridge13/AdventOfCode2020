package main

import "testing"

func TestParseIns(t *testing.T) {
	res := parseIns("F10")
	if res.op != 'F' || res.val != 10 {
		t.Errorf("Could not parse instruction: %+v", res)
	}
}

func TestNavigate(t *testing.T) {
	res := navigate("./example.txt")
	if res != 25 {
		t.Errorf("We're lost at sea! Got %d; want %d", res, 25)
	}
}

func TestNavigateWaypoint(t *testing.T) {
	res := navigateWaypoint("./example.txt")
	if res != 286 {
		t.Errorf("We're lost at sea following a waypoint! Got %d; want %d", res, 286)
	}
}

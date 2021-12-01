package main

import "testing"

func TestIncreased(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 7
	real := Increased(input)
	if real != expected {
		t.Errorf("Expected %v and real %v)", expected, real)
	}
}

func TestSlidingWindows(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 5
	real := SlidingWindows(input)
	if real != expected {
		t.Errorf("Expected %v and real %v)", expected, real)
	}
}

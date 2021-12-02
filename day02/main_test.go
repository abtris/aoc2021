package main

import "testing"

func TestPart1(t *testing.T) {
	expected := 150
	real := processPart1("input_test")
	if real != expected {
		t.Errorf("Expected %v and real %v)", expected, real)
	}
}

func TestPart2(t *testing.T) {
	expected := 900
	real := processPart2("input_test")
	if real != expected {
		t.Errorf("Expected %v and real %v)", expected, real)
	}
}

package main

import "testing"

func TestPart1(t *testing.T) {
	input := readFile("test_input.txt")
	puzzle := parseInput(input)

	expected := 2
	result := part1(puzzle)
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"123 328  51 64",
		"45 64  387 23",
		"6 98  215 314",
		"*   +   *   +",
	}
	expected := 4277556
	result := part1(input)
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}
	expected := 3263827
	result := part2(input)
	if result != expected {
		t.Errorf("part2() = %d, want %d", result, expected)
	}
}

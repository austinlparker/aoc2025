package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
	expected := 3
	result := part1(input)
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
	expected := 6
	result := part2(input)
	if result != expected {
		t.Errorf("part2() = %d, want %d", result, expected)
	}
}

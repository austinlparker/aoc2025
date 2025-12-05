package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		" ",
		"1",
		"5",
		"8",
		"11",
		"17",
		"32",
	}
	expected := 3
	result := part1(input)
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		" ",
		"1",
		"5",
		"8",
		"11",
		"17",
		"32",
	}
	expected := 14
	result := part2(input)
	if result != expected {
		t.Errorf("part2() = %d, want %d", result, expected)
	}
}

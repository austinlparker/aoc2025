package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
	expected := 357
	result := part1(input)
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
	expected := 3121910778619
	result := part2(input)
	if result != expected {
		t.Errorf("part2() = %d, want %d", result, expected)
	}
}

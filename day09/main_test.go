package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"7,1",
		"11,1",
		"11,7",
		"9,7",
		"9,5",
		"2,5",
		"2,3",
		"7,3",
	}
	expected := 50
	result := part1(input)
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"7,1",
		"11,1",
		"11,7",
		"9,7",
		"9,5",
		"2,5",
		"2,3",
		"7,3",
	}
	expected := 24
	result := part2(input)
	if result != expected {
		t.Errorf("part2() = %d, want %d", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	lines := readInput("input.txt")
	b.ResetTimer()
	for range b.N {
		part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := readInput("input.txt")
	b.ResetTimer()
	for range b.N {
		part2(lines)
	}
}

package main

import (
	"fmt"
	"maps"
	"os"
	"strings"
)

func main() {
	lines := readInput("input.txt")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
}

func part1(lines []string) int {
	startRow, startCol := 0, 0
	for row, line := range lines {
		if col := strings.Index(line, "S"); col >= 0 {
			startRow, startCol = row, col
			break
		}
	}

	rays := map[int]bool{startCol: true}
	splits := 0

	for _, line := range lines[startRow+1:] {
		newRays := make(map[int]bool)

		for col := range rays {
			if col >= 0 && col < len(line) && line[col] == '^' {
				splits++
				newRays[col-1] = true
				newRays[col+1] = true
			} else {
				newRays[col] = true
			}
		}

		rays = newRays
	}

	return splits
}

func part2(lines []string) int {
	startRow, startCol := 0, 0
	for row, line := range lines {
		if col := strings.Index(line, "S"); col >= 0 {
			startRow, startCol = row, col
			break
		}
	}

	rays := map[int]int{startCol: 1}

	for _, line := range lines[startRow+1:] {
		newRays := make(map[int]int)

		for col, count := range maps.All(rays) {
			if col < 0 || col >= len(line) {
				continue
			}

			if line[col] == '^' {
				if col-1 >= 0 {
					newRays[col-1] += count
				}
				if col+1 < len(line) {
					newRays[col+1] += count
				}
			} else {
				newRays[col] += count
			}
		}

		rays = newRays
	}

	total := 0
	for _, count := range rays {
		total += count
	}
	return total
}

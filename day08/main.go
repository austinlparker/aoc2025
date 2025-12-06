package main

import (
	"fmt"
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
	// TODO: implement
	return 0
}

func part2(lines []string) int {
	// TODO: implement
	return 0
}

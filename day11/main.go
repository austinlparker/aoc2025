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

func parseGraph(lines []string) map[string][]string {
	graph := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}
		name := parts[0]
		connections := strings.Fields(parts[1])
		graph[name] = connections
	}
	return graph
}

func countPaths(graph map[string][]string, current, target string, memo map[string]int) int {
	if current == target {
		return 1
	}

	key := current + "->" + target
	if val, ok := memo[key]; ok {
		return val
	}

	neighbors, exists := graph[current]
	if !exists {
		return 0
	}

	total := 0
	for _, neighbor := range neighbors {
		total += countPaths(graph, neighbor, target, memo)
	}

	memo[key] = total
	return total
}

func part1(lines []string) int {
	graph := parseGraph(lines)
	memo := make(map[string]int)
	return countPaths(graph, "you", "out", memo)
}

func part2(lines []string) int {
	graph := parseGraph(lines)
	memo := make(map[string]int)

	case1 := countPaths(graph, "svr", "dac", memo) *
		countPaths(graph, "dac", "fft", memo) *
		countPaths(graph, "fft", "out", memo)

	case2 := countPaths(graph, "svr", "fft", memo) *
		countPaths(graph, "fft", "dac", memo) *
		countPaths(graph, "dac", "out", memo)

	return case1 + case2
}

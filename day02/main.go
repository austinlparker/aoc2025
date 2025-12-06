package main

import (
	"fmt"
	"os"
	"strconv"
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
	total := 0
	for idRange := range strings.SplitSeq(lines[0], ",") {
		startStr, endStr, _ := strings.Cut(idRange, "-")
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)
		for i := start; i <= end; i++ {
			if isRepeatingSequence(strconv.Itoa(i)) {
				total += i
			}
		}
	}
	return total
}

func isRepeatingSequence(s string) bool {
	n := len(s)
	if n <= 1 || n%2 != 0 {
		return false
	}

	period := n / 2
	for i := range period {
		if s[i] != s[i+period] {
			return false
		}
	}
	return true
}

func part2(lines []string) int {
	total := 0
	for idRange := range strings.SplitSeq(lines[0], ",") {
		startStr, endStr, _ := strings.Cut(idRange, "-")
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)
		for i := start; i <= end; i++ {
			if isRepeatingSequenceAtLeastTwice(strconv.Itoa(i)) {
				total += i
			}
		}
	}
	return total
}

func isRepeatingSequenceAtLeastTwice(s string) bool {
	n := len(s)
	if n <= 1 {
		return false
	}
	for period := 1; period <= n/2; period++ {
		if n%period != 0 {
			continue
		}
		pattern := s[:period]
		if strings.Repeat(pattern, n/period) == s {
			return true
		}
	}
	return false
}

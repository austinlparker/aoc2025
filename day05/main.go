package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := readInput("input.txt")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

type Range struct {
	start, end int
}

func (r Range) Contains(num int) bool {
	return num >= r.start && num <= r.end
}

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseInput(lines []string) ([]Range, []int) {
	var rangeLines []string
	var numbers []int

	inRanges := true
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			inRanges = false
			continue
		}
		if inRanges {
			rangeLines = append(rangeLines, line)
		} else {
			num, _ := strconv.Atoi(line)
			numbers = append(numbers, num)
		}
	}

	return buildRanges(rangeLines), numbers
}

func buildRanges(rangeStrs []string) []Range {
	var ranges []Range
	for _, rangeStr := range rangeStrs {
		parts := strings.Split(rangeStr, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, Range{start, end})
	}
	return ranges
}

func part1(lines []string) int {
	ranges, numbers := parseInput(lines)

	count := 0
	for _, num := range numbers {
		for _, r := range ranges {
			if r.Contains(num) {
				count++
				break
			}
		}
	}

	return count
}

func part2(lines []string) int {
	ranges, _ := parseInput(lines)
	count := 0

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []Range{ranges[0]}
	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]
		if r.start <= last.end+1 {
			if r.end > last.end {
				last.end = r.end
			}
		} else {
			merged = append(merged, r)
		}
	}

	for _, r := range merged {
		count += r.end - r.start + 1
	}

	return count
}

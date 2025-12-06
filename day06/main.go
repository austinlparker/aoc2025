package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
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

func parseInput(lines []string) ([][]int, []string) {
	lastLine := lines[len(lines)-1]
	var operators []string
	for _, op := range lastLine {
		if strings.ContainsRune("*+", op) {
			operators = append(operators, string(op))
		}
	}

	columns := make([][]int, len(operators))

	for _, line := range lines[:len(lines)-1] {
		for i, field := range strings.Fields(line) {
			num, _ := strconv.Atoi(field)
			columns[i] = append(columns[i], num)
		}
	}

	return columns, operators
}

func parseInputRTL(lines []string) ([][]int, []string) {
	maxLen := 0
	for _, line := range lines {
		maxLen = max(maxLen, len(line))
	}

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = slices.Repeat([]rune{' '}, maxLen)
		copy(grid[i], []rune(line))
	}

	lastRow := len(lines) - 1

	type opInfo struct {
		col int
		op  string
	}
	var operators []opInfo
	for col := range maxLen {
		ch := grid[lastRow][col]
		if ch == '*' || ch == '+' {
			operators = append(operators, opInfo{col: col, op: string(ch)})
		}
	}

	isSeparator := func(col int) bool {
		for row := range lines {
			if grid[row][col] != ' ' {
				return false
			}
		}
		return true
	}

	var results [][]int
	var ops []string

	for _, op := range operators {
		endCol := op.col
		for endCol < maxLen && !isSeparator(endCol) {
			endCol++
		}

		var numbers []int
		for col := endCol - 1; col >= op.col; col-- {
			var numStr strings.Builder
			for row := range lastRow {
				ch := grid[row][col]
				if unicode.IsDigit(ch) {
					numStr.WriteRune(ch)
				}
			}
			if numStr.Len() > 0 {
				num, _ := strconv.Atoi(numStr.String())
				numbers = append(numbers, num)
			}
		}

		results = append(results, numbers)
		ops = append(ops, op.op)
	}

	return results, ops
}

func part1(lines []string) int {
	columns, operators := parseInput(lines)
	total := 0
	for i, col := range columns {
		result := col[0]
		for _, n := range col[1:] {
			if operators[i] == "*" {
				result *= n
			} else {
				result += n
			}
		}
		total += result
	}
	return total
}

func part2(lines []string) int {
	columns, operators := parseInputRTL(lines)
	total := 0
	for i, col := range columns {
		result := col[0]
		for _, n := range col[1:] {
			if operators[i] == "*" {
				result *= n
			} else {
				result += n
			}
		}
		total += result
	}
	return total
}

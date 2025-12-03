package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readInput("input.txt")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
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

func part1(lines []string) int {
	var total int64
	for _, line := range lines {
		total += largestKDigitNumber(line, 2)
	}
	return int(total)
}

func part2(lines []string) int {
	var total int64
	for _, line := range lines {
		total += largestKDigitNumber(line, 12)
	}
	return int(total)
}

func largestKDigitNumber(s string, k int) int64 {
	var result int64
	lastIndex := -1

	for i := range k {
		start := lastIndex + 1
		end := len(s) - k + i

		maxDigit := int64(0)
		for j := start; j <= end; j++ {
			d := int64(s[j] - '0')
			if d > maxDigit {
				maxDigit = d
				lastIndex = j
			}
		}
		result = result*10 + maxDigit
	}
	return result
}

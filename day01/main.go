package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	password := 0
	current := 50

	for _, item := range lines {
		direction := item[0]
		number, _ := strconv.Atoi(item[1:])
		switch direction {
		case 'L':
			current = (current - number + 100) % 100
		case 'R':
			current = (current + number) % 100
		}
		if current == 0 {
			password++
		}
	}
	return password
}

func part2(lines []string) int {
	password := 0
	current := 50

	for _, item := range lines {
		direction := item[0]
		number, _ := strconv.Atoi(item[1:])

		start := current
		password += number / 100
		remainder := number % 100

		switch direction {
		case 'L':
			current = (current - remainder + 100) % 100
			if remainder >= start && start > 0 {
				password++
			} else if current == 0 {
				password++
			}
		case 'R':
			current = (current + remainder) % 100
			if start+remainder > 100 {
				password++
			} else if current == 0 {
				password++
			}
		}
	}
	return password
}

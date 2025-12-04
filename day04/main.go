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

func parseGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func countNeighbors(grid [][]rune, r, c int, match rune) int {
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	count := 0
	for _, d := range directions {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
			if grid[nr][nc] == match {
				count++
			}
		}
	}
	return count
}

func part1(lines []string) int {
	grid := parseGrid(lines)
	count := 0

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '@' && countNeighbors(grid, r, c, '@') < 4 {
				count++
			}
		}
	}
	return count
}

func part2(lines []string) int {
	grid := parseGrid(lines)
	count := 0

	for {
		var toRemove [][2]int
		for r := range grid {
			for c := range grid[r] {
				if grid[r][c] == '@' && countNeighbors(grid, r, c, '@') < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}
		if len(toRemove) == 0 {
			break
		}
		for _, pos := range toRemove {
			grid[pos[0]][pos[1]] = '.'
		}
		count += len(toRemove)
	}
	return count
}

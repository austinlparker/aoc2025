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

type Point struct {
	x, y int
}

func parsePoints(lines []string) []Point {
	points := make([]Point, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		points[i].x, _ = strconv.Atoi(parts[0])
		points[i].y, _ = strconv.Atoi(parts[1])
	}
	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func area(p1, p2 Point) int {
	return (abs(p1.x-p2.x) + 1) * (abs(p1.y-p2.y) + 1)
}

func part1(lines []string) int {
	points := parsePoints(lines)

	maxArea := 0
	for i := range len(points) {
		for j := i + 1; j < len(points); j++ {
			maxArea = max(maxArea, area(points[i], points[j]))
		}
	}
	return maxArea
}

type Outline map[Point]bool

func (o Outline) paint(from, to Point) {
	if from.x == to.x {
		for y := min(from.y, to.y); y <= max(from.y, to.y); y++ {
			o[Point{from.x, y}] = true
		}
	} else if from.y == to.y {
		for x := min(from.x, to.x); x <= max(from.x, to.x); x++ {
			o[Point{x, from.y}] = true
		}
	}
}

func buildOutline(points []Point) Outline {
	outline := make(Outline)
	n := len(points)
	for i := range n {
		outline.paint(points[i], points[(i+1)%n])
	}
	return outline
}

func part2(lines []string) int {
	points := parsePoints(lines)
	outline := buildOutline(points)

	maxArea := 0
	for i, a := range points {
	outer:
		for j, b := range points {
			if i >= j {
				continue
			}

			candidateArea := area(a, b)
			if candidateArea <= maxArea {
				continue
			}

			xLo, xHi := min(a.x, b.x), max(a.x, b.x)
			yLo, yHi := min(a.y, b.y), max(a.y, b.y)

			for k, c := range points {
				if k == i || k == j {
					continue
				}
				if xLo < c.x && c.x < xHi && yLo < c.y && c.y < yHi {
					continue outer
				}
			}

			for r, x := range [2]int{xLo, xHi} {
				for y := yLo + 1; y <= yHi-1; y++ {
					if outline[Point{x + 1 - 2*r, y}] {
						continue outer
					}
				}
			}

			for s, y := range [2]int{yLo, yHi} {
				for x := xLo + 1; x <= xHi-1; x++ {
					if outline[Point{x, y + 1 - 2*s}] {
						continue outer
					}
				}
			}

			maxArea = candidateArea
		}
	}
	return maxArea
}

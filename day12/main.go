package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct{ Row, Col int }

type Shape struct {
	Index  int
	Points []Point
}

type Region struct {
	Width, Height int
	Quantities    []int
}

type PuzzleInput struct {
	Shapes  []Shape
	Regions []Region
}

func main() {
	input := parseInput(readFile("input.txt"))
	fmt.Println("Part 1:", part1(input))
}

func readFile(filename string) string {
	data, _ := os.ReadFile(filename)
	return strings.TrimSpace(string(data))
}

func parseInput(data string) PuzzleInput {
	var input PuzzleInput
	blocks := strings.Split(data, "\n\n")
	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		firstLine := lines[0]
		if strings.Contains(strings.Split(firstLine, ":")[0], "x") {
			for _, line := range lines {
				input.Regions = append(input.Regions, parseRegion(line))
			}
		} else {
			input.Shapes = append(input.Shapes, parseShape(lines))
		}
	}
	return input
}

func parseShape(lines []string) Shape {
	idx, _ := strconv.Atoi(strings.TrimSuffix(lines[0], ":"))
	var points []Point
	for row, line := range lines[1:] {
		for col, ch := range line {
			if ch == '#' {
				points = append(points, Point{row, col})
			}
		}
	}
	return Shape{Index: idx, Points: points}
}

func parseRegion(line string) Region {
	parts := strings.SplitN(line, ":", 2)
	dims := strings.Split(parts[0], "x")
	w, _ := strconv.Atoi(dims[0])
	h, _ := strconv.Atoi(dims[1])
	var quantities []int
	for _, q := range strings.Fields(parts[1]) {
		n, _ := strconv.Atoi(q)
		quantities = append(quantities, n)
	}
	return Region{w, h, quantities}
}

func (s Shape) GetAllOrientations() [][]Point {
	seen := make(map[string]bool)
	var result [][]Point
	points := s.Points
	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			norm := normalize(points)
			key := toKey(norm)
			if !seen[key] {
				seen[key] = true
				result = append(result, norm)
			}
			points = rotate(points)
		}
		points = flipH(s.Points)
	}
	return result
}

func rotate(pts []Point) []Point {
	out := make([]Point, len(pts))
	for i, p := range pts {
		out[i] = Point{p.Col, -p.Row}
	}
	return out
}

func flipH(pts []Point) []Point {
	out := make([]Point, len(pts))
	for i, p := range pts {
		out[i] = Point{p.Row, -p.Col}
	}
	return out
}

func normalize(pts []Point) []Point {
	minR, minC := pts[0].Row, pts[0].Col
	for _, p := range pts {
		if p.Row < minR {
			minR = p.Row
		}
		if p.Col < minC {
			minC = p.Col
		}
	}
	out := make([]Point, len(pts))
	for i, p := range pts {
		out[i] = Point{p.Row - minR, p.Col - minC}
	}
	return out
}

func toKey(pts []Point) string {
	sorted := make([]Point, len(pts))
	copy(sorted, pts)
	for i := range sorted {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j].Row < sorted[i].Row ||
				(sorted[j].Row == sorted[i].Row && sorted[j].Col < sorted[i].Col) {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	var b strings.Builder
	for _, p := range sorted {
		fmt.Fprintf(&b, "%d,%d;", p.Row, p.Col)
	}
	return b.String()
}

type Piece struct {
	ShapeIdx   int
	Placements [][]uint64
}

func canFitRegion(region Region, shapes []Shape) bool {
	totalCells := 0
	for shapeIdx, qty := range region.Quantities {
		totalCells += qty * len(shapes[shapeIdx].Points)
	}
	if totalCells > region.Width*region.Height {
		return false
	}

	gridSize := region.Width * region.Height
	numWords := (gridSize + 63) / 64

	var pieces []Piece
	for shapeIdx, qty := range region.Quantities {
		if qty == 0 {
			continue
		}
		orientations := shapes[shapeIdx].GetAllOrientations()
		var placements [][]uint64

		for _, orient := range orientations {
			maxR, maxC := 0, 0
			for _, p := range orient {
				if p.Row > maxR {
					maxR = p.Row
				}
				if p.Col > maxC {
					maxC = p.Col
				}
			}
			for row := 0; row <= region.Height-maxR-1; row++ {
				for col := 0; col <= region.Width-maxC-1; col++ {
					mask := make([]uint64, numWords)
					for _, p := range orient {
						bit := (row+p.Row)*region.Width + (col + p.Col)
						mask[bit/64] |= 1 << (bit % 64)
					}
					placements = append(placements, mask)
				}
			}
		}

		for q := 0; q < qty; q++ {
			pieces = append(pieces, Piece{shapeIdx, placements})
		}
	}

	if len(pieces) == 0 {
		return true
	}

	sort.SliceStable(pieces, func(i, j int) bool {
		return len(pieces[i].Placements) < len(pieces[j].Placements)
	})

	grid := make([]uint64, numWords)
	minIdx := make([]int, len(pieces))
	return solve(grid, pieces, 0, minIdx)
}

func solve(grid []uint64, pieces []Piece, idx int, minIdx []int) bool {
	if idx == len(pieces) {
		return true
	}

	piece := pieces[idx]
	startFrom := 0

	if idx > 0 && pieces[idx-1].ShapeIdx == piece.ShapeIdx {
		startFrom = minIdx[idx-1] + 1
	}

	for i := startFrom; i < len(piece.Placements); i++ {
		mask := piece.Placements[i]
		if canPlace(grid, mask) {
			place(grid, mask)
			minIdx[idx] = i
			if solve(grid, pieces, idx+1, minIdx) {
				return true
			}
			unplace(grid, mask)
		}
	}
	return false
}

func canPlace(grid, mask []uint64) bool {
	for i := range grid {
		if grid[i]&mask[i] != 0 {
			return false
		}
	}
	return true
}

func place(grid, mask []uint64) {
	for i := range grid {
		grid[i] |= mask[i]
	}
}

func unplace(grid, mask []uint64) {
	for i := range grid {
		grid[i] &^= mask[i]
	}
}

func part1(input PuzzleInput) int {
	count := 0
	for _, region := range input.Regions {
		if canFitRegion(region, input.Shapes) {
			count++
		}
	}
	return count
}

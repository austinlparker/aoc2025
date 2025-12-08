package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := readInput("input.txt")
	fmt.Println("Part 1:", part1(lines, 1000))
	fmt.Println("Part 2:", part2(lines))
}

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
}

type Point3D struct {
	X, Y, Z int
}

func parsePoint(line string) Point3D {
	parts := strings.Split(line, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return Point3D{x, y, z}
}

func (p Point3D) Distance(other Point3D) float64 {
	dx := float64(p.X - other.X)
	dy := float64(p.Y - other.Y)
	dz := float64(p.Z - other.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type Edge struct {
	From, To int
	Dist     float64
}

type UnionFind struct {
	parent []int
	rank   []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   make([]int, n),
	}
	for i := range n {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return
	}

	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.parent[py] = px
	uf.size[px] += uf.size[py]
	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
}

func part1(lines []string, numConnections int) int {
	points := make([]Point3D, len(lines))
	for i, line := range lines {
		points[i] = parsePoint(line)
	}
	edges := make([]Edge, 0, len(points)*(len(points)-1)/2)
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{
				From: i,
				To:   j,
				Dist: points[i].Distance(points[j]),
			})
		}
	}

	slices.SortFunc(edges, func(a, b Edge) int {
		if a.Dist < b.Dist {
			return -1
		}
		if a.Dist > b.Dist {
			return 1
		}
		return 0
	})

	uf := NewUnionFind(len(points))
	for i := range min(numConnections, len(edges)) {
		uf.Union(edges[i].From, edges[i].To)
	}

	circuitSizes := make(map[int]int)
	for i := range points {
		root := uf.Find(i)
		circuitSizes[root] = uf.size[root]
	}

	sizes := make([]int, 0, len(circuitSizes))
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}
	slices.Sort(sizes)
	slices.Reverse(sizes)

	result := 1
	for i := range min(3, len(sizes)) {
		result *= sizes[i]
	}

	return result
}

func part2(lines []string) int {
	points := make([]Point3D, len(lines))
	for i, line := range lines {
		points[i] = parsePoint(line)
	}

	edges := make([]Edge, 0, len(points)*len(points)-1/2)
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{
				From: i,
				To:   j,
				Dist: points[i].Distance(points[j]),
			})
		}
	}

	slices.SortFunc(edges, func(a, b Edge) int {
		if a.Dist < b.Dist {
			return -1
		}
		if a.Dist > b.Dist {
			return 1
		}
		return 0
	})

	uf := NewUnionFind(len(points))
	numComponents := len(points)
	result := 0

	for _, edge := range edges {
		if uf.Find(edge.From) != uf.Find(edge.To) {
			uf.Union(edge.From, edge.To)
			numComponents--
			if numComponents == 1 {
				result = points[edge.From].X * points[edge.To].X
			}
		}
	}
	return result
}

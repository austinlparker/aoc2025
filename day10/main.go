package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Machine represents a single machine configuration.
// All lights start OFF - Target indicates which should be ON.
type Machine struct {
	Target   int   // bitmask of desired end state
	Buttons  []int // each button's toggle mask
	Joltages []int // joltage requirements per light
}

func main() {
	machines := parseInput(readInput("input.txt"))
	fmt.Println("Part 1:", part1(machines))
	fmt.Println("Part 2:", part2(machines))
}

func readInput(filename string) []string {
	data, _ := os.ReadFile(filename)
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func parseInput(lines []string) []Machine {
	machines := make([]Machine, len(lines))
	for i, line := range lines {
		machines[i] = parseMachine(line)
	}
	return machines
}

func parseMachine(line string) Machine {
	brackets := extract(line, '[', ']')
	braces := extract(line, '{', '}')

	var target int
	for i, ch := range brackets {
		if ch == '#' {
			target |= 1 << i
		}
	}

	var buttons []int
	for rest := line; ; {
		if start := strings.IndexByte(rest, '('); start >= 0 {
			end := strings.IndexByte(rest[start:], ')') + start
			buttons = append(buttons, toBitmask(rest[start+1:end]))
			rest = rest[end+1:]
		} else {
			break
		}
	}

	return Machine{target, buttons, parseInts(braces)}
}

func extract(s string, open, close byte) string {
	start := strings.IndexByte(s, open)
	return s[start+1 : strings.IndexByte(s[start:], close)+start]
}

func toBitmask(s string) int {
	mask := 0
	for _, p := range strings.Split(s, ",") {
		n, _ := strconv.Atoi(p)
		mask |= 1 << n
	}
	return mask
}

func parseInts(s string) []int {
	parts := strings.Split(s, ",")
	nums := make([]int, len(parts))
	for i, p := range parts {
		nums[i], _ = strconv.Atoi(p)
	}
	return nums
}

func minPresses(m Machine) int {
	if m.Target == 0 {
		return 0
	}

	visited := map[int]bool{0: true}
	queue := []int{0}

	for presses := 1; len(queue) > 0; presses++ {
		var next []int
		for _, state := range queue {
			for _, btn := range m.Buttons {
				s := state ^ btn
				if s == m.Target {
					return presses
				}
				if !visited[s] {
					visited[s] = true
					next = append(next, s)
				}
			}
		}
		queue = next
	}

	return -1
}

func minPressesJoltage(m Machine) int {
	n := len(m.Joltages)
	numBtns := len(m.Buttons)

	// Build augmented matrix [A|b]
	matrix := make([][]float64, n)
	for j := range n {
		matrix[j] = make([]float64, numBtns+1)
		matrix[j][numBtns] = float64(m.Joltages[j])
		for i, btn := range m.Buttons {
			if btn&(1<<j) != 0 {
				matrix[j][i] = 1
			}
		}
	}

	// Gaussian elimination to RREF
	pivotCol := 0
	pivotRows := make([]int, numBtns)
	for i := range pivotRows {
		pivotRows[i] = -1
	}

	for row := 0; row < n && pivotCol < numBtns; {
		maxRow := row
		for r := row + 1; r < n; r++ {
			if math.Abs(matrix[r][pivotCol]) > math.Abs(matrix[maxRow][pivotCol]) {
				maxRow = r
			}
		}

		if math.Abs(matrix[maxRow][pivotCol]) < 1e-9 {
			pivotCol++
			continue
		}

		matrix[row], matrix[maxRow] = matrix[maxRow], matrix[row]

		scale := matrix[row][pivotCol]
		for c := range numBtns + 1 {
			matrix[row][c] /= scale
		}

		for r := range n {
			if r != row && math.Abs(matrix[r][pivotCol]) > 1e-9 {
				factor := matrix[r][pivotCol]
				for c := range numBtns + 1 {
					matrix[r][c] -= factor * matrix[row][c]
				}
			}
		}

		pivotRows[pivotCol] = row
		row++
		pivotCol++
	}

	// Identify free variables
	var freeVars []int
	for col := range numBtns {
		if pivotRows[col] == -1 {
			freeVars = append(freeVars, col)
		}
	}

	// Compute bounds for free variables based on non-negativity of basic vars
	maxFreeVal := 0
	for _, t := range m.Joltages {
		if t > maxFreeVal {
			maxFreeVal = t
		}
	}

	bestSum := math.MaxInt

	var search func(idx int, freeVals []int, currentSum int)
	search = func(idx int, freeVals []int, currentSum int) {
		if currentSum >= bestSum {
			return
		}

		if idx == len(freeVars) {
			x := make([]float64, numBtns)
			for i, fv := range freeVars {
				x[fv] = float64(freeVals[i])
			}

			sum := currentSum
			valid := true
			for col, row := range pivotRows {
				if row == -1 {
					continue
				}
				val := matrix[row][numBtns]
				for c := range numBtns {
					if c != col {
						val -= matrix[row][c] * x[c]
					}
				}
				rv := int(math.Round(val))
				if rv < 0 || math.Abs(val-float64(rv)) > 1e-6 {
					valid = false
					break
				}
				sum += rv
			}

			if valid && sum < bestSum {
				bestSum = sum
			}
			return
		}

		for v := 0; v <= maxFreeVal; v++ {
			search(idx+1, append(freeVals, v), currentSum+v)
		}
	}

	search(0, nil, 0)

	if bestSum == math.MaxInt {
		return -1
	}
	return bestSum
}

func part1(machines []Machine) int {
	total := 0
	for _, m := range machines {
		total += minPresses(m)
	}
	return total
}

func part2(machines []Machine) int {
	total := 0
	for _, m := range machines {
		total += minPressesJoltage(m)
	}
	return total
}

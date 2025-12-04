package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

var directions = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Reading input file went horribly wrong")
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	grid := make([][]byte, len(lines))
	for i := range lines {
		grid[i] = []byte(lines[i])
	}

	startTime := time.Now()

	part1sum := part1(grid)
	fmt.Println("Part 1 answer: ", part1sum)

	part2sum := part2(grid)
	fmt.Println("Part 2 answer: ", part2sum)

	fmt.Printf("Time: %.2fms\n", float64(time.Since(startTime).Microseconds())/1000)
}

const paper = byte('@')
const removed = byte('x')

func part1(grid [][]byte) int {
	sum, _ := scanGrid(false, grid)

	return sum
}

func part2(grid [][]byte) int {

	sum := 0

	for {
		var partSum int
		partSum, grid = scanGrid(true, grid)
		if partSum == 0 {
			break
		}
		sum += partSum
	}

	return sum
}

func scanGrid(part2 bool, grid [][]byte) (int, [][]byte) {
	sum := 0
	nextGrid := make([][]byte, len(grid))
	for i := range grid {
		nextGrid[i] = bytes.Clone(grid[i])
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != paper {
				continue
			}

			neighbors := 0
			for _, dir := range directions {
				x := i + dir[0]
				y := j + dir[1]
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) {
					continue
				}
				adjacent := grid[x][y]
				if adjacent == paper {
					neighbors++
				}
			}

			if neighbors < 4 {
				if part2 {
					nextGrid[i][j] = removed
				}
				sum++
			}

		}
	}
	return sum, nextGrid
}

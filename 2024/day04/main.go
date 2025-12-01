package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// input := readFileContents("test.txt")
	input := readFileContents("input.txt")
	result := part2(input)

	fmt.Println("Result: ", result)
}

var directions = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

var directions2 = [4][2]int{
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

func part1(input []string) int {
	sum := 0

	for i, line := range input {
		for j := range len(line) {
			if input[i][j] == byte('X') {
				for _, dir := range directions {
					if getWord(input, i, j, dir) == "XMAS" {
						sum += 1
					}
				}
			}

		}
	}

	return sum
}

func getWord(grid []string, row int, col int, dir [2]int) string {
	word := []byte{grid[row][col]}
	for i := 0; i < 3; i++ {
		row += dir[0]
		col += dir[1]
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			return ""
		}
		word = append(word, grid[row][col])
	}
	return string(word)
}

func getWord2(grid []string, row int, col int, dir [2]int) string {
	word := []byte{grid[row][col]}
	for i := 0; i < 3; i++ {
		row += dir[0]
		col += dir[1]
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			return ""
		}
		if grid[row][col] == 'S' {

		}
		word = append(word, grid[row][col])
	}
	return string(word)
}

func part2(input []string) int {
	sum := 0

	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[0])-1; j++ {
			if input[i][j] == byte('A') {
				tl := input[i-1][j-1]
				br := input[i+1][j+1]
				bl := input[i-1][j+1]
				tr := input[i+1][j-1]

				tlbr := string(tl) + string(br)
				bltr := string(bl) + string(tr)

				if (tlbr == "MS" || tlbr == "SM") && (bltr == "MS" || bltr == "SM") {
					sum += 1
				}
			}
		}
	}

	return sum
}

func readFileContents(filePath string) []string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	dataString := string(data)
	grid := strings.Split(dataString, "\n")

	return grid
}

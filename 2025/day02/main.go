package main

import (
	"fmt"
	"os"

	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	dataString := string(data)
	lines := strings.Split(dataString, ",")

	var part1sum = part1(lines)
	fmt.Println("Part 1 answer: ", part1sum)

	var part2sum = part2(lines)
	fmt.Println("Part 2 answer: ", part2sum)
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		stringRange := strings.Split(line, "-")
		start, _ := strconv.Atoi(stringRange[0])
		end, _ := strconv.Atoi(stringRange[1])

		for i := start; i < end+1; i++ {
			stringValue := strconv.Itoa(i)
			length := len(stringValue)
			if length%2 != 0 {
				continue
			}

			firstHalf := stringValue[:length/2]
			secondHalf := stringValue[length/2:]
			if firstHalf == secondHalf {
				sum += i
			}
		}

	}

	return sum
}

func repeatingPattern(s string, size int) bool {
	if len(s)%size != 0 {
		return false
	}

	for i := size; i < len(s); i++ {
		if s[i] != s[i%size] {
			return false
		}
	}
	return true
}

func part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			half := len(s) / 2

			for size := 1; size <= half; size++ {
				if repeatingPattern(s, size) {
					sum += i
					break
				}
			}
		}
	}
	return sum
}

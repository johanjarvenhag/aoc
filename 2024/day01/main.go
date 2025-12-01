package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	dataString := string(data)
	lines := strings.Split(dataString, "\n")

	var part1sum = part1(lines)
	fmt.Println("Part 1 answer: ", part1sum)

	var part2sum = part2(lines)
	fmt.Println("Part 2 answer: ", part2sum)
}

func part1(lines []string) int {
	var firstList []int
	var secondList []int

	// Fill the lists
	for i := range lines {
		line := strings.Split(lines[i], "   ")
		lineInt, _ := strconv.Atoi(line[0])
		firstList = append(firstList, lineInt)
		lineSecondInt, _ := strconv.Atoi(line[1])
		secondList = append(secondList, lineSecondInt)
	}

	// Sort the lists
	slices.Sort(firstList)
	slices.Sort(secondList)

	// Calculate the sum of the difference in the lists
	sum := 0

	for i := 0; i < len(firstList); i++ {
		diff := firstList[i] - secondList[i]
		actualDiff := max(diff, -diff)
		sum = sum + actualDiff
	}

	return sum
}

func part2(lines []string) int {
	var firstList []int
	var secondList []int

	// Fill the lists
	for i := range lines {
		line := strings.Split(lines[i], "   ")
		lineInt, _ := strconv.Atoi(line[0])
		firstList = append(firstList, lineInt)
		lineSecondInt, _ := strconv.Atoi(line[1])
		secondList = append(secondList, lineSecondInt)
	}

	sum := 0

	for i := range firstList {
		count := 0
		for _, v := range secondList {
			if v == firstList[i] {
				count++
			}
		}
		sum += firstList[i] * count
	}

	return sum
}

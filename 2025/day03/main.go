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
	// lines := strings.Split(dataString, "\n")

	var part1sum = part1(dataString)
	fmt.Println("Part 1 answer: ", part1sum)

	var part2sum = part2(dataString)
	fmt.Println("Part 2 answer: ", part2sum)
}

func part1(input string) int {
	sum := 0
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		length := len(line)
		ogArr := make([]byte, length)

		for i, r := range line {
			ogArr[i] = byte(r)
		}

		sortedOgArr := make([]byte, length-1)
		copy(sortedOgArr, ogArr[:len(ogArr)-1])
		slices.Sort(sortedOgArr)
		firstMaxValue := sortedOgArr[len(sortedOgArr)-1]

		maxValueIndex := slices.Index(ogArr, firstMaxValue)
		sortedSecondArr := make([]byte, len(ogArr[maxValueIndex:]))
		copy(sortedSecondArr, ogArr[maxValueIndex+1:])
		slices.Sort(sortedSecondArr)
		secondMaxValue := sortedSecondArr[len(sortedSecondArr)-1]

		maxJoltage := string(firstMaxValue) + string(secondMaxValue)
		a, _ := strconv.Atoi(maxJoltage)
		sum += a
	}

	return sum
}

func part2(input string) int {
	sum := 0
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		length := len(line)
		inputBytes := make([]byte, length)

		for i, r := range line {
			inputBytes[i] = byte(r)
		}

		stringSum := ""
		for i := 12; i >= 1; i-- {
			inputCopy := inputBytes[:len(inputBytes)-i+1]
			maxValue := maxValue(inputCopy)
			stringSum += string(maxValue)
			maxValueIndex := slices.Index(inputBytes, maxValue)
			inputBytes = inputBytes[maxValueIndex+1:]
		}

		a, _ := strconv.Atoi(stringSum)
		sum += a
	}

	return sum
}

func maxValue(batteries []byte) byte {
	sortedOgArr := make([]byte, len(batteries))
	copy(sortedOgArr, batteries)
	slices.Sort(sortedOgArr)
	firstMaxValue := sortedOgArr[len(batteries)-1]
	return firstMaxValue
}

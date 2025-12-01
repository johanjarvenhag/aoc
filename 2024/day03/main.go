package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// input := readFileContents("test.txt")
	input := readFileContents("input.txt")
	result := part2(input)

	fmt.Println("Result: ", result)
}

func part1(input string) int {
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	matches := re.FindAllStringSubmatch(input, -1)
	sum := 0

	for _, m := range matches {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])

		result := a * b
		sum += result
		fmt.Println(result)
	}

	return sum
}

func part2(input string) int {
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don\'t\(\)`)

	matches := re.FindAllStringSubmatch(input, -1)
	sum := 0

	do := true

	for _, m := range matches {
		if m[0] == "do()" {
			do = true
			continue
		}

		if m[0] == "don't()" {
			do = false
		}

		if do {
			a, _ := strconv.Atoi(m[1])
			b, _ := strconv.Atoi(m[2])

			result := a * b
			sum += result
		}
	}

	return sum
}

func readFileContents(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	return string(data)
}

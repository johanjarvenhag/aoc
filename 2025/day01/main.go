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
	lines := strings.Split(dataString, "\n")

	var part1sum = part1(lines)
	fmt.Println("Part 1 answer: ", part1sum)

	var part2sum = part2(lines)
	fmt.Println("Part 2 answer: ", part2sum)
}

// 1129
func part1(lines []string) int {
	sum := 0
	dialPosition := 50

	for _, line := range lines {
		dialPosition, _ = turnDial(dialPosition, line)

		if dialPosition == 0 {
			sum++
		}
	}

	return sum
}

// 6638
func part2(lines []string) int {
	sum := 0
	dialPosition := 50

	for _, line := range lines {
		zeroHits := 0
		dialPosition, zeroHits = turnDial(dialPosition, line)
		sum += zeroHits
	}

	return sum
}

func turnDial(pos int, cmd string) (int, int) {
	if len(cmd) < 2 {
		return pos, 0
	}

	dir := cmd[0]
	steps, err := strconv.Atoi(cmd[1:])
	if err != nil {
		return pos, 0
	}

	zeroHits := steps / 100   // full rotations
	diff := steps % 100       // remainder movement
	startedOnZero := pos == 0 // special rule

	incZeroHit := func() {
		if !startedOnZero {
			zeroHits++
		}
	}

	switch dir {
	case 'L':
		switch {
		case diff > pos:
			// crosses zero
			incZeroHit()
			return 100 + pos - diff, zeroHits

		case diff == pos:
			// lands on zero
			incZeroHit()
			return 0, zeroHits

		default:
			// no crossing
			return pos - diff, zeroHits
		}

	case 'R':
		if pos+diff >= 100 {
			// crosses zero
			incZeroHit()
			return pos + diff - 100, zeroHits
		}
		return pos + diff, zeroHits

	default:
		return pos, zeroHits
	}
}

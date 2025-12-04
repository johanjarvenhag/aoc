package main

import (
	"os"
	"strings"
	"testing"
)

var input = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestPart2(t *testing.T) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	grid := make([][]byte, len(lines))
	for i := range lines {
		grid[i] = []byte(lines[i])
	}
	got := part2(grid)
	want := 43
	if got != want {
		t.Errorf("Got = %v; want %v", got, want)
	}
}

func Benchmark(b *testing.B) {
	for b.Loop() {
		main()
	}
}

func BenchmarkPart2(b *testing.B) {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	grid := make([][]byte, len(lines))
	for i := range lines {
		grid[i] = []byte(lines[i])
	}
	for b.Loop() {
		part2(grid)
	}
}

// func TestPart1(t *testing.T) {
// 	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
// 	got := part1(input)
// 	want := 12277755546

// 	if got != want {
// 		t.Errorf("Got = %v; want %v", got, want)
// 	}
// }

// func TestTurnDial(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		start    int
// 		move     int
// 		expected int
// 	}{
// 		// No movement cases
// 		{"example", 50, 1000, 10},
// 		{"example", 50, -1001, 10},
// 		{"iavr", 50, 150, 2},
// 		{"iavr", 99, 1, 1},
// 		{"ivar 2", 0, 100, 1},
// 		{"ivar 2", 52, 48, 1},
// 		{"ivar 2", 0, -5, 0},
// 		{"no movement", 50, 0, 0},
// 		{"no movement at zero", 0, 0, 0},

// 		// Positive movement cases (no zero crossing)
// 		{"positive move within range", 10, 30, 0},
// 		{"positive move to exactly 99", 50, 49, 0},
// 		{"positive move starting at 0", 0, 50, 0},

// 		// Positive movement cases (crossing zero once)
// 		{"positive move crossing zero once", 50, 60, 1},
// 		{"positive move from 99 crossing zero", 99, 2, 1},
// 		{"positive move from 90 crossing zero", 90, 20, 1},

// 		// Positive movement cases (crossing zero multiple times)
// 		{"positive move crossing zero twice", 50, 160, 2},
// 		{"positive move crossing zero three times", 10, 291, 3},

// 		// Negative movement cases (no zero crossing)
// 		{"negative move within range", 50, -30, 0},
// 		{"negative move to exactly 1", 50, -49, 0},
// 		{"negative move starting at 99", 99, -50, 0},

// 		// Negative movement cases (crossing zero once)
// 		{"negative move crossing zero once", 30, -50, 1},
// 		{"negative move from 10 crossing zero", 10, -20, 1},

// 		// Negative movement cases (crossing zero multiple times)
// 		{"negative move crossing zero twice", 30, -150, 2},
// 		{"negative move crossing zero three times", 50, -290, 3},

// 		// Edge cases
// 		{"edge case - move exactly to next cycle", 0, 100, 1},
// 		{"edge case - large positive move", 0, 1000, 10},
// 		{"edge case - large negative move", 0, -1000, 10},
// 		{"edge case - move exactly to previous cycle", 0, -100, 1},
// 		{"edge case - negative move from 0 leaving zero", 0, -1, 0},
// 		{"edge case - positive move from 99 to 0, not crossing 0", 99, 1, 1},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			_, got := turnDial(tt.start, tt.move)
// 			if got != tt.expected {
// 				t.Errorf("turnDial(%d, %d) = %d; want %d", tt.start, tt.move, got, tt.expected)
// 			}
// 		})
// 	}
// }

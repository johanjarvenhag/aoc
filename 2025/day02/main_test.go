package main

import (
	"strings"
	"testing"
)

var lines = []string{
	"78847-119454",
	"636-933",
	"7143759788-7143793713",
	"9960235-10043487",
	"44480-68595",
	"23468-43311",
	"89-123",
	"785189-1014654",
	"3829443354-3829647366",
	"647009-692765",
	"2-20",
	"30-42",
	"120909-197026",
	"5477469-5677783",
	"9191900808-9191943802",
	"1045643-1169377",
	"46347154-46441299",
	"2349460-2379599",
	"719196-779497",
	"483556-641804",
	"265244-450847",
	"210541-230207",
	"195-275",
	"75702340-75883143",
	"58-84",
	"2152-3237",
	"3367-5895",
	"1552-2029",
	"9575-13844",
	"6048-8966",
	"419388311-419470147",
	"936-1409",
	"9292901468-9292987321",
}

func TestMain(t *testing.T) {

}

func TestRepeater(t *testing.T) {
	got := repeatingPattern("1698522", 1)
	want := false
	if got != want {
		t.Errorf("Got = %v; want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	lines := strings.Split(input, ",")
	got := part2(lines)
	want := 4174379265
	if got != want {
		t.Errorf("Got = %v; want %v", got, want)
	}
}

func BenchmarkPart2(b *testing.B) {
	for b.Loop() {
		part2(lines)
	}
}

func BenchmarkPart2b(b *testing.B) {
	for b.Loop() {
		part2b(lines)
	}
}

func BenchmarkPart2c(b *testing.B) {
	for b.Loop() {
		part2c(lines)
	}
}

func BenchmarkPart2Original(b *testing.B) {
	for n := 0; n < b.N; n++ {
		part2(lines)
	}
}

func BenchmarkPart2Optimized(b *testing.B) {
	for b.Loop() {
		part2Optimized(lines)
	}
}

func BenchmarkPart2AI(b *testing.B) {
	for b.Loop() {
		part2Fast(lines)
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

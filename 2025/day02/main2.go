package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Range type
type rng struct {
	lo uint64
	hi uint64
}

// parseRanges parses strings like "2-20" into []rng
func parseRanges(lines []string) ([]rng, error) {
	out := make([]rng, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("bad range: %q", line)
		}
		lo, err := strconv.ParseUint(parts[0], 10, 64)
		if err != nil {
			return nil, err
		}
		hi, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			return nil, err
		}
		if lo > hi {
			lo, hi = hi, lo
		}
		out = append(out, rng{lo: lo, hi: hi})
	}
	return out, nil
}

// pow10 returns 10^n
func pow10(n int) uint64 {
	return uint64(math.Pow10(n))
}

// repeatBase constructs the number by repeating `base` k times
func repeatBase(base uint64, k int, d int) uint64 {
	num := uint64(0)
	mult := pow10(d)
	for i := 0; i < k; i++ {
		num = num*mult + base
	}
	return num
}

// part2Fast generates only repeating numbers within each range
func part2Fast(lines []string) (uint64, error) {
	ranges, err := parseRanges(lines)
	if err != nil {
		return 0, err
	}

	var sum uint64 = 0

	for _, rg := range ranges {
		// number of digits in lo and hi
		minDigits := 1
		if rg.lo >= 10 {
			minDigits = int(math.Floor(math.Log10(float64(rg.lo)))) + 1
		}
		maxDigits := 1
		if rg.hi >= 10 {
			maxDigits = int(math.Floor(math.Log10(float64(rg.hi)))) + 1
		}

		// try each total length L in [minDigits, maxDigits]
		for L := minDigits; L <= maxDigits; L++ {
			// possible pattern lengths d (divides L)
			for d := 1; d*2 <= L; d++ {
				if L%d != 0 {
					continue
				}
				k := L / d
				lowBase := uint64(1)
				if d > 1 {
					lowBase = pow10(d - 1)
				}
				highBase := pow10(d) - 1

				for base := lowBase; base <= highBase; base++ {
					num := repeatBase(base, k, d)
					if num < rg.lo {
						continue
					}
					if num > rg.hi {
						break
					}
					sum += num
				}
			}
		}
	}

	return sum, nil
}

// Example usage
func main() {
	lines := []string{
		"78847-119454", "636-933", "7143759788-7143793713", "9960235-10043487",
		"44480-68595", "23468-43311", "89-123", "785189-1014654",
		"3829443354-3829647366", "647009-692765", "2-20", "30-42",
		"120909-197026", "5477469-5677783", "9191900808-9191943802",
		"1045643-1169377", "46347154-46441299", "2349460-2379599",
		"719196-779497", "483556-641804", "265244-450847", "210541-230207",
		"195-275", "75702340-75883143", "58-84", "2152-3237", "3367-5895",
		"1552-2029", "9575-13844", "6048-8966", "419388311-419470147",
		"936-1409", "9292901468-9292987321",
	}

	sum, err := part2Fast(lines)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("sum:", sum)
}

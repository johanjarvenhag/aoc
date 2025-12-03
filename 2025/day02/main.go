package main

import (
	"runtime"

	"strconv"
	"strings"
)

// func main() {
// 	data, _ := os.ReadFile("input.txt")
// 	dataString := string(data)
// 	lines := strings.Split(dataString, ",")

// 	// var part1sum = part1(lines)
// 	// fmt.Println("Part 1 answer: ", part1sum)

// 	var part2sum = part2(lines)
// 	fmt.Println("Part 2 answer: ", part2sum)

// 	var part2bsum = part2b(lines)
// 	fmt.Println("Part 2b answer: ", part2bsum)

// 	var part2csum = part2c(lines)
// 	fmt.Println("Part 2c answer: ", part2csum)

// 	var part2OptimizedSum = part2Optimized(lines)
// 	fmt.Println("Part 2 Optimized: ", part2OptimizedSum)
// }

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

func isRepeating(s string) bool {
	for size := 1; size <= len(s)/2; size++ {
		if len(s)%size == 0 && strings.Repeat(s[:size], len(s)/size) == s {
			return true
		}
	}
	return false
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

func part2b(lines []string) int {
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			half := len(s) / 2

			for size := 1; size <= half; size++ {
				if isRepeating(s) {
					sum += i
					break
				}
			}
		}
	}
	return sum
}

func part2c(lines []string) int {
	jobs := make(chan int, 1000)
	results := make(chan int, 1000)

	workerCount := runtime.NumCPU()

	// Start workers
	for w := 0; w < workerCount; w++ {
		go func() {
			for n := range jobs {
				s := strconv.Itoa(n)
				half := len(s) / 2
				added := false

				for size := 1; size <= half; size++ {
					if repeatingPattern(s, size) {
						results <- n
						added = true
						break
					}
				}

				if !added {
					results <- 0
				}
			}
		}()
	}

	// Publisher goroutine
	go func() {
		for _, line := range lines {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			for i := start; i <= end; i++ {
				jobs <- i
			}
		}
		close(jobs)
	}()

	// Collector
	total := 0
	totalJobs := 0
	for _, line := range lines {
		p := strings.Split(line, "-")
		start, _ := strconv.Atoi(p[0])
		end, _ := strconv.Atoi(p[1])
		totalJobs += (end - start + 1)
	}

	for i := 0; i < totalJobs; i++ {
		total += <-results
	}

	return total
}

func intToDigits(n int, buf []byte) []byte {
	// Fill buffer backwards
	i := len(buf)
	for n >= 10 {
		i--
		q := n / 10
		buf[i] = byte(n - q*10)
		n = q
	}
	i--
	buf[i] = byte(n)
	return buf[i:]
}

// Check if digits are repeating with a given period
func repeatingDigits(d []byte, size int) bool {
	// length divisible?
	if len(d)%size != 0 {
		return false
	}
	// Compare each digit with repeating base pattern
	for i := size; i < len(d); i++ {
		if d[i] != d[i%size] {
			return false
		}
	}
	return true
}

// Full optimized test
func isRepeatingNumber(n int, buf []byte) bool {
	digits := intToDigits(n, buf)
	l := len(digits)

	// Only try divisor sizes
	for size := 1; size*2 <= l; size++ {
		if l%size == 0 && repeatingDigits(digits, size) {
			return true
		}
	}
	return false
}

func part2Optimized(lines []string) int {
	sum := 0
	var buf [32]byte // supports up to 32 digits (far above what you need)

	for _, line := range lines {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for n := start; n <= end; n++ {
			if isRepeatingNumber(n, buf[:]) {
				sum += n
			}
		}
	}
	return sum
}

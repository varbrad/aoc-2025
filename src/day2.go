package main

import (
	"aoc-2025/src/utils"
	"strconv"
	"strings"
)

var D2 = Day{
	P1:    D2P1,
	P2:    D2P2,
	Input: D2Input,
}

func D2P1(input string) int {
	return d2_solve(d2_parse(input), true)
}

func D2P2(input string) int {
	return d2_solve(d2_parse(input), false)
}

func d2_solve(ranges [][2]int, fixedRepeat bool) int {
	set := make(map[int]bool)

	// Reusable buffer for int→string conversion
	buf := make([]byte, 0, 16)

	for _, r := range ranges {
		start := r[0]
		end := r[1]

		for n := start; n <= end; n++ {
			// Reset buffer
			buf = buf[:0]
			// Convert number to slice of bytes
			s := strconv.AppendInt(buf, int64(n), 10)
			length := len(s)

			if fixedRepeat {
				// Length must be even
				if length&1 != 0 {
					continue
				}

				// Get midpoint of the byte slice
				mid := length / 2

				// Check if both halves are equal
				if bytesEqual(s[:mid], s[mid:]) {
					set[n] = true
				}
				continue
			}

			// Try to repeat lengths that divide the string length
			for rep := 1; rep <= length/2; rep++ {
				// Length must be divisible by repeat length
				if length%rep != 0 {
					continue
				}

				// Compare each block with the first block
				first := s[:rep]
				ok := true

				for offset := rep; offset < length; offset += rep {
					if !bytesEqual(first, s[offset:offset+rep]) {
						ok = false
						break
					}
				}

				if ok {
					set[n] = true
					break // no need to check larger rep lengths, we found a match already
				}
			}
		}
	}

	sum := 0
	for k := range set {
		sum += k
	}
	return sum
}

// byte slice equality fn
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func d2_parse(input string) [][2]int {
	ranges := [][2]int{}

	for s := range strings.SplitSeq(strings.TrimSpace(input), ",") {
		parts := strings.Split(s, "-")
		rng := [2]int{
			utils.ToIntMust(parts[0]),
			utils.ToIntMust(parts[1]),
		}
		ranges = append(ranges, rng)
	}
	return ranges
}

package main

import (
	"aoc-2025/src/utils"
	"fmt"
	"strings"
)

func D2() {
	utils.RunDays(D2P1, D2P2, D2Input)
}

func D2P1(input string) int {
	return solve(parse(input), true)
}

func D2P2(input string) int {
	return solve(parse(input), false)
}

func parse(input string) [][2]int {
	ranges := [][2]int{}

	split := strings.Split(strings.TrimSpace(input), ",")

	for _, s := range split {
		parts := strings.Split(s, "-")
		rng := [2]int{
			utils.ToIntMust(parts[0]),
			utils.ToIntMust(parts[1]),
		}
		ranges = append(ranges, rng)
	}

	return ranges
}

func solve(ranges [][2]int, fixedRepeat bool) int {
	set := make(map[int]bool)

	for _, r := range ranges {
		start := r[0]
		end := r[1]

		for i := start; i <= end; i++ {
			// Turn it into a string
			s := fmt.Sprintf("%d", i)
			length := len(s)

			if fixedRepeat {
				if length%2 != 0 {
					continue
				}
				firstHalf := s[:length/2]
				secondHalf := s[length/2:]

				if firstHalf == secondHalf {
					set[i] = true
				}
				continue
			}

			for j := 1; j <= length/2; j++ {
				substr := s[:j]
				repeated := strings.Repeat(substr, length/j)

				if repeated == s {
					set[i] = true
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

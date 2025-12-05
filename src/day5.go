package main

import (
	"aoc-2025/src/utils"
	"slices"
	"strings"
)

func D5() {
	utils.RunDays(D5P1, D5P2, D5Input)
}

func D5P1(input string) int {
	state := d5_parse(input)

	count := 0

	for _, value := range state.values {
		if value < state.min || value > state.max {
			continue
		}

		for _, rng := range state.ranges {
			if value >= rng[0] && value <= rng[1] {
				count++
				break
			}
		}
	}

	return count
}

func D5P2(input string) int {
	ranges := d5_parse(input).ranges
	ranges = d5_merge_ranges(ranges)

	count := 0
	for _, r := range ranges {
		count += r[1] - r[0] + 1
	}

	return count
}

func d5_merge_ranges(ranges [][2]int) [][2]int {
	if len(ranges) == 0 {
		return ranges
	}

	// Sort ranges by start value
	slices.SortFunc(ranges, func(a, b [2]int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] > b[0] {
			return 1
		}
		return 0
	})

	// Merge overlapping ranges
	ix := len(ranges) - 1
	for {
		if ix == 0 {
			break
		}
		prev := ranges[ix-1]
		curr := ranges[ix]

		// if the end of range prev is greater than or equal to the start of range curr, we can combine!
		if prev[1] >= curr[0] {
			// merge the ranges
			ranges[ix-1][1] = utils.MaxInt(prev[1], curr[1])
			// remove the current range
			ranges = append(ranges[:ix], ranges[ix+1:]...)
		}

		ix--
	}

	// Are any ranges fully contained in another
	for j := len(ranges) - 1; j > 0; j-- {
		curr := ranges[j]
		prev := ranges[j-1]

		if (curr[0] >= prev[0] && curr[1] <= prev[1]) || (prev[0] >= curr[0] && prev[1] <= curr[1]) {
			// remove the current range
			ranges = append(ranges[:j], ranges[j+1:]...)
		}
	}

	return ranges
}

type d5_state struct {
	ranges [][2]int
	values []int
	min    int
	max    int
}

func d5_parse(input string) *d5_state {
	ab := strings.Split(strings.TrimSpace(input), "\n\n")

	a := ab[0]
	b := ab[1]

	ranges := [][2]int{}
	values := []int{}

	min := -1
	max := -1

	for _, line := range strings.Split(a, "\n") {
		parts := strings.Split(line, "-")
		start := utils.ToIntMust(parts[0])
		end := utils.ToIntMust(parts[1])

		rng := [2]int{
			start,
			end,
		}

		if min == -1 || start < min {
			min = start
		}
		if max == -1 || end > max {
			max = end
		}

		ranges = append(ranges, rng)
	}

	for _, line := range strings.Split(b, "\n") {
		values = append(values, utils.ToIntMust(line))
	}

	return &d5_state{
		ranges: ranges,
		values: values,
		min:    min,
		max:    max,
	}
}

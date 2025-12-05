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

	return utils.Reduce(state.values, func(count int, value int) int {
		for _, rng := range state.ranges {
			if value >= rng[0] && value <= rng[1] {
				return count + 1
			}
		}
		return count
	}, 0)
}

func D5P2(input string) int {
	ranges := d5_parse(input).ranges

	return utils.Reduce(ranges, func(count int, rng [2]int) int {
		return count + (rng[1] - rng[0] + 1)
	}, 0)
}

type d5_state struct {
	ranges [][2]int
	values []int
}

func d5_parse(input string) d5_state {
	a, b, _ := strings.Cut(strings.TrimSpace(input), "\n\n")

	ranges := [][2]int{}
	values := []int{}

	for line := range strings.SplitSeq(a, "\n") {
		parts := strings.Split(line, "-")
		start := utils.ToIntMust(parts[0])
		end := utils.ToIntMust(parts[1])

		ranges = append(ranges, [2]int{start, end})
	}

	for line := range strings.SplitSeq(b, "\n") {
		values = append(values, utils.ToIntMust(line))
	}

	return d5_state{
		ranges: d5_merge_ranges(ranges), // pre-merge ranges for efficiency
		values: values,
	}
}

func d5_merge_ranges(ranges [][2]int) [][2]int {
	// Sort ranges by start value (low -> high)
	slices.SortFunc(ranges, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	merged := make([][2]int, 0, len(ranges))
	current := ranges[0] // current range

	for i := 1; i < len(ranges); i++ {
		r := ranges[i] // next range

		// if the next range start is smaller or equal to the current range end
		if r[0] <= current[1] {
			// if the next range end is larger than the current range end
			if r[1] > current[1] {
				// extend the current range
				current[1] = r[1]
			}
			// skip onto the next range, we've merged it
			continue
		}

		// add the currently modified range to the list and move to the next
		merged = append(merged, current)
		// set current to next range
		current = r
	}

	// add the last range to the list
	merged = append(merged, current)

	return merged
}

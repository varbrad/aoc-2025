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

func d5_parse(input string) *d5_state {
	ab := strings.Split(strings.TrimSpace(input), "\n\n")

	a, b := ab[0], ab[1]

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

	return &d5_state{
		ranges: d5_merge_ranges(ranges), // pre-merge ranges for efficiency
		values: values,
	}
}

func d5_merge_ranges(ranges [][2]int) [][2]int {
	// Sort ranges by start value (low -> high)
	slices.SortFunc(ranges, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	// Merge overlapping ranges
	for ix := len(ranges) - 1; ix > 0; ix-- {
		prev, curr := ranges[ix-1], ranges[ix]

		// if the end of range prev is greater than or equal to the start of range curr, we can combine!
		if prev[1] >= curr[0] {
			// merge the ranges by setting the end of the previous range to the max of both ends
			ranges[ix-1][1] = utils.MaxInt(prev[1], curr[1])
			// remove the current range as it's now merged
			ranges = slices.Delete(ranges, ix, ix+1)
		}
	}

	// Now consider, are there any ranges fully contained within another range?
	for j := len(ranges) - 1; j > 0; j-- {
		curr, prev := ranges[j], ranges[j-1]

		// If the current range is fully contained in the previous range
		if curr[0] >= prev[0] && curr[1] <= prev[1] {
			// remove the current range
			ranges = slices.Delete(ranges, j, j+1)
		}
	}

	return ranges
}

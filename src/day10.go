package main

import (
	"aoc-2025/src/utils"
	"slices"
	"strings"
)

var D10 = Day{
	P1:    D10P1,
	P2:    D10P2,
	Input: D10Input,
}

func D10P1(input string) int {
	states := d10_parse(input)

	sum := 0
	for _, state := range states {
		steps := d10_solve_machine(state)
		sum += steps
	}

	return sum
}

func d10_solve_machine(state d10_state) int {
	permutations := d10_permute(len(state.buttons))

	for _, perm := range permutations {
		lights := make([]bool, len(state.lights))
		copy(lights, state.lights)

		for b, press := range perm {
			if press {
				for _, lightIdx := range state.buttons[b] {
					lights[lightIdx] = !lights[lightIdx]
				}
			}
		}

		allOff := true
		for _, light := range lights {
			if light {
				allOff = false
				break
			}
		}

		if allOff {
			steps := 0
			for _, press := range perm {
				if press {
					steps++
				}
			}
			return steps
		}
	}

	return -1
}

var _cache map[int][][]bool = make(map[int][][]bool)

func d10_permute(n int) [][]bool {
	if val, ok := _cache[n]; ok {
		return val
	}

	total := 1 << n
	out := make([][]bool, total)

	for i := 0; i < total; i++ {
		row := make([]bool, n)
		for b := 0; b < n; b++ {
			shift := n - 1 - b
			row[b] = ((i >> shift) & 1) == 1
		}
		out[i] = row
	}

	slices.SortFunc(out, func(a, b []bool) int {
		trueCountA := 0
		for _, v := range a {
			if v {
				trueCountA++
			}
		}
		trueCountB := 0
		for _, v := range b {
			if v {
				trueCountB++
			}
		}
		return trueCountA - trueCountB
	})

	_cache[n] = out

	return out[1:]
}

func D10P2(input string) int {
	return -1
}

type d10_state struct {
	lights  []bool
	buttons [][]int
}

func d10_parse(input string) []d10_state {
	lines := utils.SplitLines(input)

	machines := make([]d10_state, len(lines))

	for i, line := range lines {
		machines[i] = d10_state{
			lights:  []bool{},
			buttons: [][]int{},
		}

		for part := range strings.SplitSeq(strings.TrimSpace(line), " ") {
			switch part[0] {
			case '[':
				lightStr := strings.Trim(part, "[]")
				for _, ch := range lightStr {
					if ch == '#' {
						machines[i].lights = append(machines[i].lights, true)
					} else {
						machines[i].lights = append(machines[i].lights, false)
					}
				}
			case '(':
				buttonStr := strings.Trim(part, "()")
				buttonParts := strings.Split(buttonStr, ",")
				buttons := make([]int, len(buttonParts))
				for j, b := range buttonParts {
					buttons[j] = utils.ToIntMust(b)
				}
				machines[i].buttons = append(machines[i].buttons, buttons)
			}
		}
	}

	return machines
}

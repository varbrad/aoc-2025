package main

import (
	"aoc-2025/src/utils"
	"fmt"
	"slices"
)

var D7 = Day{
	P1:    D7P1,
	P2:    D7P2,
	Input: D7Input,
}

func D7P1(input string) int {
	state := d7_parse(input)

	sim := d7_simulate(&state)

	return len(sim.visitedSplitters)
}

func D7P2(input string) int {
	state := d7_parse(input)
	d7_simulate(&state)

	// Store the valid splitters after the first simulation
	validSplitters := state.splitters
	cache := make(map[string]int)
	paths := d7_recurse(&validSplitters, validSplitters[0], cache)

	return paths
}

func d7_recurse(splitters *[]d7_pos, pos d7_pos, cache map[string]int) int {
	// Check the cache first
	key := fmt.Sprintf("%d,%d|", pos.x, pos.y)
	if val, exists := cache[key]; exists {
		return val
	}

	// We know this splitter will split the beam into two paths, left and right
	// Find what splitters are to the left and right
	leftSplitter := d7_pos{-1, -1}
	rightSplitter := d7_pos{-1, -1}
	for _, splitter := range *splitters {
		// If we have both splitters, break
		hasLeft := leftSplitter.x != -1
		hasRight := rightSplitter.x != -1
		if hasLeft && hasRight {
			break
		}

		// Firstly, this splitter must be below the current position
		if splitter.y <= pos.y {
			continue
		}

		// Check for left splitter
		if !hasLeft && splitter.x == pos.x-1 && splitter.y > pos.y {
			leftSplitter = splitter
			continue
		}

		// Check for right splitter
		if !hasRight && splitter.x == pos.x+1 && splitter.y > pos.y {
			rightSplitter = splitter
			continue
		}
	}

	// If we found no splitters, this is an end path where both beams just go off into infinity
	if leftSplitter.x == -1 && rightSplitter.x == -1 {
		return 2 // Both beams go to infinity
	}

	paths := 0

	// Recurse left
	if leftSplitter.x != -1 {
		paths += d7_recurse(splitters, leftSplitter, cache)
	} else {
		paths += 1 // Left beam goes to infinity
	}

	// Recurse right
	if rightSplitter.x != -1 {
		paths += d7_recurse(splitters, rightSplitter, cache)
	} else {
		paths += 1 // Right beam goes to infinity
	}

	// Store in cache
	cache[key] = paths

	return paths
}

type d7_pos struct {
	x int
	y int
}

type d7_state struct {
	start     d7_pos
	splitters []d7_pos
}

type d7_simulation struct {
	visitedSplitters map[string]bool
}

func d7_simulate(state *d7_state) d7_simulation {
	sim := d7_simulation{
		visitedSplitters: make(map[string]bool),
	}

	beams := []d7_pos{state.start}

	for len(beams) > 0 {
		// Take the first beam and find a splitter below it
		beam := beams[0]
		beams = beams[1:]

		for _, splitter := range state.splitters {
			// Check if the splitter is directly below the beam
			if splitter.x == beam.x && splitter.y > beam.y {
				key := fmt.Sprintf("%d,%d", splitter.x, splitter.y)
				if !sim.visitedSplitters[key] {
					sim.visitedSplitters[key] = true
					// Create new beams to the left and right
					beams = append(beams, d7_pos{x: splitter.x - 1, y: splitter.y})
					beams = append(beams, d7_pos{x: splitter.x + 1, y: splitter.y})
				}
				break
			}
		}
	}

	for i := len(state.splitters) - 1; i >= 0; i-- {
		key := fmt.Sprintf("%d,%d", state.splitters[i].x, state.splitters[i].y)
		if !sim.visitedSplitters[key] {
			fmt.Printf("Splitter at (%d, %d) was not visited\n", state.splitters[i].x, state.splitters[i].y)
			state.splitters = slices.Delete(state.splitters, i, i+1)
		}
	}

	return sim
}

func d7_parse(input string) d7_state {
	rows := utils.SplitLines(input)
	state := d7_state{
		start:     d7_pos{},
		splitters: []d7_pos{},
	}

	for y, row := range rows {
		for x, char := range row {
			switch char {
			case 'S':
				{
					state.start.x = x
					state.start.y = y
				}
			case '^':
				{
					state.splitters = append(state.splitters, d7_pos{x: x, y: y})
				}
			}
		}
	}

	return state
}

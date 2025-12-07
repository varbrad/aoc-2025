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

	tree := d7_build_tree(validSplitters)

	return d7_recurse(tree, make(map[string]int))
}

func d7_recurse(node *d7_node, cache map[string]int) int {
	// Check the cache first
	key := fmt.Sprintf("%d,%d|", node.x, node.y)
	if val, exists := cache[key]; exists {
		return val
	}

	// Base case: if leaf node
	if node.left == nil && node.right == nil {
		return 2
	}

	paths := 0

	if node.left != nil {
		paths += d7_recurse(node.left, cache)
	}

	if node.right != nil {
		paths += d7_recurse(node.right, cache)
	}

	// Store in cache
	cache[key] = paths

	return paths
}

type d7_node struct {
	x     int
	y     int
	left  *d7_node
	right *d7_node
}

func d7_build_tree(splitters []d7_pos) *d7_node {
	nodes := make([]*d7_node, len(splitters))
	for i := range splitters {
		nodes[i] = &d7_node{
			x: splitters[i].x,
			y: splitters[i].y,
		}
	}

	for _, n1 := range nodes {
		// Find left child
		leftChildIndex := -1
		for j, n2 := range nodes {
			if n2.x == n1.x-1 && n2.y > n1.y {
				leftChildIndex = j
				break
			}
		}
		if leftChildIndex != -1 {
			n1.left = nodes[leftChildIndex]
		}

		// Find right child
		rightChildIndex := -1
		for j, n2 := range nodes {
			if n2.x == n1.x+1 && n2.y > n1.y {
				rightChildIndex = j
				break
			}
		}
		if rightChildIndex != -1 {
			n1.right = nodes[rightChildIndex]
		}
	}

	return nodes[0]
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

package main

import (
	"aoc-2025/src/utils"
	"strings"
)

var D11 = Day{
	P1:    D11P1,
	P2:    D11P2,
	Input: D11Input,
}

func D11P1(input string) int {
	nodes := d11_parse(input)
	return d11_calculate_routes(nodes, "you", "out")
}

func D11P2(input string) int {
	nodes := d11_parse(input)

	// route A: svr -> fft -> dac -> out
	svr_to_fft := d11_calculate_routes(nodes, "svr", "fft", "dac")
	fft_to_dac := d11_calculate_routes(nodes, "fft", "dac")
	dac_to_out := d11_calculate_routes(nodes, "dac", "out")

	// route B: svr -> dac -> fft -> out
	svr_to_dac := d11_calculate_routes(nodes, "svr", "dac", "fft")
	dac_to_fft := d11_calculate_routes(nodes, "dac", "fft")
	fft_to_out := d11_calculate_routes(nodes, "fft", "out")

	waysA := svr_to_fft * fft_to_dac * dac_to_out
	waysB := svr_to_dac * dac_to_fft * fft_to_out

	return waysA + waysB
}

func d11_calculate_routes(nodes map[string][]string, from string, target string, ignoring ...string) int {
	waysMap := make(map[string]int)
	waysMap[target] = 1

	// mark ignoring nodes as 0 ways (dead end)
	for _, ign := range ignoring {
		waysMap[ign] = 0
	}

	stack := []string{from}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]

		// already in the map - if so pop the stack and continue
		if _, ok := waysMap[curr]; ok {
			stack = stack[:len(stack)-1]
			continue
		}

		children, ok := nodes[curr]
		if !ok {
			// leaf node, no ways to target from here
			waysMap[curr] = 0
			stack = stack[:len(stack)-1]
			continue
		}

		allChildrenInMap := true
		totalWays := 0
		// loop thru children, and see if they are all in the map and count ways
		for _, child := range children {
			if v, ok := waysMap[child]; !ok {
				allChildrenInMap = false
				stack = append(stack, child)
			} else {
				totalWays += v
			}
		}

		// if not all children are in the map, continue (we've pushed them to the stack already)
		if !allChildrenInMap {
			continue
		}

		// otherwise, all children are in the map, so set ways and pop
		waysMap[curr] = totalWays
		stack = stack[:len(stack)-1]
	}

	// finally return the ways from 'from' to 'target'
	return waysMap[from]
}

func d11_parse(input string) map[string][]string {
	rows := utils.SplitLines(input)
	nodes := make(map[string][]string)

	for _, row := range rows {
		spl := strings.Split(row, ":")

		node := spl[0]
		edges := strings.Fields(spl[1])

		nodes[node] = edges
	}

	return nodes
}

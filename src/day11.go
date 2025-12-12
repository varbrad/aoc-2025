package main

import (
	"aoc-2025/src/utils"
	"fmt"
	"strings"
)

var D11 = Day{
	P1:    D11P1,
	P2:    D11P2,
	Input: D11Input,
}

func D11P1(input string) int {
	rows := utils.SplitLines(input)

	nodes := make(map[string][]string)

	for _, row := range rows {
		spl := strings.Split(row, ":")

		node := spl[0]
		edges := strings.Fields(spl[1])

		nodes[node] = edges
	}

	return d11_find_paths(nodes, "you", "out")
}

func D11P2(input string) int {
	rows := utils.SplitLines(input)

	nodes := make(map[string][]string)

	for _, row := range rows {
		spl := strings.Split(row, ":")

		node := spl[0]
		edges := strings.Fields(spl[1])

		nodes[node] = edges
	}

	nodesThatCanReachFFT := d11_nodes_that_can_reach(nodes, "fft")
	nodesThatCanReachDAC := d11_nodes_that_can_reach(nodes, "dac")

	fmt.Println("Nodes that can reach FFT:", nodesThatCanReachFFT)
	fmt.Println("Nodes that can reach DAC:", nodesThatCanReachDAC)

	fft_to_dac := d11_find_paths(nodes, "fft", "dac")
	fmt.Println("Paths from FFT to DAC:", fft_to_dac)
	// return d11_find_paths_p2(nodes, nodesThatCanReachFFT, nodesThatCanReachDAC)

	return -1
}

type d11_p2 struct {
	path       []string
	visitedFFT bool
	visitedDAC bool
}

func d11_find_paths_p2(nodes map[string][]string, fft map[string]bool, dac map[string]bool) int {
	paths := []d11_p2{{path: []string{"svr"}, visitedFFT: false, visitedDAC: false}}
	totalPaths := 0

	for {
		if len(paths)%1000 == 0 {
			fmt.Println("Paths to explore:", len(paths), "Total paths found:", totalPaths)
		}
		if len(paths) == 0 {
			break
		}
		current := paths[0]
		currentNode := current.path[len(current.path)-1]
		visitedFFT := current.visitedFFT
		visitedDAC := current.visitedDAC
		paths = paths[1:]

		// Find all next steps
		nextNodes := nodes[currentNode]

		for _, nextNode := range nextNodes {
			// If we've visted both fft and dac and are at out, count a path
			if nextNode == "out" && visitedFFT && visitedDAC {
				fmt.Println("Found path:", append(current.path, nextNode), visitedFFT, visitedDAC)
				totalPaths++
				continue
			}

			// If we've not yet reached fft, check if we can reach it from here
			if !visitedFFT && !fft[nextNode] {
				continue
			}

			// If we've not yet reached dac, check if we can reach it from here
			if !visitedDAC && !dac[nextNode] {
				continue
			}

			newPath := make([]string, len(current.path))
			copy(newPath, current.path)
			newPath = append(newPath, nextNode)
			paths = append(paths, d11_p2{
				path:       newPath,
				visitedFFT: visitedFFT || nextNode == "fft",
				visitedDAC: visitedDAC || nextNode == "dac",
			})
		}
	}

	return totalPaths
}

func d11_find_paths(unoptimised_nodes map[string][]string, start string, end string) int {
	optimised_nodes := d11_nodes_that_can_reach(unoptimised_nodes, end)
	nodes := make(map[string][]string)
	for node := range optimised_nodes {
		nodes[node] = unoptimised_nodes[node]
	}

	paths := [][]string{{start}}
	totalPaths := 0

	for {
		if len(paths) == 0 {
			break
		}

		currentPath := paths[0]
		currentNode := currentPath[len(currentPath)-1]
		paths = paths[1:]

		// Find all next steps
		nextNodes := nodes[currentNode]

		for _, nextNode := range nextNodes {
			if nextNode == end {
				totalPaths++
				continue
			}

			newPath := make([]string, len(currentPath))
			copy(newPath, currentPath)
			newPath = append(newPath, nextNode)
			paths = append(paths, newPath)
		}
	}

	return totalPaths
}

func d11_nodes_that_can_reach(nodes map[string][]string, target string) map[string]bool {
	result := map[string]bool{}

	for node := range nodes {
		if d11_can_get_to(nodes, node, target) {
			result[node] = true
		}
	}

	return result
}

func d11_can_get_to(nodes map[string][]string, start string, end string) bool {
	visited := make(map[string]bool)
	toVisit := []string{start}

	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]

		if current == end {
			return true
		}

		if visited[current] {
			continue
		}

		visited[current] = true

		for _, neighbor := range nodes[current] {
			if !visited[neighbor] {
				toVisit = append(toVisit, neighbor)
			}
		}
	}

	return false
}

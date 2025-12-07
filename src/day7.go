package main

import (
	"aoc-2025/src/utils"
	"fmt"
)

var D7 = Day{
	P1:    D7P1,
	P2:    D7P2,
	Input: D7Input,
}

func D7P1(input string) int {
	return d7_tree_unique_nodes(d7_parse(input), make(map[string]bool))
}

func D7P2(input string) int {
	return d7_recurse(d7_parse(input), make(map[string]int))
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
	} else {
		paths += 1
	}

	if node.right != nil {
		paths += d7_recurse(node.right, cache)
	} else {
		paths += 1
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

func d7_tree_unique_nodes(node *d7_node, seen map[string]bool) int {
	if node == nil {
		return 0
	}

	key := fmt.Sprintf("%d,%d", node.x, node.y)
	if seen[key] {
		return 0
	}

	seen[key] = true

	count := 1

	count += d7_tree_unique_nodes(node.left, seen)
	count += d7_tree_unique_nodes(node.right, seen)

	return count
}

type d7_pos struct {
	x int
	y int
}

func d7_parse(input string) *d7_node {
	rows := utils.SplitLines(input)
	splitters := []d7_pos{}

	for y, row := range rows {
		for x, char := range row {
			switch char {
			case '^':
				{
					splitters = append(splitters, d7_pos{x: x, y: y})
				}
			}
		}
	}

	return d7_build_tree(splitters)
}

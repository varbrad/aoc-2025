package main

import (
	"aoc-2025/src/utils"
)

var D7 = Day{
	P1:    D7P1,
	P2:    D7P2,
	Input: D7Input,
}

func D7P1(input string) int {
	return d7_tree_unique_nodes(d7_parse(input), make(map[d7_pos]bool))
}

func D7P2(input string) int {
	return d7_recurse(d7_parse(input), make(map[d7_pos]int))
}

func d7_recurse(node *d7_node, cache map[d7_pos]int) int {
	// Check the cache first
	if val, exists := cache[*node.pos]; exists {
		return val
	}

	left := 1
	if node.left != nil {
		left = d7_recurse(node.left, cache)
	}

	right := 1
	if node.right != nil {
		right = d7_recurse(node.right, cache)
	}

	// Store in cache
	cache[*node.pos] = left + right
	return cache[*node.pos]
}

type d7_node struct {
	pos   *d7_pos
	left  *d7_node
	right *d7_node
}

func d7_build_tree(splitters []d7_pos) *d7_node {
	nodes := make([]*d7_node, len(splitters))
	for i := range splitters {
		nodes[i] = &d7_node{
			pos: &d7_pos{splitters[i].x, splitters[i].y},
		}
	}

	for _, n1 := range nodes {
		// Find left child
		leftChildIndex := -1
		for j, n2 := range nodes {
			if n2.pos.x == n1.pos.x-1 && n2.pos.y > n1.pos.y {
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
			if n2.pos.x == n1.pos.x+1 && n2.pos.y > n1.pos.y {
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

func d7_tree_unique_nodes(node *d7_node, seen map[d7_pos]bool) int {
	if node == nil {
		return 0
	}

	if seen[*node.pos] {
		return 0
	}

	seen[*node.pos] = true

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

package main

import (
	"aoc-2025/src/utils"
	"sort"
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
	xNodes := make(map[int][]*d7_node)

	// Group nodes by their x coordinate
	for _, p := range splitters {
		n := &d7_node{pos: &d7_pos{x: p.x, y: p.y}}
		xNodes[p.x] = append(xNodes[p.x], n)
	}

	// Sort each x group by y coordinate
	for _, list := range xNodes {
		sort.Slice(list, func(i, j int) bool {
			return list[i].pos.y < list[j].pos.y
		})
	}

	for x, list := range xNodes {
		for _, n := range list {
			// left child
			if leftList, ok := xNodes[x-1]; ok {
				i := sort.Search(len(leftList), func(i int) bool {
					return leftList[i].pos.y > n.pos.y
				})
				if i < len(leftList) {
					n.left = leftList[i]
				}
			}

			// right child
			if rightList, ok := xNodes[x+1]; ok {
				i := sort.Search(len(rightList), func(i int) bool {
					return rightList[i].pos.y > n.pos.y
				})
				if i < len(rightList) {
					n.right = rightList[i]
				}
			}
		}
	}

	return xNodes[splitters[0].x][0]
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

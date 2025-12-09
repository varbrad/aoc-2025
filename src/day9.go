package main

import (
	"aoc-2025/src/utils"
	"fmt"
)

var D9 = Day{
	P1:    D9P1,
	P2:    D9P2,
	Input: D9Input,
}

type d9_xy struct {
	x, y int
}

func D9P1(input string) int {
	rows := utils.SplitLines(input)

	cells := []d9_xy{}

	for _, row := range rows {
		var x, y int
		_, err := fmt.Sscanf(row, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		cells = append(cells, d9_xy{x: x, y: y})
	}

	max := -1
	maxA, maxB := d9_xy{}, d9_xy{}

	for i := 0; i < len(cells)-1; i++ {
		a := cells[i]
		for j := i + 1; j < len(cells); j++ {
			b := cells[j]

			dist := d9_abs(a.x-b.x+1) + d9_abs(a.y-b.y+1)
			if dist > max {
				max = dist
				maxA = a
				maxB = b
			}
		}
	}

	dx := d9_abs(maxA.x - maxB.x + 1)
	dy := d9_abs(maxA.y - maxB.y + 1)

	return dx * dy
}

type d9_bounds struct {
	x1, y1, x2, y2 int
}

func D9P2(input string) int {
	rows := utils.SplitLines(input)

	redCells := []d9_xy{}

	bounds := d9_bounds{x1: -1, y1: -1, x2: -1, y2: -1}
	for _, row := range rows {
		var x, y int
		_, err := fmt.Sscanf(row, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		if bounds.x1 == -1 || x < bounds.x1 {
			bounds.x1 = x
		}
		if bounds.y1 == -1 || y < bounds.y1 {
			bounds.y1 = y
		}
		if bounds.x2 == -1 || x > bounds.x2 {
			bounds.x2 = x
		}
		if bounds.y2 == -1 || y > bounds.y2 {
			bounds.y2 = y
		}
		redCells = append(redCells, d9_xy{x: x, y: y})
	}

	cellMap := map[d9_xy]rune{}

	d9_make_borders(bounds, redCells, cellMap)
	d9_flood_fill(bounds, cellMap)
	d9_print_map(bounds, cellMap)

	return -1

}

func d9_make_borders(bounds d9_bounds, redCells []d9_xy, cellMap map[d9_xy]rune) {
	// Add every red cell to the map
	for _, cell := range redCells {
		cellMap[cell] = '#'
	}

	// Now for every red cell, we need to find its orthagonal neighbors (it will have two)

	for _, cell := range redCells {
		neighbors := []d9_xy{}

		// Try and find the neighbour to our left
		for i := cell.x - 1; i >= bounds.x1; i-- {
			if v, ok := cellMap[d9_xy{x: i, y: cell.y}]; ok && v == '#' {
				// Found a red cell
				neighbors = append(neighbors, d9_xy{x: i, y: cell.y})
				break
			}
		}

		// Try and find the neighbour to our right
		for i := cell.x + 1; i <= bounds.x2; i++ {
			if v, ok := cellMap[d9_xy{x: i, y: cell.y}]; ok && v == '#' {
				// Found a red cell
				neighbors = append(neighbors, d9_xy{x: i, y: cell.y})
				break
			}
		}

		// Try and find the neighbour above
		for i := cell.y - 1; i >= bounds.y1; i-- {
			if v, ok := cellMap[d9_xy{x: cell.x, y: i}]; ok && v == '#' {
				// Found a red cell
				neighbors = append(neighbors, d9_xy{x: cell.x, y: i})
				break
			}
		}

		// Try and find the neighbour below
		for i := cell.y + 1; i <= bounds.y2; i++ {
			if v, ok := cellMap[d9_xy{x: cell.x, y: i}]; ok && v == '#' {
				// Found a red cell
				neighbors = append(neighbors, d9_xy{x: cell.x, y: i})
				break
			}
		}

		// Now for these neighbours, we need to draw the border lines
		for _, n := range neighbors {
			if n.x == cell.x {
				// Vertical line
				minY := cell.y
				maxY := n.y
				if minY > maxY {
					minY, maxY = maxY, minY
				}
				for y := minY + 1; y < maxY; y++ {
					cellMap[d9_xy{x: cell.x, y: y}] = 'X'
				}
			} else if n.y == cell.y {
				// Horizontal line
				minX := cell.x
				maxX := n.x
				if minX > maxX {
					minX, maxX = maxX, minX
				}
				for x := minX + 1; x < maxX; x++ {
					cellMap[d9_xy{x: x, y: cell.y}] = 'X'
				}
			}
		}
	}
}

func d9_flood_fill(bounds d9_bounds, cellMap map[d9_xy]rune) {
	// We must find a cell inside the shape first
	// Try the very middle of the grid

	start := d9_xy{x: 8, y: 2}
	cellMap[start] = 'O'

	queue := []d9_xy{start}

	directions := []d9_xy{
		{x: 0, y: -1},
		{x: 0, y: 1},
		{x: -1, y: 0},
		{x: 1, y: 0},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			neighbor := d9_xy{x: current.x + dir.x, y: current.y + dir.y}
			if _, ok := cellMap[neighbor]; !ok {
				// Empty cell, fill it
				cellMap[neighbor] = 'O'
				queue = append(queue, neighbor)
			}
		}
	}
}

func d9_print_map(bounds d9_bounds, cellMap map[d9_xy]rune) {
	for y := bounds.y1 - 1; y <= bounds.y2+1; y++ {
		for x := bounds.x1 - 1; x <= bounds.x2+1; x++ {
			if v, ok := cellMap[d9_xy{x: x, y: y}]; ok {
				fmt.Print(string(v))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func d9_abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package main

import (
	"aoc-2025/src/utils"
	"fmt"
	"math"
	"slices"
)

var D9 = Day{
	P1:    D9P1,
	P2:    D9P2,
	Input: D9Input,
}

type d9_xy struct {
	x, y float64
}

type d9_pair struct {
	a, b *d9_xy
	area int
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
		cells = append(cells, d9_xy{x: float64(x), y: float64(y)})
	}

	pairs := []d9_pair{}

	for i := 0; i < len(cells)-1; i++ {
		a := cells[i]
		for j := i + 1; j < len(cells); j++ {
			b := cells[j]

			area := (d9_abs(a.x-b.x) + 1) * (d9_abs(a.y-b.y) + 1)
			pairs = append(pairs, d9_pair{a: &a, b: &b, area: int(area)})
		}
	}

	slices.SortFunc(pairs, func(a, b d9_pair) int {
		return b.area - a.area
	})

	return pairs[0].area
}

func D9P2(input string) int {
	rows := utils.SplitLines(input)

	allCells := []d9_xy{}
	cellsX := []d9_xy{}
	cellsY := []d9_xy{}

	for _, row := range rows {
		var x, y int
		_, err := fmt.Sscanf(row, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		xy := d9_xy{x: float64(x), y: float64(y)}
		allCells = append(allCells, xy)
		cellsX = append(cellsX, xy)
		cellsY = append(cellsY, xy)
	}

	slices.SortFunc(cellsX, func(a, b d9_xy) int {
		if a.x == b.x {
			return int(a.y - b.y)
		}
		return int(a.x - b.x)
	})

	slices.SortFunc(cellsY, func(a, b d9_xy) int {
		if a.y == b.y {
			return int(a.x - b.x)
		}
		return int(a.y - b.y)
	})

	polygon := d9_polygon(allCells, cellsX, cellsY)

	pairs := []d9_pair{}

	for i := 0; i < len(allCells)-1; i++ {
		a := allCells[i]
		for j := i + 1; j < len(allCells); j++ {
			b := allCells[j]

			area := (d9_abs(a.x-b.x) + 1) * (d9_abs(a.y-b.y) + 1)
			pairs = append(pairs, d9_pair{a: &a, b: &b, area: int(area)})
		}
	}

	slices.SortFunc(pairs, func(a, b d9_pair) int {
		return b.area - a.area
	})

	// Now go through all pairs, and see if the rectangle they form intersects with any polygon edge
	for _, pair := range pairs {
		a := *pair.a
		b := *pair.b

		// Apply small inset to avoid overlapping edges
		minX := math.Min(a.x, b.x) + 0.5
		maxX := math.Max(a.x, b.x) - 0.5
		minY := math.Min(a.y, b.y) + 0.5
		maxY := math.Max(a.y, b.y) - 0.5

		rectPoints := []d9_xy{
			{x: minX, y: minY},
			{x: maxX, y: minY},
			{x: maxX, y: maxY},
			{x: minX, y: maxY},
		}

		intersects := false

	outer:
		for i := range polygon {
			p1 := polygon[i]
			p2 := polygon[(i+1)%len(polygon)]

			for j := range rectPoints {
				r1 := rectPoints[j]
				r2 := rectPoints[(j+1)%len(rectPoints)]

				if d9_lines_intersect(p1, p2, r1, r2) {
					intersects = true
					break outer
				}
			}
		}

		if !intersects {
			return pair.area
		}
	}

	return len(polygon)
}

func d9_polygon(allCells, cellsX, cellsY []d9_xy) []d9_xy {
	polygon := []d9_xy{allCells[0]}

	lastDir := ' '

outer:
	for i := 1; i < len(allCells); i++ {
		current := polygon[i-1]

		// Is it to the right of us?
		if lastDir != 'L' {
			for i := 0; i < len(cellsX); i++ {
				if cellsX[i].y == current.y && cellsX[i].x > current.x {
					polygon = append(polygon, cellsX[i])
					lastDir = 'R'
					continue outer
				}
			}
		}

		// Is it below us?
		if lastDir != 'U' {
			for i := 0; i < len(cellsY); i++ {
				if cellsY[i].x == current.x && cellsY[i].y > current.y {
					polygon = append(polygon, cellsY[i])
					lastDir = 'D'
					continue outer
				}
			}
		}

		// Is it above us?
		if lastDir != 'D' {
			for i := len(cellsY) - 1; i >= 0; i-- {
				if cellsY[i].x == current.x && cellsY[i].y < current.y {
					polygon = append(polygon, cellsY[i])
					lastDir = 'U'
					continue outer
				}
			}
		}

		// Is it to the left of us?
		if lastDir != 'R' {
			for i := len(cellsX) - 1; i >= 0; i-- {
				if cellsX[i].y == current.y && cellsX[i].x < current.x {
					polygon = append(polygon, cellsX[i])
					lastDir = 'L'
					continue outer
				}
			}
		}
	}

	return polygon
}

func d9_lines_intersect(a1, a2, b1, b2 d9_xy) bool {
	// Compute orientation of ordered triplet (p, q, r)
	orient := func(p, q, r d9_xy) float64 {
		return (q.y-p.y)*(r.x-q.x) - (q.x-p.x)*(r.y-q.y)
	}

	o1 := orient(a1, a2, b1)
	o2 := orient(a1, a2, b2)
	o3 := orient(b1, b2, a1)
	o4 := orient(b1, b2, a2)

	// General case
	if o1*o2 < 0 && o3*o4 < 0 {
		return true
	}

	// Collinear + on-segment checks
	onSegment := func(p, q, r d9_xy) bool {
		return q.x >= min(p.x, r.x) &&
			q.x <= max(p.x, r.x) &&
			q.y >= min(p.y, r.y) &&
			q.y <= max(p.y, r.y)
	}

	if o1 == 0 && onSegment(a1, b1, a2) {
		return true
	}
	if o2 == 0 && onSegment(a1, b2, a2) {
		return true
	}
	if o3 == 0 && onSegment(b1, a1, b2) {
		return true
	}
	if o4 == 0 && onSegment(b1, a2, b2) {
		return true
	}

	return false
}

func d9_abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

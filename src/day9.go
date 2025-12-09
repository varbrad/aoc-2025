package main

import (
	"aoc-2025/src/utils"
	"fmt"
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

	max := -1
	maxA, maxB := d9_xy{}, d9_xy{}

	for i := 0; i < len(cells)-1; i++ {
		a := cells[i]
		for j := i + 1; j < len(cells); j++ {
			b := cells[j]

			dist := d9_abs(int(a.x-b.x+1)) + d9_abs(int(a.y-b.y+1))
			if dist > max {
				max = dist
				maxA = a
				maxB = b
			}
		}
	}

	dx := d9_abs(int(maxA.x - maxB.x + 1))
	dy := d9_abs(int(maxA.y - maxB.y + 1))

	return dx * dy
}

type d9_bounds struct {
	x1, y1, x2, y2 float64
}

type d9_pair struct {
	a, b d9_xy
	area int
}

func D9P2(input string) int {
	rows := utils.SplitLines(input)
	redCells := []d9_xy{}
	bounds := d9_bounds{-1, -1, -1, -1}

	for _, row := range rows {
		var xI, yI int
		_, err := fmt.Sscanf(row, "%d,%d", &xI, &yI)
		x := float64(xI)
		y := float64(yI)
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

	pairs := []d9_pair{}
	for i := 0; i < len(redCells)-1; i++ {
		a := redCells[i]
		for j := i + 1; j < len(redCells); j++ {
			b := redCells[j]

			area := d9_abs(int(a.x-b.x+1)) * d9_abs(int(a.y-b.y+1))
			pairs = append(pairs, d9_pair{a: a, b: b, area: area})
		}
	}

	slices.SortFunc(pairs, func(i, j d9_pair) int {
		return j.area - i.area
	})

	// Construct a polygon from the red cells
	polygon := d9_polygon(bounds, redCells)

	for ix, pair := range pairs {
		if d9_rectangle_inside_poly(polygon, pair) {
			fmt.Println("Found at index", ix)
			fmt.Println("Pair:", pair)
			return pair.area
		}
	}

	return -1
}

func d9_rectangle_inside_poly(polygon []d9_xy, pair d9_pair) bool {
	n := len(polygon)

	minX := min(pair.a.x, pair.b.x) + 0.01
	maxX := max(pair.a.x, pair.b.x) - 0.01
	minY := min(pair.a.y, pair.b.y) + 0.01
	maxY := max(pair.a.y, pair.b.y) - 0.01

	rect := []d9_xy{
		{minX, minY},
		{maxX, minY},
		{maxX, maxY},
		{minX, maxY},
	}

	// rectangle edges: (0→1), (1→2), (2→3), (3→0)
	for i := 0; i < 4; i++ {
		r1 := rect[i]
		r2 := rect[(i+1)%4]

		for j := 0; j < n; j++ {
			p1 := polygon[j]
			p2 := polygon[(j+1)%n]

			if d9_linesIntersect(r1, r2, p1, p2) {
				return false
			}
		}
	}
	return true
}

// Helper function: cross product of two vectors (p1->p2) × (p1->p3)
func d9_cross(p1, p2, p3 d9_xy) float64 {
	return (p2.x-p1.x)*(p3.y-p1.y) - (p2.y-p1.y)*(p3.x-p1.x)
}

// Check if line segments a1->a2 and b1->b2 intersect
func d9_linesIntersect(a1, a2, b1, b2 d9_xy) bool {
	c1 := d9_cross(a1, a2, b1)
	c2 := d9_cross(a1, a2, b2)
	c3 := d9_cross(b1, b2, a1)
	c4 := d9_cross(b1, b2, a2)

	return c1*c2 <= 0 && c3*c4 <= 0
}

func d9_polygon(bounds d9_bounds, cells []d9_xy) []d9_xy {
	pofloat64s := make([]d9_xy, len(cells))

	// Begin with a random cell
	pofloat64s[0] = cells[0]
	lastDir := ' '
outer:
	for ix := 1; ix < len(cells); ix++ {
		// Find the next pofloat64 in the polygon, it will be orthogonal to the last
		curr := pofloat64s[ix-1]

		for _, candidate := range cells {
			if candidate == curr {
				continue
			}
			// is it to our left?
			if lastDir != 'R' {
				for i := curr.x - 1; i >= bounds.x1; i-- {
					if candidate.x == i && candidate.y == curr.y {
						pofloat64s[ix] = candidate
						lastDir = 'L'
						continue outer
					}
				}
			}
			// is it to our right?
			if lastDir != 'L' {
				for i := curr.x + 1; i <= bounds.x2; i++ {
					if candidate.x == i && candidate.y == curr.y {
						pofloat64s[ix] = candidate
						lastDir = 'R'
						continue outer
					}
				}
			}
			// is it above us?
			if lastDir != 'D' {
				for i := curr.y - 1; i >= bounds.y1; i-- {
					if candidate.y == i && candidate.x == curr.x {
						pofloat64s[ix] = candidate
						lastDir = 'U'
						continue outer
					}
				}
			}
			// is it below us?
			if lastDir != 'U' {
				for i := curr.y + 1; i <= bounds.y2; i++ {
					if candidate.y == i && candidate.x == curr.x {
						pofloat64s[ix] = candidate
						lastDir = 'D'
						continue outer
					}
				}
			}
		}
	}

	return pofloat64s
}

func d9_abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

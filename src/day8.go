package main

import (
	"aoc-2025/src/utils"
	"fmt"
	"slices"
)

var D8 = Day{
	P1:    D8P1,
	P2:    D8P2,
	Input: D8Input,
}

func D8P1(input string) int {
	state := d8_parse(input)
	coords := state.pos

	dists := []d8_dist{}

	for i := 0; i < len(coords)-1; i++ {
		a := coords[i]
		for j := i + 1; j < len(coords); j++ {
			b := coords[j]

			dist := a.distance(b)

			dists = append(dists, d8_dist{a, b, dist})
		}
	}

	slices.SortFunc(dists, func(a, b d8_dist) int {
		return a.dist - b.dist
	})

	circuits := [][]*d8_3d{}

	for _, dist := range dists[:state.n] {
		// We will find a matching circuit
		for cix, circuit := range circuits {
			for _, point := range circuit {
				if point == dist.a || point == dist.b {
					circuits[cix] = append(circuits[cix], dist.a, dist.b)
				}
			}
		}

		// otherwise, make a new circuit
		circuits = append(circuits, []*d8_3d{dist.a, dist.b})
	}

	// Try and combine circuits
outer:
	for i := len(circuits) - 1; i > 0; i-- {
		ca := circuits[i]
		for j := i - 1; j >= 0; j-- {
			cb := circuits[j]

			// see if they share any points
			shared := false
		inner:
			for _, pa := range ca {
				for _, pb := range cb {
					if pa == pb {
						shared = true
						break inner
					}
				}
			}

			if shared {
				circuits[j] = append(circuits[j], circuits[i]...)
				circuits = append(circuits[:i], circuits[i+1:]...)
				continue outer
			}
		}
	}

	// make each circuit a unique set
	for cix, circuit := range circuits {
		unique := []*d8_3d{}
		seen := map[*d8_3d]bool{}

		for _, point := range circuit {
			if !seen[point] {
				unique = append(unique, point)
				seen[point] = true
			}
		}

		circuits[cix] = unique
	}

	slices.SortFunc(circuits, func(i, j []*d8_3d) int {
		return len(j) - len(i)
	})

	sum := 1
	for _, p := range circuits[:3] {
		sum *= len(p)
	}

	return sum
}

func D8P2(input string) int {
	return -1
}

type d8_3d struct {
	x, y, z int
}

type d8_dist struct {
	a, b *d8_3d
	dist int
}

func (vec3d *d8_3d) distance(other *d8_3d) int {
	dx := vec3d.x - other.x
	dy := vec3d.y - other.y
	dz := vec3d.z - other.z

	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	if dz < 0 {
		dz = -dz
	}

	return dx*dx + dy*dy + dz*dz
}

type d8_state struct {
	pos []*d8_3d
	n   int
}

func d8_parse(input string) d8_state {
	rows := utils.SplitLines(input)

	result := make([]*d8_3d, len(rows)-1)
	for ix, row := range rows[1:] {
		var point d8_3d
		fmt.Sscanf(row, "%d,%d,%d", &point.x, &point.y, &point.z)
		result[ix] = &point
	}

	slices.SortFunc(result, func(a, b *d8_3d) int {
		return a.x - b.x
	})

	return d8_state{
		pos: result,
		n:   utils.ToIntMust(rows[0]),
	}
}

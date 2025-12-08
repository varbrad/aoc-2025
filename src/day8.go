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

	circuits := make([][]*d8_3d, len(state.nodes))
	for i, p := range state.nodes {
		circuits[i] = []*d8_3d{p}
	}

	for i := 0; i < state.n; i++ {
		d8_merge_circuits(&circuits, state.connections[i])
	}

	slices.SortFunc(circuits, func(a, b []*d8_3d) int {
		return len(b) - len(a)
	})

	prod := 1
	for _, circuit := range circuits[:3] {
		prod *= len(circuit)
	}

	return prod
}

func D8P2(input string) int {
	state := d8_parse(input)

	circuits := make([][]*d8_3d, len(state.nodes))
	for i, p := range state.nodes {
		circuits[i] = []*d8_3d{p}
	}

	i := 0
	for {
		conn := state.connections[i]
		d8_merge_circuits(&circuits, conn)

		if len(circuits) == 1 {
			return conn.a.x * conn.b.x
		}

		i++
	}
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

func d8_merge_circuits(circuits *[][]*d8_3d, conn *d8_dist) {
	// Find the circuits for a and b
	var caix, cbix int
	for cix, circuit := range *circuits {
		for _, point := range circuit {
			if point == conn.a {
				caix = cix
			}
			if point == conn.b {
				cbix = cix
			}
		}
	}

	if caix != cbix {
		// Merge circuits
		(*circuits)[caix] = append((*circuits)[caix], (*circuits)[cbix]...)
		*circuits = append((*circuits)[:cbix], (*circuits)[cbix+1:]...)
	}
}

type d8_state struct {
	nodes       []*d8_3d
	connections []*d8_dist
	n           int
}

func d8_parse(input string) d8_state {
	rows := utils.SplitLines(input)

	nodes := make([]*d8_3d, len(rows)-1)
	for ix, row := range rows[1:] {
		var point d8_3d
		fmt.Sscanf(row, "%d,%d,%d", &point.x, &point.y, &point.z)
		nodes[ix] = &point
	}

	slices.SortFunc(nodes, func(a, b *d8_3d) int {
		return a.x - b.x
	})

	connections := []*d8_dist{}

	for i := 0; i < len(nodes)-1; i++ {
		a := nodes[i]
		for j := i + 1; j < len(nodes); j++ {
			b := nodes[j]

			dist := a.distance(b)

			connections = append(connections, &d8_dist{a, b, dist})
		}
	}

	slices.SortFunc(connections, func(a, b *d8_dist) int {
		return a.dist - b.dist
	})

	return d8_state{
		nodes:       nodes,
		connections: connections,
		n:           utils.ToIntMust(rows[0]),
	}
}

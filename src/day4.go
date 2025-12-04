package main

import (
	"aoc-2025/src/utils"
)

func D4() {
	utils.RunDays(D4P1, D4P2, D4Input)
}

func D4P1(input string) int {
	grid := d4_parse(input)

	count := 0

	for y := range grid.h {
		for x := range grid.w {
			char := grid.cells[y][x]
			if char != '@' {
				continue
			}

			adjacent := d4_countAdjacent(grid, x, y)
			if adjacent < 4 {
				count++
			}
		}
	}

	return count
}

func D4P2(input string) int {
	grid := d4_parse(input)

	count := 0
	flag := false

	for {
		for y := range grid.h {
			for x := range grid.w {
				char := grid.cells[y][x]
				if char != '@' {
					continue
				}

				adjacent := d4_countAdjacent(grid, x, y)
				if adjacent < 4 {
					count++
					grid.cells[y][x] = '.'
					flag = true
				}
			}
		}

		if !flag {
			break
		}

		flag = false
	}

	return count
}

type d4_grid struct {
	cells [][]rune
	w     int
	h     int
}

func d4_parse(input string) *d4_grid {
	lines := utils.SplitLines(input)

	cells := make([][]rune, len(lines))

	for i, line := range lines {
		cells[i] = make([]rune, len(line))
		for j, char := range line {
			cells[i][j] = char
		}
	}

	return &d4_grid{
		cells: cells,
		w:     len(lines[0]),
		h:     len(lines),
	}
}

func d4_countAdjacent(grid *d4_grid, x, y int) int {
	count := 0
	directions := [][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for _, dir := range directions {
		dx := dir[0]
		dy := dir[1]

		nx := x + dx
		ny := y + dy

		if nx >= 0 && nx < grid.w && ny >= 0 && ny < grid.h {
			if grid.cells[ny][nx] == '@' {
				count++
			}
		}
	}

	return count
}

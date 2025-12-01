package main

import (
	"aoc-2025/src/utils"
	"fmt"
)

func main() {
	// Read input from file or other source
	input := utils.ReadInput("input/day1")

	p1 := D1P1(input)
	p2 := D1P2(input)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func D1P1(input string) int {
	rows := utils.SplitLines(input)

	dial := 50
	onZero := 0

	for _, row := range rows {
		forward := row[0] == 'R'
		steps := utils.ToInt(row[1:])

		if forward {
			dial += steps
		} else {
			dial -= steps
		}

		dial = (dial + 100) % 100

		if dial == 0 {
			onZero++
		}
	}

	return onZero
}

func D1P2(input string) int {
	rows := utils.SplitLines(input)

	dial := 50
	passedZero := 0

	for _, row := range rows {
		forward := row[0] == 'R'
		steps := utils.ToInt(row[1:])

		for range steps {
			if forward {
				dial++
			} else {
				dial--
			}

			dial = (dial + 100) % 100

			if dial == 0 {
				passedZero++
			}
		}
	}

	return passedZero
}

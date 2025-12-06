package main

import (
	"aoc-2025/src/utils"
)

var D1 = Day{
	P1:    D1P1,
	P2:    D1P2,
	Input: D1Input,
}

func D1P1(input string) int {
	rows := utils.SplitLines(input)

	dial := 50
	onZero := 0

	for _, row := range rows {
		forward := row[0] == 'R'
		steps := utils.ToIntMust(row[1:])

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
		steps := utils.ToIntMust(row[1:])

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

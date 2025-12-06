package main

import (
	"aoc-2025/src/utils"
	"strings"
)

func D6() {
	utils.RunDays(D6P1, D6P2, D6Input)
}

func D6P1(input string) int {
	state := d6_parse(input)
	sum := 0
	for ix, op := range state.ops {
		l := []int{}
		for i := 0; i < state.num_lists; i++ {
			l = append(l, state.ns[i][ix])
		}
		switch op {
		case '+':
			sum += utils.SumInts(l)
		case '*':
			sum += utils.ProdInts(l)
		}
	}

	return sum
}

func D6P2(input string) int {
	rows := utils.SplitLinesNoTrim(input)

	sum := 0

	ops := []rune{}
	for _, item := range strings.Fields(rows[len(rows)-1]) {
		ops = append(ops, rune(item[0]))
	}

	lenRow1 := len(rows[0])

	group := []int{}
	for ix := lenRow1 - 1; ix >= 0; ix-- {
		// Scan through each column
		colChars := []rune{}

		for j, row := range rows {
			// Skip the last row (operations)
			if j == len(rows)-1 {
				continue
			}

			colChar := rune(row[ix])

			if colChar != ' ' {
				colChars = append(colChars, colChar)
			}
		}

		str := string(colChars)

		if str == "" || ix == 0 {
			if ix == 0 {
				group = append(group, utils.ToIntMust(str))
			}
			lastOp := ops[len(ops)-1]
			ops = ops[:len(ops)-1]

			switch lastOp {
			case '+':
				sum += utils.SumInts(group)
			case '*':
				sum += utils.ProdInts(group)
			}

			group = []int{}
		} else {
			group = append(group, utils.ToIntMust(str))
		}
	}

	return sum
}

type d6_state struct {
	ns        [][]int
	ops       []rune
	num_lists int
}

func d6_parse(input string) d6_state {
	rows := utils.SplitLines(input)
	lenRows := len(rows)

	ns := make([][]int, lenRows-1)
	ops := []rune{}

	for ix, row := range rows {
		items := strings.Fields(row)

		if ix == lenRows-1 {
			// ops
			for _, item := range items {
				ops = append(ops, rune(item[0]))
			}
			break
		} else {
			nums := []int{}
			for _, item := range items {
				nums = append(nums, utils.ToIntMust(item))
			}
			ns[ix] = nums
		}
	}

	return d6_state{
		ns:        ns,
		ops:       ops,
		num_lists: len(ns),
	}
}

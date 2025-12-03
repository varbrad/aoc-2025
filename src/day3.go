package main

import (
	"aoc-2025/src/utils"
	"strings"
)

func D3() {
	utils.RunDays(D3P1, D3P2, D3Input)
}

func D3P1(input string) int {
	return d3_solve(input, 2)
}

func D3P2(input string) int {
	return d3_solve(input, 12)
}

func d3_solve(input string, n int) int {
	rows := d3_parse(input)
	sum := 0
	for _, row := range rows {
		val := d3_recurse(row, n, 0, len(row)-n)
		str := ""
		for _, v := range val {
			str += utils.ToString(v)
		}
		sum += utils.ToIntMust(str)
	}
	return sum
}

func d3_recurse(row []int, n int, ix1 int, ix2 int) []int {
	if n == 0 {
		return []int{}
	}
	ix, value := d3_findHighest(row, ix1, ix2)
	rest := d3_recurse(row, n-1, ix+1, ix2+1)
	return append([]int{value}, rest...)
}

func d3_findHighest(row []int, ix1 int, ix2 int) (ix int, value int) {
	ix, value = -1, -1
	for i := ix1; i <= ix2; i++ {
		if row[i] > value {
			ix = i
			value = row[i]
		}
	}
	return ix, value
}

func d3_parse(input string) [][]int {
	rows := utils.SplitLines(input)
	result := make([][]int, len(rows))
	for i, row := range rows {
		ints := strings.Split(row, "")
		result[i] = make([]int, len(ints))
		for j, s := range ints {
			v := utils.ToIntMust(s)
			result[i][j] = v
		}
	}
	return result
}

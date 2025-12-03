package main

import (
	"aoc-2025/src/utils"
	"strings"
)

func D3() {
	utils.RunDays(D3P1, D3P2, D3Input)
}

func D3P1(input string) int {
	rows := parseD3(input)

	sum := 0

	for _, row := range rows {

		minIx, minVal := -1, -1

		for i := 0; i < len(row)-1; i++ {
			if row[i] > minVal {
				minIx = i
				minVal = row[i]
			}
		}

		maxVal := -1

		for j := len(row) - 1; j > minIx; j-- {
			if row[j] > maxVal {
				maxVal = row[j]
			}
		}

		sum += (minVal * 10) + maxVal
	}

	return sum
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

func D3P2(input string) int {
	rows := parseD3(input)

	sum := 0

	for _, row := range rows {
		ix1, val1 := d3_findHighest(row, 0, len(row)-12)
		ix2, val2 := d3_findHighest(row, ix1+1, len(row)-11)
		ix3, val3 := d3_findHighest(row, ix2+1, len(row)-10)
		ix4, val4 := d3_findHighest(row, ix3+1, len(row)-9)
		ix5, val5 := d3_findHighest(row, ix4+1, len(row)-8)
		ix6, val6 := d3_findHighest(row, ix5+1, len(row)-7)
		ix7, val7 := d3_findHighest(row, ix6+1, len(row)-6)
		ix8, val8 := d3_findHighest(row, ix7+1, len(row)-5)
		ix9, val9 := d3_findHighest(row, ix8+1, len(row)-4)
		ix10, val10 := d3_findHighest(row, ix9+1, len(row)-3)
		ix11, val11 := d3_findHighest(row, ix10+1, len(row)-2)
		_, val12 := d3_findHighest(row, ix11+1, len(row)-1)

		str := utils.ToString(val1) + utils.ToString(val2) + utils.ToString(val3) + utils.ToString(val4) + utils.ToString(val5) +
			utils.ToString(val6) + utils.ToString(val7) + utils.ToString(val8) + utils.ToString(val9) + utils.ToString(val10) + utils.ToString(val11) + utils.ToString(val12)

		sum += utils.ToIntMust(str)
	}

	return sum
}

func parseD3(input string) [][]int {
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

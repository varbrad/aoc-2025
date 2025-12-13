package main

import (
	"aoc-2025/src/utils"
	"flag"
	"fmt"
	"maps"
	"os"
	"slices"
)

type Day struct {
	P1    func(string) int
	P2    func(string) int
	Input string
}

var Days = map[int]Day{
	1:  D1,
	2:  D2,
	3:  D3,
	4:  D4,
	5:  D5,
	6:  D6,
	7:  D7,
	8:  D8,
	9:  D9,
	10: D10,
	11: D11,
	12: D12,
}

func benchmark() {
	days := []int{}
	dayNumbers := maps.Keys(Days)
	for day := range dayNumbers {
		days = append(days, day)
	}
	slices.Sort(days)

	averageRuntime := int64(0)

	for _, day := range days {
		dayStruct := Days[day]

		solve1 := dayStruct.P1(dayStruct.Input) != -1
		solve2 := dayStruct.P2(dayStruct.Input) != -1

		p1 := utils.BenchmarkFn(dayStruct.P1, dayStruct.Input)
		p2 := utils.BenchmarkFn(dayStruct.P2, dayStruct.Input)

		fmt.Println("Day", day)
		// Colour red if not solved
		if solve1 {
			averageRuntime += p1.AverageNs
			fmt.Printf("  Part 1: %v (avg), %.2f (iter/s), \n", utils.ToFriendlyTime(p1.AverageNs), p1.IterationsPerSecond)
		} else {
			fmt.Println("\033[31m  Part 1: Not solved\033[0m")
		}
		if solve2 {
			averageRuntime += p2.AverageNs
			fmt.Printf("  Part 2: %v (avg), %.2f (iter/s), \n", utils.ToFriendlyTime(p2.AverageNs), p2.IterationsPerSecond)
		} else {
			fmt.Println("\033[31m  Part 2: Not solved\033[0m")
		}
	}

	fmt.Println("")
	fmt.Println("Overall total avg runtime: ", utils.ToFriendlyTime(averageRuntime))
}

func main() {
	dayPtr := flag.Int("day", -1, "which day to run")
	benchPtr := flag.Bool("bench", false, "whether to run in benchmark mode")
	flag.Parse()

	day := *dayPtr
	bench := *benchPtr

	if bench {
		benchmark()
		return
	}

	if day == -1 {
		fmt.Println("Please specify a day to run with -day=N")
		os.Exit(1)
	}

	if dayStruct, ok := Days[day]; ok {
		utils.RunDays(dayStruct.P1, dayStruct.P2, dayStruct.Input)
	} else {
		fmt.Printf("Day %d is not implemented yet.\n", day)
		os.Exit(1)
	}
}

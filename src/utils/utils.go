package utils

import (
	"fmt"
	"strings"
	"time"
)

func SplitLines(input string) []string {
	slice := strings.Split(strings.TrimSpace(input), "\n")
	for i := range slice {
		slice[i] = strings.TrimSpace(slice[i])
	}
	return slice
}

func SplitLinesNoTrim(input string) []string {
	slice := strings.Split(input, "\n")
	return slice
}

func ToInt(s string) (int, error) {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		return 0, fmt.Errorf("could not convert string to int: %q", s)
	}
	return n, nil
}

func ToIntMust(s string) int {
	n, err := ToInt(s)
	if err != nil {
		panic(err)
	}
	return n
}

func TimeFn[T any](fn func() T) (T, int64) {
	start := time.Now()
	result := fn()
	elapsedNs := time.Since(start).Nanoseconds()
	return result, elapsedNs
}

func ToString(v any) string {
	return fmt.Sprintf("%v", v)
}

func ToFriendlyTime(ns int64) string {
	if ns < 1_000 {
		return fmt.Sprintf("%d ns", ns)
	}
	if ns < 1_000_000 {
		return fmt.Sprintf("%.2f µs", float64(ns)/1_000)
	}
	if ns < 1_000_000_000 {
		return fmt.Sprintf("%.2f ms", float64(ns)/1_000_000)
	}
	return fmt.Sprintf("%.2f s", float64(ns)/1_000_000_000)
}

func RunDays(p1, p2 func(i string) int, input string) {
	r1, t1 := TimeFn(func() int { return p1(input) })
	r2, t2 := TimeFn(func() int { return p2(input) })

	fmt.Println("Part 1:", r1, "("+ToFriendlyTime(t1)+")")
	fmt.Println("Part 2:", r2, "("+ToFriendlyTime(t2)+")")
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Reduce[ArrType any, RetValue any](slice []ArrType, fn func(RetValue, ArrType) RetValue, initial RetValue) RetValue {
	result := initial

	for _, v := range slice {
		result = fn(result, v)
	}

	return result
}

func SumInts(ints []int) int {
	return Reduce(ints, func(acc int, v int) int {
		return acc + v
	}, 0)
}

func ProdInts(ints []int) int {
	return Reduce(ints, func(acc int, v int) int {
		return acc * v
	}, 1)
}

type BenchmarkResult struct {
	TotalIterations     int
	IterationsPerSecond float64
	TotalNs             int64
	AverageNs           int64
}

func BenchmarkFn(fn func(string) int, input string) BenchmarkResult {
	start := time.Now()
	iters := 0
	for {
		iters++
		fn(input)
		if iters >= 1000 || time.Since(start).Milliseconds() >= 500 {
			break
		}
	}
	totalNs := time.Since(start).Nanoseconds()
	avgNs := totalNs / int64(iters)

	return BenchmarkResult{
		TotalIterations:     iters,
		IterationsPerSecond: float64(iters) / (float64(totalNs) / 1_000_000_000),
		TotalNs:             totalNs,
		AverageNs:           avgNs,
	}
}

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

func toFriendlyTime(ns int64) string {
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

	fmt.Println("Part 1:", r1, "("+toFriendlyTime(t1)+")")
	fmt.Println("Part 2:", r2, "("+toFriendlyTime(t2)+")")
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

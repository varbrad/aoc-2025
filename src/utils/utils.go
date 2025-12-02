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
	elapsedMs := time.Since(start).Milliseconds()
	return result, elapsedMs
}

func ToString(v any) string {
	return fmt.Sprintf("%v", v)
}

func RunDays(p1, p2 func(i string) int, input string) {
	r1, t1 := TimeFn(func() int { return p1(input) })
	r2, t2 := TimeFn(func() int { return p2(input) })

	fmt.Println("Part 1:", r1, "("+ToString(t1)+"ms)")
	fmt.Println("Part 2:", r2, "("+ToString(t2)+"ms)")
}

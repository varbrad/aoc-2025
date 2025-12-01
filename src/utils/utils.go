package utils

import (
	"fmt"
	"strings"
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

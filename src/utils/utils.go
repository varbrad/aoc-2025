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

func ToInt(s string) int {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		panic(err)
	}
	return n
}

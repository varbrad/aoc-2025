package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}

func SplitLines(input string) []string {
	slice := strings.Split(input, "\n")
	for i := range slice {
		slice[i] = strings.TrimSpace(slice[i])
	}
	return slice
}

func ReadInputAndSplitLines(path string) []string {
	input := ReadInput(path)
	return SplitLines(input)
}

func ToInt(s string) int {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		panic(err)
	}
	return n
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

func TestD7P1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  testInput,
			expect: 21,
		},
		{
			name:   "Input",
			input:  D7Input,
			expect: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D7P1(tt.input))
		})
	}
}

func TestD7P2(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  testInput,
			expect: 40,
		},
		// {
		// 	name:   "Input",
		// 	input:  D7Input,
		// 	expect: 0,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D7P2(tt.input))
		})
	}
}

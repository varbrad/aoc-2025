package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d8_test_input = `
10
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func TestD8P1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  d8_test_input,
			expect: 40,
		},
		// {
		// 	name:   "Input",
		// 	input:  D8.Input,
		// 	expect: 164475,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D8P1(tt.input))
		})
	}
}

func TestD8P2(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  d8_test_input,
			expect: 25272,
		},
		{
			name:   "Input",
			input:  D8.Input,
			expect: 169521198,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D8P2(tt.input))
		})
	}
}

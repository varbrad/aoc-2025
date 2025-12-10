package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d10_test_input = `
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
[###] (0) (1) (2) {1,1,1}
`

func TestD10P1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  d10_test_input,
			expect: 10,
		},
		{
			name:   "Input",
			input:  D10.Input,
			expect: 375,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D10P1(tt.input))
		})
	}
}

// func TestD10P2(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		input  string
// 		expect int
// 	}{
// 		{
// 			name:   "Example Test 1",
// 			input:  d10_test_input,
// 			expect: 0,
// 		},
// 		{
// 			name:   "Input",
// 			input:  D10.Input,
// 			expect: 1501292304,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			assert.Equal(t, tt.expect, D10P2(tt.input))
// 		})
// 	}
// }

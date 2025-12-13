package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d12_test_input = `
0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2
`

func TestD12P1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		// {
		// 	name:   "Example Test 1",
		// 	input:  d12_test_input,
		// 	expect: 2,
		// },
		// {
		// 	name:   "Input",
		// 	input:  D12.Input,
		// 	expect: 0,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D12P1(tt.input))
		})
	}
}

// func TestD12P2(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		input  string
// 		expect int
// 	}{
// 		{
// 			name:   "Example Test 1",
// 			input:  d11_p2_test_input,
// 			expect: 2,
// 		},
// 		{
// 			name:   "Input",
// 			input:  D12.Input,
// 			expect: 0,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			assert.Equal(t, tt.expect, D12P2(tt.input))
// 		})
// 	}
// }

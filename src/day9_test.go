package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d9_test_input = `
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestD9P1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  d9_test_input,
			expect: 50,
		},
		{
			name:   "Input",
			input:  D9.Input,
			expect: 4763932976,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D9P1(tt.input))
		})
	}
}

func TestD9P2(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  d9_test_input,
			expect: 24,
		},
		// {
		// 	name:   "Input",
		// 	input:  D9.Input,
		// 	expect: 0,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D9P2(tt.input))
		})
	}
}

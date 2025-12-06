package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD6P1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  "123 328  51 64 \n45 64  387 23 \n 6 98  215 314\n*   +   *   +  ",
			expect: 4277556,
		},
		{
			name:   "Input",
			input:  D6Input,
			expect: 4693159084994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D6P1(tt.input))
		})
	}
}

func TestD6P2(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  ",
			expect: 3263827,
		},
		{
			name:   "Input",
			input:  D6Input,
			expect: 11643736116335,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D6P2(tt.input))
		})
	}
}

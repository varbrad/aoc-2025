package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD3P1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  "987654321111111\n811111111111119\n234234234234278\n818181911112111",
			expect: 357,
		},
		{
			name:   "Input",
			input:  D3Input,
			expect: 17332,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D3P1(tt.input))
		})
	}
}

func TestD3P2(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  "987654321111111\n811111111111119\n234234234234278\n818181911112111",
			expect: 3121910778619,
		},
		{
			name:   "Input",
			input:  D3Input,
			expect: 172516781546707,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D3P2(tt.input))
		})
	}
}

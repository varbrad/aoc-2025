package main

import (
	"testing"
)

func TestD1P1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82",
			expect: 3,
		},
		{
			name:   "Example Test 2",
			input:  "R50\nL50\nR50\nL50",
			expect: 2,
		},
		{
			name:   "Input",
			input:  D1Input,
			expect: 1_168,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := D1P1(tt.input); got != tt.expect {
				t.Errorf("D1P1() = %v, expected %v", got, tt.expect)
			}
		})
	}
}

func TestD1P2(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82",
			expect: 6,
		},
		{
			name:   "Example Test 2",
			input:  "R50\nL50\nR50\nL50",
			expect: 2,
		},
		{
			name:   "Input",
			input:  D1Input,
			expect: 7_199,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := D1P2(tt.input); got != tt.expect {
				t.Errorf("D1P2() = %v, expected %v", got, tt.expect)
			}
		})
	}
}

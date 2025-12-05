package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD5P1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32",
			expect: 3,
		},
		{
			name:   "Input",
			input:  D5Input,
			expect: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D5P1(tt.input))
		})
	}
}

func TestD5P2(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32",
			expect: 14,
		},
		{
			name:   "Example Test 2",
			input:  "1-20\n2-4\n7-11\n3-9\n4-5\n9-15\n\n1\n5",
			expect: 20,
		},
		{
			name:   "Input",
			input:  D5Input,
			expect: 353507173555373,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D5P2(tt.input))
		})
	}
}

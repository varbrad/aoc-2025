package utils

import (
	"testing"
)

func TestSplitLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Basic Test",
			input:    "line1\nline2\nline3",
			expected: []string{"line1", "line2", "line3"},
		},
		{
			name:     "With Extra Spaces",
			input:    "  line1  \nline2 \n  line3",
			expected: []string{"line1", "line2", "line3"},
		},
		{
			name:     "Single Line",
			input:    "onlyline",
			expected: []string{"onlyline"},
		},
		{
			name:     "Empty Input",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "Leading/Trailing Newline/spaces ignored",
			input:    " \nline1\nline2\n",
			expected: []string{"line1", "line2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitLines(tt.input)
			if len(result) != len(tt.expected) {
				t.Fatalf("expected length %d, got %d", len(tt.expected), len(result))
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("at index %d, expected %q, got %q", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Basic Number",
			input:    "42",
			expected: 42,
		},
		{
			name:     "Negative Number",
			input:    "-17",
			expected: -17,
		},
		{
			name:     "Zero",
			input:    "0",
			expected: 0,
		},
		{
			name:     "With Leading/Trailing Spaces",
			input:    "   123   ",
			expected: 123,
		},
		{
			name:     "Large Number",
			input:    "987654321",
			expected: 987_654_321,
		},
		{
			name:     "With Plus Sign",
			input:    "+56",
			expected: 56,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToInt(tt.input)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

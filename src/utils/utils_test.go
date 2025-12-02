package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			t.Parallel()
			result := SplitLines(tt.input)
			assert.Equal(t, tt.expected, result)
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
			t.Parallel()
			result, err := ToInt(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}

	t.Run("Invalid Input", func(t *testing.T) {
		t.Parallel()
		result, err := ToInt("invalid")
		assert.Error(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestToString(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{name: "Integer", input: 42, expected: "42"},
		{name: "String", input: "hello", expected: "hello"},
		{name: "Float", input: 3.14, expected: "3.14"},
		{name: "Boolean True", input: true, expected: "true"},
		{name: "Boolean False", input: false, expected: "false"},
		{name: "Nil", input: nil, expected: "<nil>"},
		{name: "Slice", input: []int{1, 2, 3}, expected: "[1 2 3]"},
		{name: "Map", input: map[string]int{"a": 1, "b": 2}, expected: "map[a:1 b:2]"},
		{name: "Struct", input: struct {
			Name  string
			Other int
		}{Name: "Test", Other: 498}, expected: "{Test 498}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTimeFn(t *testing.T) {
	t.Run("Basic Function", func(t *testing.T) {
		t.Parallel()
		fn := func() int {
			sum := 0
			for i := 1; i <= 50; i++ {
				sum += i
			}
			return sum
		}
		result, elapsed := TimeFn(fn)
		assert.Equal(t, 1275, result)
		assert.GreaterOrEqual(t, elapsed, int64(0))
	})
}

func TestToIntMust(t *testing.T) {
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
			t.Parallel()
			result := ToIntMust(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}

	t.Run("Invalid Input", func(t *testing.T) {
		t.Parallel()
		assert.PanicsWithError(t, "could not convert string to int: \"invalid\"", func() {
			ToIntMust("invalid")
		})
	})
}

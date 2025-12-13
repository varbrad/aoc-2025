package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d11_test_input = `
aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
`

func TestD11P1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  d11_test_input,
			expect: 5,
		},
		{
			name:   "Input",
			input:  D11.Input,
			expect: 603,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D11P1(tt.input))
		})
	}
}

var d11_p2_test_input = `
svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
`

func TestD11P2(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Example Test 1",
			input:  d11_p2_test_input,
			expect: 2,
		},
		{
			name:   "Input",
			input:  D11.Input,
			expect: 380961604031372,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, D11P2(tt.input))
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0013_roman_to_integer(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{"I", 1},
		{"IV", 4},
		{"XIV", 14},
		{"XXXI", 31},
		{"XLI", 41},
		{"CMXCIX", 999},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0013_roman_to_integer(c.input))
	}
}

func lc_0013_roman_to_integer(s string) int {
	last, res, m := 1001, 0, map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for _, v := range s {
		val := m[v]
		if val > last {
			res += val - 2*last
		} else {
			res += val
		}
		last = val
	}
	return res
}

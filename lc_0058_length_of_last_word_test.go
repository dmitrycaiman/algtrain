package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/length-of-last-word/
func Test_lc_0058_length_of_last_word(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{"kek lol haha", 4}, {"", 0}, {"123 1       ", 1}, {"   ", 0},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0058_length_of_last_word(c.input))
	}
}

// Итерация от конца строки до первого пробела. Учесть, что в конце строки тоже могут быть пробелы, их нужно пропустить.
func lc_0058_length_of_last_word(s string) int {
	start, res := false, 0
	for i := len(s) - 1; i >= 0; i-- {
		switch {
		case !start && s[i] != ' ':
			start = true
			fallthrough
		case start && s[i] != ' ':
			res++
		case start && s[i] == ' ':
			return res
		}
	}
	return res
}

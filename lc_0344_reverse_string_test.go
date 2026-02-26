package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-string
func Test_lc_0344_reverse_string(t *testing.T) {
	cases := []struct {
		input, result string
	}{
		{"hi", "ih"},
		{"hello", "olleh"},
		{"h", "h"},
		{"bebra", "arbeb"},
	}
	for _, c := range cases {
		input := []byte(c.input)
		lc_0344_reverse_string(input)
		assert.Equal(t, []byte(c.result), input)
	}
}

// Двумя метками с концов идём в середину и меняемся символам, пока не встретимся.
func lc_0344_reverse_string(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

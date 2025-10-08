package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_1446_consecutive_characters(t *testing.T) {
	cases := []struct {
		input  string
		result int
	}{
		{"aaaaabbccdddddddefg", 7},
		{"f", 1},
		{"adawegweg", 1},
		{"aaaaaaga", 6},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_1446_consecutive_characters(c.input))
	}
}

func lc_1446_consecutive_characters(s string) int {
	result := 1
	for i, j := 0, 1; j <= len(s); j++ {
		if j == len(s) || s[j-1] != s[j] {
			if l := j - i; l > result {
				result = l
			}
			i = j
		}
	}
	return result
}

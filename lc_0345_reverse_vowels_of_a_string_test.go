package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-vowels-of-a-string
func Test_lc_0345_reverse_vowels_of_a_string(t *testing.T) {
	cases := []struct {
		input, output string
	}{
		{"aabee", "eebaa"},
		{"aeibbb", "ieabbb"},
		{"icecrEAm", "AcEcreim"},
		{"ai", "ia"},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0345_reverse_vowels_of_a_string(c.input))
	}
}

// Двумя метками идём с концов. Сначала одной меткой ищем гласную, потом другой. Если нашли, меняемся.
// Если метки встретились, то выходим: осталась либо одна средняя гласная, либо не осталось вообще.
func lc_0345_reverse_vowels_of_a_string(s string) string {
	isVowel, b := func(b byte) bool {
		switch b {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		}
		return false
	}, []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		for ; i < j && !isVowel(b[i]); i++ {
		}
		for ; i < j && !isVowel(b[j]); j-- {
		}
		if i < j {
			b[i], b[j], i, j = b[j], b[i], i+1, j-1
		}
	}
	return string(b)
}

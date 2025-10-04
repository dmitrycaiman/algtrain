package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-palindromic-substring
func Test_lc_0005_longest_palindromic_substring(t *testing.T) {
	cases := []struct{ input, output string }{
		{"estdgsadgwescbaabcergreghe", "cbaabc"},
		{"aaaaaaaa", "aaaaaaaa"},
		{"aabbbcccc", "cccc"},
		{"aaaabcacbartyutr", "abcacba"},
		{"aaaabcacbartyutrabcacbar", "rabcacbar"},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0005_longest_palindromic_substring(c.input))
	}
}

// Ставим две метки: на начало и конец массива. Двигая правую метку, проверяем каждое слово в границах меток на палиндром.
// При нахождении палиндрома сохраняем его границы. Далее сдвигаем левую метку и повторяем действия с правой меткой.
// Таким образом мы переберём все возможные подстроки. Важно прерывать исследование текущих границ,
// если их размер заведомо меньше, чем длина уже найденного ранее палиндрома.
func lc_0005_longest_palindromic_substring(s string) string {
	check := func(i, j int) bool {
		for ; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	a, b := 0, 0
	for i := range s {
		for j := len(s) - 1; j > 0; j-- {
			if j-i <= b-a {
				break
			}
			if check(i, j) {
				a, b = i, j
				break
			}
		}
	}
	return s[a : b+1]
}

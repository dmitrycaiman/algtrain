package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/is-subsequence
func Test_lc_0392_is_subsequence(t *testing.T) {
	cases := []struct {
		s, t   string
		result bool
	}{{"kek", "keeeeek", true}, {"kek", "kep", false}, {"", "keeeeek", true}, {"", "", true}}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0392_is_subsequence(c.s, c.t))
	}
}

// 0ms (100%), 3.9MB (99.85%)
// Две метки. Двигаем первую по целевой строке, вторую по подстроке.
// При совпадении символов сдвигаем вторую. Как только вторая добралась до конца подстроки, выходим с успехом.
// Учитываем, что пустая строка является подстрокой к любой строке.
func lc_0392_is_subsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(t) && j != len(s); i++ {
		if t[i] == s[j] {
			j++
			if j == len(s) {
				return true
			}
		}
	}
	return j == len(s)
}

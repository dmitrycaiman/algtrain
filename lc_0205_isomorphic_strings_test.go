package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/isomorphic-strings
func Test_lc_0205_isomorphic_strings(t *testing.T) {
	cases := []struct {
		s, t   string
		result bool
	}{
		{"tree", "broo", true},
		{"topkek", "farlol", true},
		{"tree", "trea", false},
		{"a", "b", true},
		{"abcdefg", "bcdefgh", true},
		{"abcdefg", "abcdeff", false},
		{"paper", "title", true},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0205_isomorphic_strings(c.s, c.t))
	}
}

// 100/70
// Строки нужно проверить на то, что в них каждый символ соответствует какому-то другому "1 к 1".
// Организуем два словаря: один ставит соответствие символов s -> t, другой t -> s.
// Идём по строкам и начинаем заполнять словари. Если оба символа отсутствуют в обоих словарях, ставим их во взаимное соответствие.
// Если оба символа присутствуют, то проверяем правильность соответствия и при отрицательном результате сразу выходим с неудачей.
// Также выходим с неудачей, если в одном словаре есть соответствие, а в другом нет (в этом случае соответствие будет не "1 к 1").
func lc_0205_isomorphic_strings(s string, t string) bool {
	ms, mt := map[byte]byte{}, map[byte]byte{}
	for i := range s {
		symt, oks := ms[s[i]]
		_, okt := mt[t[i]]
		switch {
		case !oks && !okt:
			ms[s[i]], mt[t[i]] = t[i], s[i]
		case !oks || !okt || symt != t[i]:
			return false
		}
	}
	return true
}

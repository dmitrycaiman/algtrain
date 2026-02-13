package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/word-pattern
func Test_lc_0290_word_pattern(t *testing.T) {
	cases := []struct {
		pattern, s string
		result     bool
	}{
		{"ababb", "bob kak bob kak kak", true},
		{"ababb", "bob kak bob kak kak kek", false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0290_word_pattern(c.pattern, c.s))
	}
}

// При несовпадении количества слов и длины паттерна сразу выходим с неудачей.
// Далее итерируемся по элементам и собираем 2 мапы: слов в символы и символов в слова.
// Итерация будет успешной, только если встреченная комбинация слова и символа:
// - ещё не задана в мапах (тогда мы её задаём);
// - уже задана и совпадает с содержимым мап.
func lc_0290_word_pattern(pattern, s string) bool {
	words, m1, m2 := strings.Split(s, " "), map[string]byte{}, map[byte]string{}
	if len(words) != len(pattern) {
		return false
	}
	for i := range words {
		p, ok1 := m1[words[i]]
		w, ok2 := m2[pattern[i]]
		switch {
		case !ok1 && !ok2:
			m1[words[i]], m2[pattern[i]] = pattern[i], words[i]
		case ok1 && ok2 && p == pattern[i] && w == words[i]:
		default:
			return false
		}
	}
	return true
}

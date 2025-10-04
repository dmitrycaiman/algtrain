package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-anagram
func Test_lc_0242_valid_anagram(t *testing.T) {
	cases := []struct {
		s, t   string
		result bool
	}{
		{"kobaaa", "kaboaa", true}, {"abcdef", "fedcba", true}, {"abcde", "fedcba", false}, {"anagram", "nagaram", true}, {"anagram", "naga", false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0242_valid_anagram(c.s, c.t))
	}
}

// 0ms (100%), 4.68MB (99.41%)
// Складываем в мапу все символы первого слова, в значении используем количество их появления в слове.
// Далее проходимся по символам второго слова и соответствующий элемент мапы уменьшаем на 1 или вовсе удаляем.
// Сразу выходим с неудачей, если текущего символа нет в мапе.
// По итогу мапа должна стать пустой, что и скажет об успешности результата.
func lc_0242_valid_anagram(s, t string) bool {
	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]--
		switch m[v] {
		case 0:
			delete(m, v)
		case -1:
			return false
		}
	}
	return len(m) == 0
}

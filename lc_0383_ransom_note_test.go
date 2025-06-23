package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/ransom-note
func Test_lc_0383_ransom_note(t *testing.T) {
	cases := []struct {
		ransomNote, magazine string
		result               bool
	}{
		{"abc", "abcd", true},
		{"abc", "ab", false},
		{"a", "", false},
		{"", "a", true},
		{"aa", "ab", false},
		{"aa", "aba", true},
		{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			true,
		},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0383_ransom_note(c.ransomNote, c.magazine))
	}
}

// 0ms (100%), 5.7MB (85.99%)
// Используем в качестве мапы слайс фиксированной величины 26 элементов, где 0 — символ 'a', 25 — символ 'z'.
// Проходимся по алфавиту magazine и инкрементируем элемент по индексу, соответствующему порядковому номеру ASCII-кода символа в рамках строчных английских букв.
// Затем проходимся по слову и так же декрементируем элементы слайса. Если встречаем нулевой элемент — выходим с неудачей (значит, элемент не встречался
// или был использован максимально возможное число раз).
func lc_0383_ransom_note(ransomNote string, magazine string) bool {
	s := make([]uint16, 'z'-'a'+1)
	for _, v := range magazine {
		s[v-'a']++
	}
	for _, v := range ransomNote {
		if s[v-'a'] == 0 {
			return false
		}
		s[v-'a']--
	}
	return true
}

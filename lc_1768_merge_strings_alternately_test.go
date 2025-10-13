package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-strings-alternately
func Test_lc_1768_merge_strings_alternately(t *testing.T) {
	cases := []struct{ word1, word2, result string }{
		{"abc", "defgh", "adbecfgh"},
		{"abcef", "d", "adbcef"},
		{"aaaa", "aaaa", "aaaaaaaa"},
		{"abc", "", "abc"},
		{"", "", ""},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_1768_merge_strings_alternately(c.word1, c.word2))
	}
}

// 0/100, 4/87
// Заводим общий счётчик позиции символа в слове.
// Проверяем, что счётчик не превышает длины каждого из слов (сначала первого, потом второго),
// и добавляем символ слова на позиции, соответствующей счётчику.
// Также мы знаем, что количество взятий равно сумме длин слов — соответственно ограничиваем этим количество наших взятий.
func lc_1768_merge_strings_alternately(word1, word2 string) string {
	result := []byte{}
	for i, j := 0, 0; j < len(word1)+len(word2); i++ {
		if i < len(word1) {
			result = append(result, word1[i])
			j++
		}
		if i < len(word2) {
			result = append(result, word2[i])
			j++
		}
	}
	return string(result)
}

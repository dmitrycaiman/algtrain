package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-common-prefix
func Test_lc_0014_longest_common_prefix(t *testing.T) {
	cases := []struct {
		input  []string
		output string
	}{
		{[]string{"ab", "ac", "ad"}, "a"},
		{[]string{"ab", "ab", "ab"}, "ab"},
		{[]string{"abc", "abd", "abe"}, "ab"},
		{[]string{"abc", "abd", ""}, ""},
		{[]string{"abcd"}, "abcd"},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0014_longest_common_prefix(c.input))
	}
}

// Заводим массив байт, в который будем складывать символы, одинаковые у всех строк на идущих подряд от начала позициях.
// Итерируемся по символам первой строки. Если у всех строк на текущей позиции есть данный символ, то добавляем его в массив результатов.
// Если у какой-либо из строк на этой позиции другой символ, или позиция находится за пределами её длины,
// то сразу возвращаем накопленный результат.
func lc_0014_longest_common_prefix(strs []string) string {
	res := []byte{}
	for i := range len(strs[0]) {
		checked := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 || strs[j][i] != checked {
				return string(res)
			}
		}
		res = append(res, checked)
	}
	return string(res)
}

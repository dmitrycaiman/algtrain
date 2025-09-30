package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0003_longest_substring_without_repeating_characters(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{"aaaaabcc", 3},
		{"abba", 2},
		{"abcccccqwertyui", 9},
		{"doggies", 4},
		{"fraudrate", 6},
		{"a", 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0003_longest_substring_without_repeating_characters(c.input))
	}
}

// Итерируемся по массиву двумя метками. Они обозначают начало и конец участка без повторений.
// Левая метка изначально стоит на нулевом элементе, правая идёт по массиву.
// Сохраняем в хеш-таблице встречаемые символы и их последнюю встреченную позицию.
// Если мы ранее уже видели символ, то смещаем левую метку на его предыдущую позицию из хеш-таблицы.
// Если не видели, то считаем длину участка без повторений и сохраняем её при необходимости.
func lc_0003_longest_substring_without_repeating_characters(s string) int {
	m, maxLength := map[byte]int{}, 0
	for i, j := 0, 0; j < len(s); j++ {
		pos, ok := m[s[j]]
		if ok && i <= pos {
			i = pos + 1
		} else {
			if length := j - i + 1; length > maxLength {
				maxLength = length
			}
		}
		m[s[j]] = j
	}
	return maxLength
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/edit-distance
func Test_lc_0072_edit_distance(t *testing.T) {
	cases := []struct {
		word1, word2 string
		result       int
	}{
		{"abc", "bbc", 1},
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "aaa", 3},
		{"aaa", "", 3},
		{"", "", 0},
		{"aaa", "aaa", 0},
		{"a", "ab", 1},
		{"dinitrophenylhydrazine", "acetylphenylhydrazine", 6},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0072_edit_distance(c.word1, c.word2))
	}
}

// Рекурсивный перебор с мемоизацией.
// В начале рекурсивной функции рассматриваем все краевые случаи,
// а потом пытаемся сделать все возможные действия и выбираем наилучший результат.
// Результат запоминаем в хранилище и используем повторно при необходимости.
func lc_0072_edit_distance(word1, word2 string) int {
	return lc_0072([]byte(word1), []byte(word2), 0, 0, map[[2]int]int{})
}

func lc_0072(b1, b2 []byte, i, j int, storage map[[2]int]int) (output int) {
	result, ok := storage[[2]int{i, j}]
	if ok {
		return result
	}
	defer func() {
		storage[[2]int{i, j}] = output
	}()

	switch {
	case i == len(b1) && j == len(b2):
		return 0
	case i == len(b1) && j != len(b2):
		return len(b2) - j
	case i != len(b1) && j == len(b2):
		return len(b1) - i
	case b1[i] == b2[j]:
		switch {
		case i == len(b1)-1 && j != len(b2)-1:
			return len(b2) - j - 1
		case i != len(b1)-1 && j == len(b2)-1:
			return len(b1) - i - 1
		case i != len(b1)-1 && j != len(b2)-1:
			return lc_0072(b1, b2, i+1, j+1, storage)
		default:
			return 0
		}
	}

	return 1 + min(
		lc_0072(b1, b2, i+1, j+1, storage), // Замена.
		lc_0072(b1, b2, i, j+1, storage),   // Вставка.
		lc_0072(b1, b2, i+1, j, storage),   // Удаление.
	)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string
func Test_lc_0028_find_the_index_of_the_first_occurence_in_a_string(t *testing.T) {
	cases := []struct {
		haystack, needle string
		output           int
	}{
		{"abaaaaaba", "ab", 0},
		{"abaaaaaba", "ba", 1},
		{"abaaaaaba", "aaba", 5},
		{"abaaaaaba", "ba", 1},
		{"abaaaaaba", "bab", -1},
		{"ab", "aba", -1},
		{"ab", "aba", -1},
		{"abababc", "babc", 3},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0028_find_the_index_of_the_first_occurence_in_a_string(c.haystack, c.needle))
	}
}

// 0ms (100%), 3.88MB (95.03%)
// Две метки: позиция символа в стоге и в игле. Итерируемся по стогу. Если встречаем совпадение символов, то двигаем метку иглы вперёд.
// Если не встречаем, то сдвигаем метку стога назад на следующую от той, где было предполагаемое начало иглы, а метку иглы в начало.
// Заканчиваем движение, если одна из меток вышла за пределы своей строки.
// Алгоритм считается успешным, если после итераций мы прошли всю иглу. В этом случае возвращаем разницу меток, чтобы получить начало иглы.
func lc_0028_find_the_index_of_the_first_occurence_in_a_string(haystack, needle string) int {
	i, j := 0, 0
	for ; i < len(haystack) && j < len(needle); i++ {
		if haystack[i] == needle[j] {
			j++
		} else {
			i -= j
			j = 0
		}
	}
	if j != len(needle) {
		return -1
	}
	return i - j
}

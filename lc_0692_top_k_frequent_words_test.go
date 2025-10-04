package main

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/top-k-frequent-words
func Test_lc_0692_top_k_frequent_words(t *testing.T) {
	cases := []struct {
		words  []string
		k      int
		result []string
	}{
		{[]string{"ka", "ko", "ko", "ka", "be", "be", "be", "be", "fu", "fu"}, 3, []string{"be", "fu", "ka"}},
		{[]string{"be", "be", "be", "fu", "fu", "fu"}, 1, []string{"be"}},
		{[]string{"a", "b", "c"}, 3, []string{"a", "b", "c"}},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0692_top_k_frequent_words(c.words, c.k))
	}
}

// Проходимся по массиву строк и собираем хеш-таблицу частот, а также массив ключей этой таблицы.
// Сортируем массив ключей встроенными средствами, опираясь на значения частот. Возвращаем первые k результатов.
func lc_0692_top_k_frequent_words(words []string, k int) []string {
	m, keys := map[string]int{}, []string{}
	for _, v := range words {
		m[v]++
		if m[v] == 1 {
			keys = append(keys, v)
		}
	}
	slices.SortStableFunc(
		keys,
		func(a, b string) int {
			switch {
			case m[a] > m[b]:
				return -1
			case m[a] < m[b]:
				return 1
			default:
				return strings.Compare(a, b)
			}
		},
	)
	return keys[:k]
}

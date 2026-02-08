package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/combinations
func Test_lc_0077_combinations(t *testing.T) {
	cases := []struct {
		n, k   int
		result [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{1, 1, [][]int{{1}}},
	}
	for _, c := range cases {
		assert.True(t, CheckEqualityWithFunc(c.result, lc_0077_combinations(c.n, c.k), func(a, b []int) bool { return CheckEquality(a, b) }))
	}
}

// Рекурсивный метод "берём или не берём".
func lc_0077_combinations(n, k int) [][]int {
	return lc_0077(1, n, k, []int{}, [][]int{})
}

func lc_0077(i, n, k int, set []int, result [][]int) [][]int {
	// i есть число, которые мы "берём или не берём" в набор.
	// Проверяем, хватит ли оставшихся чисел для достижения длины набора k.
	if n-i+1 < k-len(set) {
		return result
	}
	set2 := make([]int, len(set))
	copy(set2, set)
	set2 = append(set2, i)
	if len(set2) < k {
		// Если набор не достиг длины k, то продолжаем.
		// В этой ветке мы "берём" число i в набор.
		result = lc_0077(i+1, n, k, set2, result)
	} else {
		// Если набор готов, то сохраняем результат.
		result = append(result, set2)
	}
	// В этой ветке мы "не берём" число i в набор.
	return lc_0077(i+1, n, k, set, result)
}

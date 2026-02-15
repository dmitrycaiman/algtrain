package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/contains-duplicate-ii
func Test_lc_0219_contains_duplicate_ii(t *testing.T) {
	cases := []struct {
		input  []int
		k      int
		result bool
	}{
		{[]int{1, 2, 3, 4, 5, 1, 6, 7, 1}, 3, true},
		{[]int{1, 2, 3, 2, 5, 6, 7, 2}, 3, true},
		{[]int{1, 2, 1, 2, 5, 6, 5, 2}, 1, false},
		{[]int{1, 2, 1, 2, 2, 6, 5, 2}, 1, true},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0219_contains_duplicate_ii(c.input, c.k))
	}
}

// Проходим по массиву. Собираем в мапу индексы встречающихся элементов.
// Если на очередной итерации находим элемент, который уже встречался, то вычисляем разницу между индексами.
// Если разница не удовлетворяет k, то обновляем индекс в мапе на текущий. Если удолветворяет, то мы нашли ответ.
func lc_0219_contains_duplicate_ii(nums []int, k int) bool {
	m := map[int]int{}
	for i, v := range nums {
		if j, ok := m[v]; !ok || i-j > k {
			m[v] = i
			continue
		}
		return true
	}
	return false
}

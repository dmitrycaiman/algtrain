package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-consecutive-sequence
func Test_lc_0128_longest_consecutive_sequence(t *testing.T) {
	cases := []struct {
		input  []int
		result int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, 1, 2}, 3},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0128_longest_consecutive_sequence(c.input))
	}
}

// Собираем все элементы в мапу. Итерируемся по мапе.
// Создаём рекурсивную функцию, которая для каждого элемента сначала удаляет его из мапы, а затем проверяет, есть ли в мапе соседние
// элементы, то есть +1 и -1 от него. Если какой-то из соседних элементов существует, рекурсивная функция запускается и для него.
// Каждый вход в рекурсию прибавляет единицу к результату. По итогу по выходу из рекурсии получим общее число соседних элементов.
// Повторного поиска не происходит, так как при рассмотрении элемента, принадлежащего промежутку, все элементы промежутка по выходу
// из рекурсии будут удалены.
func lc_0128_longest_consecutive_sequence(nums []int) int {
	best, m := 0, map[int]struct{}{}
	for _, n := range nums {
		m[n] = struct{}{}
	}
	for n := range m {
		best = max(best, lc_0128(n, m))
	}
	return best
}

func lc_0128(n int, m map[int]struct{}) int {
	delete(m, n)
	result := 1
	if _, ok := m[n+1]; ok {
		result += lc_0128(n+1, m)
	}
	if _, ok := m[n-1]; ok {
		result += lc_0128(n-1, m)
	}
	return result
}

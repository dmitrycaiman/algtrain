package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted
func Test_lc_0167_two_sum_ii_input_array_is_sorted(t *testing.T) {
	cases := []struct {
		numbers, result []int
		target          int
	}{
		{[]int{1, 2}, []int{1, 2}, 3},
		{[]int{2, 7, 11, 15}, []int{1, 2}, 9},
		{[]int{2, 3, 5, 7, 9}, []int{2, 3}, 8},
		{[]int{0, 2, 3, 5, 7, 9}, []int{1, 6}, 9},
		{[]int{-8, -7, -4, -1, 2, 3, 5, 7, 9}, []int{2, 8}, 0},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0167_two_sum_ii_input_array_is_sorted(c.numbers, c.target))
	}
}

// Используем тот факт, что массив отсортирован: ставим метки на его концы и двигаем их навстречу.
// Считаем сумму на метках: если она больше целевой, сдвигаем правую (уменьшаемся), если меньше — левую (увеличиваемся).
// Как только попадаем метками в целевое значение, возвращаем результат с прибавлением единицы к индексам (по условию — 1-indexed array).
func lc_0167_two_sum_ii_input_array_is_sorted(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; ; {
		sum := numbers[i] + numbers[j]
		switch {
		case sum == target:
			return []int{i + 1, j + 1}
		case sum < target:
			i++
		default:
			j--
		}
	}
}

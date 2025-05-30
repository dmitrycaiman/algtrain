package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0035_search_insert_position(t *testing.T) {
	cases := []struct {
		nums           []int
		target, output int
	}{
		{[]int{}, 1, 0},
		{[]int{0}, 1, 1},
		{[]int{2}, 1, 0},
		{[]int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3}, 4, 3},
		{[]int{1, 2, 4}, 3, 2},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0035_search_insert_position(c.nums, c.target))
	}
}

// Итерируемся по массиву. Если встретили равный или больший элемент, то возвращаем их индекс.
// Если таких не встретили, то значит, что все элементы массива меньше целевого, и мы должны "встать вперёд".
func lc_0035_search_insert_position(nums []int, target int) int {
	for i, v := range nums {
		if v == target || v > target {
			return i
		}
	}
	return len(nums)
}

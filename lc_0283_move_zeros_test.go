package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/move-zeroes
func Test_lc_0283_move_zeros(t *testing.T) {
	cases := []struct {
		nums   []int
		result []int
	}{
		{[]int{1, 0, 2, 0, 3, 0, 0, 0, 0}, []int{1, 2, 3, 0, 0, 0, 0, 0, 0}},
		{[]int{1, 0, 0, 0, 0, 1}, []int{1, 1, 0, 0, 0, 0}},
		{[]int{1, 0, 0, 0, 0}, []int{1, 0, 0, 0, 0}},
		{[]int{0, 0, 1, 0, 1}, []int{1, 1, 0, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
	}
	for _, c := range cases {
		lc_0283_move_zeros(c.nums)
		assert.Equal(t, c.result, c.nums)
	}
}

// Два указателя. Сначала первый указатель ищет ноль, потом второй ищет "не ноль" перед первым.
// Как нашли, то меняем местами данные элементы. Выходим, когда один из указателей достиг конца.
func lc_0283_move_zeros(nums []int) {
	for i, j := 0, 1; i < len(nums) && j < len(nums); {
		switch {
		case nums[i] != 0:
			i++
		case j <= i:
			j = i + 1
		case nums[j] == 0:
			j++
		default:
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

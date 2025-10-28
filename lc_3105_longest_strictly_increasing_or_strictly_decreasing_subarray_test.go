package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray
func Test_lc_3105_longest_strictly_increasing_or_strictly_decreasing_subarray(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{[]int{1, 2, 1, 1, 1, 1}, 2},
		{[]int{1, 1, 2, 1, 1, 1, 1}, 2},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 1},
		{[]int{1, 1, 5, 4, 3, 9, 8, 7, 6, 5, 5, 5}, 5},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_3105_longest_strictly_increasing_or_strictly_decreasing_subarray(c.nums))
	}
}

// 100/93
// Два указателя на нулевом элементе. Задаёмся отношением первых двух элементов как текущим.
// Итерируемся и сверяем два соседних элемента. Если отношение сохраняется, двигаем правый указатель.
// Если нет, считаем длину, на протяжении которой последнее отношение сохранялось, и двигаем левый указатель.
// Причём если отношение пропало равенством двух элементов, двигаем левый указатель на позицию второго элемента.
// Если отношение просто "поменяло полярность", то сохраняем новое отношение
// и двигаем левый указатель на позицию первого сравниваемого элемента.
func lc_3105_longest_strictly_increasing_or_strictly_decreasing_subarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	increasing, result := nums[0] < nums[1], 1
	for i, j, shift := 0, 1, 0; j <= len(nums); j, shift = j+1, 0 {
		switch {
		case j == len(nums):
		case increasing && nums[j-1] < nums[j] || !increasing && nums[j-1] > nums[j]:
			continue
		case nums[j-1] == nums[j]:
			shift = 1
		default:
			increasing = nums[j-1] < nums[j]
		}
		if l := j - i; l > result {
			result = l
		}
		i = j - 1 + shift
	}
	return result
}

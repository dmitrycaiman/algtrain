package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/search-in-rotated-sorted-array-ii
func Test_lc_0081_search_in_rotated_sorted_array_ii(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		result bool
	}{
		{[]int{1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1}, 2, true},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 2, false},
		{[]int{1, 2, 3, 4, 5, 1, 1, 1}, 5, true},
		{[]int{2, 3, 1}, 1, true},
		{[]int{2, 3, 3, 3, 1}, 1, true},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0081_search_in_rotated_sorted_array_ii(c.nums, c.target))
	}
}

// Бинарный поиск. Основное отличите в том, что вывод о переходе в левую или правую часть делается на основе того,
// какая из частей точно отсортирована. Отсортированную часть определяем по величинам начала и середины массива.
// Наличие дублей может привести к такой ситуации, когда все границы равны между собой. В этом случае мы их сужаем.
func lc_0081_search_in_rotated_sorted_array_ii(nums []int, target int) bool {
	for {
		m := len(nums) / 2
		switch {
		case len(nums) <= 1:
			return len(nums) != 0 && nums[0] == target
		case nums[m] == target:
			return true
		case nums[0] == nums[m] && nums[m] == nums[len(nums)-1]:
			nums = nums[1 : len(nums)-1]
		case nums[0] <= nums[m]:
			if nums[0] <= target && target < nums[m] {
				nums = nums[:m]
			} else {
				nums = nums[m+1:]
			}
		case nums[0] > nums[m]:
			if nums[m] < target && target <= nums[len(nums)-1] {
				nums = nums[m+1:]
			} else {
				nums = nums[:m]
			}
		}
	}
}

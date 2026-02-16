package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array
func Test_lc_0034_find_first_and_last_position_of_element_in_sorted_array(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{1, 2, 3, 3, 3, 4, 5, 6}, 3, []int{2, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 10, []int{5, 5}},
		{[]int{5, 7, 7, 8, 8, 10}, 11, []int{-1, -1}},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0034_find_first_and_last_position_of_element_in_sorted_array(c.nums, c.target))
	}
}

// Бинарным поиском находим индекс целевого числа.
// Если нашли, то двигаемся в обе стороны от него, чтобы определить протяжённость промежутка повторения.
func lc_0034_find_first_and_last_position_of_element_in_sorted_array(nums []int, target int) []int {
	targetIndex := -1
loop:
	for i, j := 0, len(nums)-1; i <= j; {
		m := i + (j-i)/2
		switch {
		case nums[m] == target:
			targetIndex = m
			fallthrough
		case i == j:
			break loop
		case nums[m] < target:
			i = m + 1
		case nums[m] > target:
			j = m - 1
		}
	}
	if targetIndex == -1 {
		return []int{-1, -1}
	}
	for i, j := targetIndex-1, targetIndex+1; ; {
		switch {
		case i >= 0 && nums[i] == target:
			i--
		case j < len(nums) && nums[j] == target:
			j++
		default:
			return []int{i + 1, j - 1}
		}
	}
}

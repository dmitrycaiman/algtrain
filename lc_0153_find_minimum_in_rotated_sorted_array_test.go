package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array
func Test_lc_0153_find_minimum_in_rotated_sorted_array(t *testing.T) {
	cases := []struct {
		input  []int
		result int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{5, 1, 2, 3, 4}, 1},
		{[]int{4, 5, 1, 2, 3}, 1},
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{2, 3, 4, 5, 1}, 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0153_find_minimum_in_rotated_sorted_array(c.input))
	}
}

// Бинарный поиск. Минимум всегда лежит в точке перелома сдвинутого массива, соответственно мы стремимся туда попасть.
// Как только в массиве остаётся два элемента, мы возвращаем наименьший.
// Учитываем случай, когда массив после сдвига принял нормальное состояние, в этом случае возвращаем его первый элемент.
func lc_0153_find_minimum_in_rotated_sorted_array(nums []int) int {
	for {
		m := len(nums) / 2
		switch {
		case len(nums) == 2:
			if nums[0] < nums[1] {
				return nums[0]
			}
			return nums[1]
		case nums[0] <= nums[m] && nums[m] <= nums[len(nums)-1]:
			return nums[0]
		case nums[0] > nums[m]:
			nums = nums[:m+1]
		default:
			nums = nums[m+1:]
		}
	}
}

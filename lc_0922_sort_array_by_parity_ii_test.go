package main

import (
	"testing"
)

// https://leetcode.com/problems/sort-array-by-parity-ii
func Test_lc_0922_sort_array_by_parity_ii(t *testing.T) {
	cases := []struct{ input []int }{
		{[]int{0, 1, 2, 3}},
		{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		{[]int{4, 4, 4, 4, 3, 3, 1, 1}},
	}
	for _, c := range cases {
		for i, v := range lc_0922_sort_array_by_parity_ii(c.input) {
			if i%2 != v%2 {
				t.Errorf("%v%%2 != %v%%2", i, v)
			}
		}
	}
}

// 100/81
// Два указателя. Один на чётном индексе, другой на нечётном.
// Двигаем каждый из указателей на 2 позиции, если он стоит на верном элементе.
// В момент, когда на обоих указателях стоят неверные элементы, меняем элементы местами.
func lc_0922_sort_array_by_parity_ii(nums []int) []int {
	for i, j := 0, 1; i < len(nums) && j < len(nums); {
		switch {
		case nums[i]%2 != 0 && nums[j]%2 != 1:
			nums[i], nums[j] = nums[j], nums[i]
			i, j = i+2, j+2
		case nums[i]%2 == 0:
			i += 2
		case nums[j]%2 == 1:
			j += 2
		}
	}
	return nums
}

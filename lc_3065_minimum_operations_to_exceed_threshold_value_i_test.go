package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-i
func Test_lc_3065_minimum_operations_to_exceed_threshold_value_i(t *testing.T) {
	cases := []struct {
		nums      []int
		k, result int
	}{
		{[]int{1, 1, 1, 1, 2, 4, 3, 5, 6, 1}, 2, 5},
		{[]int{1, 1, 2}, 2, 2},
		{[]int{5, 6, 7, 8, 9}, 9, 4},
		{[]int{9, 8, 5, 4, 0, 0, 7, 6, 4, 2, 1, 9}, 9, 10},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_3065_minimum_operations_to_exceed_threshold_value_i(c.nums, c.k))
	}
}

// Просто считаем, сколько элементов в массиве меньше, чем k.
func lc_3065_minimum_operations_to_exceed_threshold_value_i(nums []int, k int) (result int) {
	for _, v := range nums {
		if v < k {
			result++
		}
	}
	return
}

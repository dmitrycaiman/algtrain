package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/contains-duplicate
func Test_lc_0217_contains_duplicate(t *testing.T) {
	cases := []struct {
		input  []int
		result bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1}, false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0217_contains_duplicate(c.input))
	}
}

// Сортируем и ищем одинаковые рядом стоящие элементы.
func lc_0217_contains_duplicate(nums []int) bool {
	slices.Sort(nums)
	for i, last := 1, nums[0]; i < len(nums); i++ {
		if nums[i] == last {
			return true
		}
		last = nums[i]
	}
	return false
}

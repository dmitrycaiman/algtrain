package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0189_rotate_array(t *testing.T) {
	cases := []struct {
		input, output []int
		k             int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 1, 2, 3, 4}, 1},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 1, 2, 3}, 2},
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 1}, 4},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 2, 3, 4, 5}, []int{5, 1, 2, 3, 4}, 6},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 1, 2, 3}, 7},
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 8},
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 1}, 9},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 10},
	}
	for _, c := range cases {
		lc_0189_rotate_array(c.input, c.k)
		assert.Equal(t, c.output, c.input)
	}
}

func lc_0189_rotate_array(nums []int, k int) {
	shift := k % len(nums)
	if shift == 0 {
		return
	}
	for i, v := range append(nums[len(nums)-shift:], nums[:len(nums)-shift]...) {
		nums[i] = v
	}
}

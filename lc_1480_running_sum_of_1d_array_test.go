package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/running-sum-of-1d-array
func Test_lc_1480_running_sum_of_1d_array(t *testing.T) {
	cases := []struct {
		input, output []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 0, 0, 0, 0}, []int{1, 1, 1, 1, 1}},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_1480_running_sum_of_1d_array(c.input))
	}
}

// Решаем по месту. Итерируемся от второго элемента, прибавляя каждый раз предыдущий элемент, и так до конца массива.
func lc_1480_running_sum_of_1d_array(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0001_two_sum(t *testing.T) {
	assert.ElementsMatch(t, []int{1, 3}, lc_0001_two_sum_1([]int{1, 4, 5, 3, 7, 1}, 7))
	assert.ElementsMatch(t, []int{0, 1}, lc_0001_two_sum_1([]int{1, 4}, 5))
	assert.ElementsMatch(t, []int{5, 6}, lc_0001_two_sum_1([]int{9, 9, 9, 9, 9, 9, 1}, 10))

	assert.ElementsMatch(t, []int{1, 3}, lc_0001_two_sum_2([]int{1, 4, 5, 3, 7, 1}, 7))
	assert.ElementsMatch(t, []int{0, 1}, lc_0001_two_sum_2([]int{1, 4}, 5))
	assert.ElementsMatch(t, []int{0, 6}, lc_0001_two_sum_2([]int{9, 9, 9, 9, 9, 9, 1}, 10))
}

func lc_0001_two_sum_1(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}

func lc_0001_two_sum_2(nums []int, target int) []int {
	for i := range len(nums) - 1 {
		for j := i + 1; j < len(nums); j++ {
			switch {
			case nums[i]+nums[j] == target:
				return []int{i, j}
			}
		}
	}
	return nil
}

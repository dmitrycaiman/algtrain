package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0080_remove_dublicates(t *testing.T) {
	cases := []struct {
		input, output []int
		res           int
	}{
		{input: []int{1, 2, 3, 3, 4, 7, 7, 8, 8, 9, 9, 9, 9}, output: []int{1, 2, 3, 3, 4, 7, 7, 8, 8, 9, 9}, res: 11},
		{input: []int{1, 2, 3, 3, 3, 3, 3, 8, 8}, output: []int{1, 2, 3, 3, 8, 8}, res: 6},
		{input: []int{1, 1, 2, 2, 3, 4, 4, 4, 8, 9, 9}, output: []int{1, 1, 2, 2, 3, 4, 4, 8, 9, 9}, res: 10},
	}
	for _, v := range cases {
		assert.Equal(t, v.res, lc_0080_remove_dublicates(v.input))
		for i := range v.res {
			assert.Equal(t, v.output[i], v.input[i])
		}
	}
}

func lc_0080_remove_dublicates(nums []int) int {
	i, dub := 0, false
	for j := 1; j < len(nums); j++ {
		if (nums[j] > nums[i]) || (!dub && nums[j] == nums[i]) {
			nums[i+1] = nums[j]
			dub = nums[i] == nums[i+1]
			i++
		}
	}
	return i + 1
}

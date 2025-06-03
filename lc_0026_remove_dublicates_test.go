package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-array
func Test_lc_0026_remove_dublicates(t *testing.T) {
	cases := []struct {
		input, output []int
		res           int
	}{
		{input: []int{1, 2, 3, 3, 4, 7, 7, 8, 8, 9, 9, 9, 9}, output: []int{1, 2, 3, 4, 7, 8, 9}, res: 7},
		{input: []int{1, 2, 3, 3, 3, 3, 3, 8, 8}, output: []int{1, 2, 3, 8}, res: 4},
		{input: []int{1, 1, 2, 2, 3, 4, 4, 4, 8, 9, 9}, output: []int{1, 2, 3, 4, 8, 9}, res: 6},
	}
	for _, v := range cases {
		assert.Equal(t, v.res, lc_0026_remove_dublicates(v.input))
		for i := range v.res {
			assert.Equal(t, v.output[i], v.input[i])
		}
	}
}

func lc_0026_remove_dublicates(nums []int) int {
	// i + 1 — позиция, на которую мы ищем подходящего кандидата.
	// j — метка последовательного перебора всех значений.
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] > nums[i] {
			// Кандидат найден. Ищем кандидата на следующую позицию.
			if j-i != 1 {
				nums[i+1] = nums[j]
			}
			i++
		}
	}
	return i + 1
}

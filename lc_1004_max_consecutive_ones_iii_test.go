package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/max-consecutive-ones-iii
func Test_lc_1004_max_consecutive_ones_iii(t *testing.T) {
	cases := []struct {
		nums      []int
		k, result int
	}{
		{[]int{1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1}, 2, 8},
		{[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0}, 6, 14},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{[]int{0, 0, 0}, 1, 1},
		{[]int{1, 1, 0}, 0, 2},
		{[]int{0, 1, 0, 0, 1, 0, 0}, 0, 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_1004_max_consecutive_ones_iii(c.nums, c.k))
	}
}

// Плавающее окно.
// Ставим границы окна на начало массива. Расширяем окно, пока не кончатся "очки нулей", то есть замены нулей единицами.
// Как только очки кончились, сужаем окно по левой границе до следующего нуля. На каждом этапе храним длину окна.
// Возвращаем наибольшую длину окна. Итого получается, что в окне всегда находится k нулей.
func lc_1004_max_consecutive_ones_iii(nums []int, k int) int {
	best := 0
	for i, j, result, counter := -1, 0, 0, k; j < len(nums); j++ {
		if nums[j] == 1 {
			result++
		} else if counter != 0 {
			counter, result = counter-1, result+1
		} else {
			for i++; nums[i] != 0; i, result = i+1, result-1 {
			}
		}
		best = max(best, result)
	}
	return best
}

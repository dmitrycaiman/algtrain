package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-increasing-subsequence
func Test_lc_0300_longest_increasing_subsequence(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 1, 1, 1, 3, 1, 2, 5, 7, 4, 9}, 6},
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
		{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0300_longest_increasing_subsequence(c.input))
	}
}

// Для решения задачи нужно решить её для каждого промежутка массива с шагом 1 от начала и до конца.
// Каждое следующее решение есть либо продолжение предыдущего уже сделанного решения, либо новое решение, которое начинается с 1.
// Нужно вычислить все решения и выбрать наилучшее.
//
// Пояснение:
// Пусть решение на длине i есть d[i] (при этом d[0] = 1).
// Оно является продолжением решения d[j], если i > j (т.е. оно предыдущее) и nums[i] > nums[j]; тогда d[i] = d[j] + 1.
// Если для d[i] не нашлось такого d[j], то d[i] = 1 (последовательность из одного элемента).
// Таким образом, для каждого нового решения d[i] ищем наилучшее из предыдущих d[j], которое удовлетворяет условию i > j и nums[i] > nums[j].
func lc_0300_longest_increasing_subsequence(nums []int) int {
	d, max := make([]int, len(nums)), 1
	d[0] = 1
	for i := 1; i < len(nums); i++ {
		d[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && d[i] < d[j]+1 {
				d[i] = d[j] + 1
			}
		}
		if d[i] > max {
			max = d[i]
		}
	}
	return max
}

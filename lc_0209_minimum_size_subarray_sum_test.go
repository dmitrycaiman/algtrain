package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-size-subarray-sum
func Test_lc_0209_minimum_size_subarray_sum(t *testing.T) {
	cases := []struct {
		input          []int
		target, result int
	}{
		{[]int{1, 2, 3, 4, 2, 1, 5, 2, 3}, 5, 1},
		{[]int{1, 1, 1, 1, 1}, 3, 3},
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0209_minimum_size_subarray_sum(c.target, c.input))
	}
}

// Метод скользящего окна.
// Изначально границы окна указывают на первый элемент: длина 1, сумма равна первому элементу.
// На каждом шаге проверяем:
// - если попали в таргет, то запоминаем длину окна (если она наилучшая), уменьшаем окно слева (нет смысла увеличивать окно).
// - если нет, увеличиваем окно справа.
// На каждом шаге пересчитываем сумму. Возвращаем наилучшую длину. Если так и не попали в таргет, возвращаем ноль.
func lc_0209_minimum_size_subarray_sum(target int, nums []int) int {
	result, length, sum := 100_001, 1, nums[0]
	for i, j := 0, 0; ; {
		if sum >= target {
			if length < result {
				result = length
			}
			sum -= nums[i]
			i++
			length--
		} else if j < len(nums)-1 {
			j++
			sum += nums[j]
			length++
		} else {
			break
		}
	}
	if result == 100_001 {
		return 0
	}
	return result
}

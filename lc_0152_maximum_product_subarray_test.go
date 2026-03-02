package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-product-subarray
func Test_lc_0152_maximum_product_subarray(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 2, 0, 3, -1, 4, -1}, 12},
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, -3, -1, -1, 0, 10}, 10},
		{[]int{3, -1, 4}, 4},
		{[]int{0}, 0},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0152_maximum_product_subarray(c.input))
	}
}

// Динамическое программирование с элементами жадности.
// Итерируемся по массиву и запоминаем два значения: минимальное и максимальное локальные произведения.
// Начинаем, положив, что локальные максимум и минимум равны первому элементу.
// На каждой итерации у нас есть выбор по действию с новым элементом:
// - умножить на локальный минимум;
// - умножить на локальный максимум;
// - оставить только его, отбросив все предыдущие значения.
// Производим данные действия и далее обновляем локальные значения:
// - локальный минимум есть наименьшее из трёх полученных;
// - локальный масимум есть наибольшее.
// Обновляем наилучший результат, сравнив его с локальным максимумом.s
func lc_0152_maximum_product_subarray(nums []int) int {
	best, globalMin, globalMax := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		localMin, localMax := globalMin*nums[i], globalMax*nums[i]
		globalMin, globalMax = min(localMin, localMax, nums[i]), max(localMin, localMax, nums[i])
		best = max(best, globalMax)
	}
	return best
}

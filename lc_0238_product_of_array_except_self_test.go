package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/product-of-array-except-self
func Test_lc_0238_product_of_array_except_self(t *testing.T) {
	cases := []struct{ input, output []int }{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0238_product_of_array_except_self(c.input))
	}
}

// 0ms (100%), 10.84MB (23.21%)
// Создаём массив результатов и все присвоения делаем в нём. Итерируемся по входному массиву.
// Проходимся вперёд, по пути создаём префиксное произведение и присваиваем его текущему элементу. Таким образом покроем часть множителей, лежащих слева.
// Проходимся назад, но теперь создаём постфиксное произведение и умножаем его на текущий элемент. Так покрываем правые множители.
func lc_0238_product_of_array_except_self(nums []int) []int {
	res := make([]int, len(nums))
	for i, p := 0, 1; i < len(nums); i++ {
		res[i] = p
		p *= nums[i]
	}
	for i, p := len(nums)-1, 1; i >= 0; i-- {
		res[i] *= p
		p *= nums[i]
	}
	return res
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/missing-number
func Test_lc_0268_missing_number(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{[]int{0, 5, 2, 1, 4}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 0},
		{[]int{0}, 1},
		{[]int{1}, 0},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0268_missing_number(c.nums))
	}
}

// В массиве представлены числа, которые при сортировке будут идти подряд с потерей одного числа.
// Это по сути индексы массива, поэтому мы знаем, сколько будет всего чисел и какими они будут.
// Складываем все индексы массива (плюс длина как последний элемент) и вычитаем из этой суммы все его значения.
// Полученная разница между всеми возможными элементами и теми, которые представлены на вводе, и есть недостающее число.
func lc_0268_missing_number(nums []int) int {
	sum := len(nums)
	for i, v := range nums {
		sum = sum + i - v
	}
	return sum
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-duplicate-number
func Test_lc_0287_find_the_duplicate_number(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{3, 3, 3, 3, 3, 3, 3, 3}, 3},
		{[]int{1, 5, 4, 7, 8, 9, 2, 1, 1, 1, 1, 1}, 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0287_find_the_duplicate_number(c.nums))
	}
}

// Алгоритм Флойда (черепаха и заяц).
// Рассматриваем массив как связный список, где значение элемента есть номер следующего узла.
// Это справедливо, так как мы имеем N элементов со значениями [1,N] и массив с индексами [0,N].
// 1. Запускаем черепаху (скорость 1) и зайца (скорость 2) из начала массива и ждём, пока они встретятся.
// 2. После встречи запускаем ещё одну черепаху вместо зайца и опять ждём их встречи.
// 3. Возвращаем индекс элемента их встречи.
func lc_0287_find_the_duplicate_number(nums []int) int {
	var t, h int
	for t, h = nums[0], nums[nums[0]]; t != h; t, h = nums[t], nums[nums[h]] {
	}
	for t = 0; t != h; t, h = nums[t], nums[h] {
	}
	return t
}

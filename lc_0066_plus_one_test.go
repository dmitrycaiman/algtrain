package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/plus-one
func Test_lc_0066_plus_one(t *testing.T) {
	cases := []struct {
		digits, result []int
	}{
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{9, 9, 8}, []int{9, 9, 9}},
		{[]int{9}, []int{1, 0}},
		{[]int{1}, []int{2}},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0066_plus_one(c.digits))
	}
}

// Идём с конца и прибавляем единичку. Если перешагнули разряд, то продолжаем. Если нет, выходим.
// Учитываем, что на самом большом разряде тоже может быть переполнение.
func lc_0066_plus_one(digits []int) []int {
	for i := len(digits) - 1; ; i-- {
		if i == -1 {
			return append([]int{1}, digits...)
		}
		if digits[i] += 1; digits[i] == 10 {
			digits[i] = 0
		} else {
			break
		}
	}
	return digits
}

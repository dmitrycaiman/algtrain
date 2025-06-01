package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/single-number
func Test_lc_0136_single_number(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{[]int{-100, -100, 200, 200, 3}, 3},
		{[]int{200, 200, -300}, -300},
		{[]int{1, 2, 200, 2, 1}, 200},
		{[]int{-100, -1000, -1000}, -100},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0136_single_number(c.nums))
	}
}

// Свойство операции XOR.
// На исходном числе применили XOR последовательно несколькими любыми числами.
// Если применить XOR повторно теми же числами в любом порядке, то получим исходное число.
func lc_0136_single_number(nums []int) int {
	m := 0
	for _, v := range nums {
		m ^= v
	}
	return m
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/add-digits
func Test_lc_0258_add_digits(t *testing.T) {
	cases := []struct{ input, output int }{{38, 2}, {435, 3}}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0258_add_digits(c.input))
	}
}

// Берём число, сохраняем его остаток от деления на 10 и целочисленно делим его на 10. И так пока число не станет нулём.
// Затем берём сумму остатков, и если она 10 и более (т.е. содержит больше одной цифры), то всё повторяем.
func lc_0258_add_digits(num int) int {
	for {
		result := 0
		for num != 0 {
			result += num % 10
			num /= 10
		}
		if result < 10 {
			return result
		}
		num = result
	}
}

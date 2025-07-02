package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-1-bits
func Test_lc_0191_number_of_1_bits(t *testing.T) {
	cases := []struct{ n, res int }{{11, 3}, {128, 1}, {2147483645, 30}}
	for _, c := range cases {
		assert.Equal(t, c.res, lc_0191_number_of_1_bits(c.n))
	}
}

// Логическое И с 1 и сдвиг числа вправо на 1 до тех пор, пока число не станет 0.
// Каждое И, дающее 1, увеличивает счётчик.
func lc_0191_number_of_1_bits(n int) (res int) {
	for ; n != 0; n >>= 1 {
		res += n & 1
	}
	return
}

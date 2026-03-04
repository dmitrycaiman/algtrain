package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/power-of-four
func Test_lc_0342_power_of_four(t *testing.T) {
	cases := []struct {
		n      int
		result bool
	}{
		{1, true}, {4, true}, {16, true}, {64, true}, {0, false}, {2, false}, {5, false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0342_power_of_four(c.n))
	}
}

// Признаки степени 4-ки:
// 1. Не ноль.
// 2. В числе возведён только один бит.
// 3. Возведённый бит является нечётным.
func lc_0342_power_of_four(n int) bool {
	f := 0b1010101010101010101010101010101
	return f|n == f && n&(n-1) == 0 && n != 0
}

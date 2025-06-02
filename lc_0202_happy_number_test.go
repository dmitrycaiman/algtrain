package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/happy-number
func Test_lc_0202_happy_number(t *testing.T) {
	cases := []struct {
		n      int
		result bool
	}{
		{19, true}, {565, true}, {564, false}, {2, false}, {999, false}, {997, false}, {998, true},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0202_happy_number(c.n))
	}
}

// Рассчитываем суммы квадратов цифр и складываем их в мапу.
// Как только встречаем повтор, то выходим с неудачей. Встречаем единицу, выходим с победой.
func lc_0202_happy_number(n int) bool {
	sumOfDigitsSquares := func(m int) int {
		div, result := 10, 0
		for m > 0 {
			rest := m % 10
			m /= div
			result += rest * rest
		}
		return result
	}
	m := map[int]struct{}{}
	for {
		if n == 1 {
			return true
		}
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}
		n = sumOfDigitsSquares(n)
	}
}

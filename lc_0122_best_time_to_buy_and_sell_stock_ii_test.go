package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0122_best_time_to_buy_and_sell_stock_ii(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 4}, 3},
		{[]int{4, 3, 2, 1}, 0},
		{[]int{4, 3, 2, 10}, 8},
		{[]int{1, 2, 10, 1}, 9},
		{[]int{1}, 0},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0122_best_time_to_buy_and_sell_stock_ii(c.input))

	}
}

// Представляем входной массив как график и ищем в нём перегибы: на низинах покупаем, на пиках продаём.
// Изначально считаем, что идём вверх и покупаем первый элемент: так как можно продавать в день покупки, то при движении вниз
// мы продадим его на первой итерации. В конце цикла, если мы шли вверх (а значит мы когда-то купили "внизу"), то нужно продать.
func lc_0122_best_time_to_buy_and_sell_stock_ii(prices []int) int {
	res, cur, up := 0, prices[0], true
	for i := 0; i < len(prices)-1; i++ {
		switch {
		case prices[i] > prices[i+1] && up:
			res += prices[i] - cur
			up = false
		case prices[i] < prices[i+1] && !up:
			cur = prices[i]
			up = true
		}
	}
	if up {
		res += prices[len(prices)-1] - cur
	}
	return res
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii
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
		assert.Equal(t, c.output, lc_0122_best_time_to_buy_and_sell_stock_ii_1(c.input))
		assert.Equal(t, c.output, lc_0122_best_time_to_buy_and_sell_stock_ii_2(c.input))

	}
}

// Жадная задача. Идём по массиву, покупаем в начале каждого промежутка возрастания и продаём сразу в его конце.
func lc_0122_best_time_to_buy_and_sell_stock_ii_1(prices []int) int {
	result := 0
	for i, j := 0, 1; j < len(prices); {
		if prices[j] < prices[j-1] {
			result += prices[j-1] - prices[i]
			i = j
		} else if j == len(prices)-1 {
			result += prices[j] - prices[i]
		}
		j++
	}
	return result
}

// Представляем входной массив как график и ищем в нём перегибы: на низинах покупаем, на пиках продаём.
// Изначально считаем, что идём вверх и покупаем первый элемент: так как можно продавать в день покупки, то при движении вниз
// мы продадим его на первой итерации. В конце цикла, если мы шли вверх (а значит мы когда-то купили "внизу"), то нужно продать.
func lc_0122_best_time_to_buy_and_sell_stock_ii_2(prices []int) int {
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

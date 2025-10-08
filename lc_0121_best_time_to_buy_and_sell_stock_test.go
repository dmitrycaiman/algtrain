package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock
func Test_lc_0121_best_time_to_buy_and_sell_stock(t *testing.T) {
	assert.Equal(t, 5, lc_0121_best_time_to_buy_and_sell_stock([]int{1, 6, 5, 0, 3, 2}))
	assert.Equal(t, 0, lc_0121_best_time_to_buy_and_sell_stock([]int{8, 6, 5, 4, 3, 2}))
	assert.Equal(t, 3, lc_0121_best_time_to_buy_and_sell_stock([]int{8, 9, 10, 1, 3, 2, 4}))
	assert.Equal(t, 7, lc_0121_best_time_to_buy_and_sell_stock([]int{0, 0, 1, 2, 3, 7}))

}

// Жадная задача. Итерируемся по массиву и либо запоминаем и обновляем наименьшее число,
// либо вычисляем разницу между элементом и текущим наименьшим числом.
// Возвращаем наилучший результат.
func lc_0121_best_time_to_buy_and_sell_stock(prices []int) int {
	min, result := prices[0], 0
	for _, v := range prices {
		if v < min {
			min = v
		} else if diff := v - min; diff > result {
			result = diff
		}
	}
	return result
}

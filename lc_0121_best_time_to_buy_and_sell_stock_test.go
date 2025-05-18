package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0121_best_time_to_buy_and_sell_stock(t *testing.T) {
	assert.Equal(t, 5, lc_0121_best_time_to_buy_and_sell_stock([]int{1, 6, 5, 0, 3, 2}))
	assert.Equal(t, 0, lc_0121_best_time_to_buy_and_sell_stock([]int{8, 6, 5, 4, 3, 2}))
	assert.Equal(t, 3, lc_0121_best_time_to_buy_and_sell_stock([]int{8, 9, 10, 1, 3, 2, 4}))
	assert.Equal(t, 7, lc_0121_best_time_to_buy_and_sell_stock([]int{0, 0, 1, 2, 3, 7}))

}

func lc_0121_best_time_to_buy_and_sell_stock(prices []int) int {
	min, res := prices[0], 0
	for i := 1; i <= len(prices)-1; i++ {
		if delta := prices[i] - min; delta > res {
			res = delta
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return res
}

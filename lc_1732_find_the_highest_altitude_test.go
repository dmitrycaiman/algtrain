package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-highest-altitude
func Test_lc_1732_find_the_highest_altitude(t *testing.T) {
	cases := []struct {
		gain   []int
		result int
	}{
		{[]int{-5, 1, 5, 0, -7}, 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_1732_find_the_highest_altitude(c.gain))
	}
}

// Последовательно определяем высоты точек, начиная с 0. Запоминаем наибольшую.
func lc_1732_find_the_highest_altitude(gain []int) int {
	result, lastPoint := 0, 0
	for _, v := range gain {
		lastPoint += v
		if lastPoint > result {
			result = lastPoint
		}
	}
	return result
}

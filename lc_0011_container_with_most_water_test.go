package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/container-with-most-water
func Test_lc_0011_container_with_most_water(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 3, 3, 3, 3, 1}, 9},
		{[]int{1, 2, 1, 3, 1, 4}, 8},
		{[]int{1, 2, 1, 5, 1, 5}, 10},
		{[]int{1, 9, 1}, 2},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0011_container_with_most_water(c.input))
	}
}

// Ставим метки на концах массива. Считаем и запоминаем объём между метками.
// Единственный способ увеличить объём — попробовать сдвинуть метку, указывающую на меньший элемент, вглубь массива.
// После каждой пробы так же считаем объём и запоминаем его при необходимости.
// Итерируемся так до встречи меток и возвращаем максимально полученный объём.
func lc_0011_container_with_most_water(height []int) int {
	maxVol := 0
	for i, j := 0, len(height)-1; i < j; {
		if vol := (j - i) * min(height[i], height[j]); vol > maxVol {
			maxVol = vol
		}
		switch {
		case height[i] < height[j]:
			i++
		case height[i] > height[j]:
			j--
		default:
			i++
			j--
		}
	}
	return maxVol
}

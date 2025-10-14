package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// #yandex
// https://leetcode.com/problems/maximize-distance-to-closest-person
func Test_lc_0849_maximize_distance_to_closest_person(t *testing.T) {
	cases := []struct {
		seats  []int
		result int
	}{
		{[]int{0, 0, 1, 1, 1, 0, 1}, 2},
		{[]int{0, 0, 1, 1, 0, 0, 0}, 3},
		{[]int{1, 0, 0, 0, 0, 0, 0}, 6},
		{[]int{1, 0}, 1},
		{[]int{1, 0, 1}, 1},
		{[]int{1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1}, 2},
		{[]int{1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1}, 3},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0849_maximize_distance_to_closest_person(c.seats))
	}
}

// 0/100, 8/88
// Два указателя. Первый указывает на начало промежутка с нулями, второй итерируется по массиву.
// Ищем промежутки нулей и считаем их длину: если промежуток в начале или конце массива, оставляем длину без изменений,
// а иначе опираемся на её чётность (от чётной берём половину, от нечётной половину с единицей).
// Возвращаем самую большую найденную длину.
func lc_0849_maximize_distance_to_closest_person(seats []int) int {
	maxLength, onZeroes := -1, false
	for i, j := 0, 0; j <= len(seats); j++ {
		switch {
		case j == len(seats): // Обработка особого случая, когда промежуток с нулями так и не кончился на правой границе массива.
			if !onZeroes {
				break
			}
			fallthrough
		case onZeroes && seats[j] == 1:
			l := j - i
			switch {
			case i == 0 || j == len(seats):
			case l%2 == 0:
				l = l / 2
			case l%2 != 0:
				l = l/2 + 1
			}
			if l > maxLength {
				maxLength = l
			}
			onZeroes = false
		case !onZeroes && seats[j] == 0:
			onZeroes = true
			i = j
		}

	}
	return maxLength
}

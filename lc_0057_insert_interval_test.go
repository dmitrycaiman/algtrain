package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/insert-interval
func Test_lc_0057_insert_interval(t *testing.T) {
	cases := []struct {
		intervals   [][]int
		newInterval []int
		result      [][]int
	}{
		{[][]int{{1, 2}, {5, 6}}, []int{3, 4}, [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{[][]int{{1, 2}, {5, 6}}, []int{0, 1}, [][]int{{0, 2}, {5, 6}}},
		{[][]int{{1, 2}, {5, 6}}, []int{6, 7}, [][]int{{1, 2}, {5, 7}}},
		{[][]int{{1, 2}, {5, 6}}, []int{2, 5}, [][]int{{1, 6}}},
		{[][]int{{1, 2}, {5, 6}}, []int{7, 8}, [][]int{{1, 2}, {5, 6}, {7, 8}}},
		{[][]int{{5, 6}}, []int{3, 4}, [][]int{{3, 4}, {5, 6}}},
		{[][]int{{5, 6}}, []int{4, 5}, [][]int{{4, 6}}},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0057_insert_interval(c.intervals, c.newInterval))
	}
}

// Итерируемся по массиву. Пытаемся определить, куда вставить сначала левую границу, потом правую.
// Левая граница встаёт либо слева от текущего интервала, либо приравнивается к его левой границе, либо продолжаем.
// Правая граница встаёт либо слева от текущего интервала, либо приравнивается к его правой границе, либо продолжаем.
// Пока ни одна граница не определена, вставляем интервалы "как есть".
// Как только обе границы будут определены, вставляем их как интервал и копируем остальную часть массива (если требуется).
// Обрабатываем краевые случаи, когда мы не смогли вставить ни одной границы, либо когда вставили только левую.
func lc_0057_insert_interval(input [][]int, n []int) [][]int {
	output, a, b := [][]int{}, -1, -1
	for i, v := range input {
		if a == -1 {
			switch {
			case n[0] < v[0]:
				a = n[0]
			case n[0] <= v[1]:
				a = v[0]
			}
		}
		if b == -1 {
			switch {
			case n[1] < v[0]:
				b = n[1]
			case n[1] <= v[1]:
				b = v[1]
			}
		}
		if a != -1 && b != -1 {
			output = append(output, []int{a, b})
			if b == v[1] {
				output = append(output, input[i+1:]...)
			} else {
				output = append(output, input[i:]...)
			}
			break
		}
		if a == -1 && b == -1 {
			output = append(output, v)
		}
	}
	switch {
	case a == -1 && b == -1:
		output = append(output, n)
	case b == -1:
		output = append(output, []int{a, n[1]})
	}
	return output
}

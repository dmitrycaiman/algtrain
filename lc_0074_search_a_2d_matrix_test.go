package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0074_search_a_2d_matrix(t *testing.T) {
	cases := []struct {
		matrix [][]int
		target int
		result bool
	}{
		{[][]int{{1, 2, 4}, {5, 7, 8}, {9, 9, 9}}, 4, true},
		{[][]int{{1, 2, 4}, {5, 7, 8}, {9, 9, 9}}, 7, true},
		{[][]int{{1, 2, 4}, {5, 7, 8}, {9, 9, 9}}, 9, true},
		{[][]int{{1, 2, 4}, {5, 7, 8}, {9, 9, 9}}, 4, true},
		{[][]int{{1, 2, 4}, {5, 7, 8}, {9, 9, 9}}, 3, false},
		{[][]int{{1, 2, 4}, {5, 7, 8}, {9, 9, 9}}, 6, false},
		{[][]int{{1, 2, 4}, {5, 7, 8}, {9, 9, 9}}, 0, false},
		{[][]int{{1, 2, 4}, {5, 7, 8}, {9, 9, 9}}, 10, false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0074_search_a_2d_matrix_1(c.matrix, c.target))
		assert.Equal(t, c.result, lc_0074_search_a_2d_matrix_2(c.matrix, c.target))
	}
}

// Выбираем ряд для поиска: если первый элемент ряда больше искомого, то нам подходит предыдущий ряд.
// Далее осуществляем бинарный поиск в выбранном ряду. Если в процессе границы сошлись на одном элементе,
// то сравниваем этот элемент с искомым. Если разошлись, то искомого элемента не существует.
func lc_0074_search_a_2d_matrix_1(matrix [][]int, target int) bool {
	row := len(matrix) - 1
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] > target {
			row = i - 1
			break
		}
	}
	s := matrix[row]
	a, b := 0, len(s)-1
	for {
		switch {
		case b == a:
			return s[a] == target
		case b < a:
			return false
		}
		m := a + (b-a)/2
		switch {
		case s[m] > target:
			b = m - 1
		case s[m] < target:
			a = m + 1
		case s[m] == target:
			return true
		}
	}
}

func lc_0074_search_a_2d_matrix_2(matrix [][]int, target int) bool {
	n := -2
	for i, v := range matrix {
		if v[0] > target {
			n = i - 1
			break
		}
	}
	switch n {
	case -1:
		return false
	case -2:
		n = len(matrix) - 1
		if target > matrix[n][len(matrix[n])-1] {
			return false
		}
	}
	m := matrix[n]
	for {
		if len(m) == 0 {
			return false
		}
		mid := len(m) / 2
		val := m[mid]
		switch {
		case val > target:
			m = m[:mid]
		case val < target:
			m = m[mid+1:]
		default:
			return true
		}
	}
}

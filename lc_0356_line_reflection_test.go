package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/line-reflection
func Test_lc_0356_line_reflection(t *testing.T) {
	cases := []struct {
		points [][]int
		result bool
	}{
		{[][]int{{1, 1}, {-1, 1}}, true},
		{[][]int{{1, 1}, {-1, -1}}, false},
		{[][]int{{2, 1}, {2, 2}, {2, 3}}, true},
		{[][]int{{4, 1}, {2, 1}, {2, 2}, {2, 3}, {0, 1}}, true},
		{[][]int{{4, 2}, {2, 1}, {2, 2}, {2, 3}, {0, 1}}, false},
		{[][]int{{1, 1}, {2, 2}, {2, 3}}, false},
		{[][]int{{2, 3}, {4, 4}, {6, 4}, {6, 1}, {10, 4}, {12, 4}, {10, 1}, {14, 3}}, true},
		{[][]int{{2, 3}, {4, 4}, {6, 4}, {6, 1}, {10, 4}, {12, 4}, {10, 1}, {14, 3}, {1, 1}, {-1, 1}}, false},
		{[][]int{{2, 3}, {4, 4}, {6, 4}, {6, 1}, {10, 4}, {12, 4}, {10, 1}, {14, 3}, {14, 3}}, false},
		{[][]int{{2, 3}, {4, 4}, {6, 4}, {6, 1}, {10, 4}, {12, 4}, {10, 1}, {13, 3}}, false},
		{[][]int{{2, 3}, {4, 4}, {6, 4}, {6, 1}, {10, 4}, {12, 4}, {8, 1}, {8, 1}, {14, 3}, {10, 1}}, true},
		{[][]int{{1, 2}, {1, 2}, {1, 2}, {1, 2}}, true},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0356_line_reflection(c.points))
	}
}

// Проходим по массиву точек и складываем их в хеш-таблицу, где ключ это координаты, значение это количество вхождений.
// Заодно высчитываем максимум и минимум по оси X: это даст нам возможность вычислить положение средней линии.
// Далее проходимся по получившейся мапе и ищем зеркальные точки: если не находим, сразу выходим с неудачей.
func lc_0356_line_reflection(points [][]int) bool {
	m, min, max := map[[2]int]int{}, points[0][0], points[0][0]
	for _, point := range points {
		m[[2]int(point)]++
		switch {
		case point[0] > max:
			max = point[0]
		case point[0] < min:
			min = point[0]
		}
	}
	mid := min + (max-min)/2
	var mirrorPoint [2]int
	for point, n := range m {
		if point[0] >= mid {
			mirrorPoint = [2]int{point[0] - 2*(point[0]-mid), point[1]}
		} else {
			mirrorPoint = [2]int{point[0] + 2*(mid-point[0]), point[1]}
		}
		if m[mirrorPoint] != n {
			return false
		}
		delete(m, point)
		delete(m, mirrorPoint)
	}
	return true
}

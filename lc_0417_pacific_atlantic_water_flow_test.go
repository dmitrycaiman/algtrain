package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/pacific-atlantic-water-flow
func Test_lc_0417_pacific_atlantic_water_flow(t *testing.T) {
	cases := []struct {
		grid   [][]int
		result [][]int
	}{
		{
			[][]int{
				{2, 1, 3},
				{3, 2, 3},
				{4, 1, 4},
			},
			[][]int{{0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 2}},
		},
		{
			[][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			[][]int{
				{10, 10, 10},
				{10, 1, 10},
				{10, 10, 10},
			},
			[][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
		{
			[][]int{
				{5, 5, 5, 5},
				{4, 4, 4, 4},
				{5, 5, 5, 5},
			},
			[][]int{{0, 3}, {0, 2}, {0, 1}, {0, 0}, {1, 3}, {2, 3}, {2, 2}, {2, 1}, {2, 0}, {1, 2}, {1, 1}, {1, 0}},
		},
	}
	for _, c := range cases {
		assert.ElementsMatch(t, c.result, lc_0417_pacific_atlantic_water_flow(c.grid))
	}
}

// 100/83
// Левую и верхнюю границы красим одним цветом. Из каждого закрашеннго элемента рекурсивно пытаемся закрасить все соседние.
// То же самое, но другим цветом, проделываем и с границами справа и снизу.
// Возвращаем те точки, которые попали в пересечение закрашиваемых областей.
func lc_0417_pacific_atlantic_water_flow(heights [][]int) [][]int {
	result, reachability := [][]int{}, make([][][3]bool, len(heights))
	for i := range reachability {
		reachability[i] = make([][3]bool, len(heights[0]))
	}
	for i := 0; i < len(heights); i++ {
		result = lc_0417_fill(i, 0, true, heights, reachability, result)
	}
	for i := 0; i < len(heights[0]); i++ {
		result = lc_0417_fill(0, i, true, heights, reachability, result)
	}
	for i := 0; i < len(heights); i++ {
		result = lc_0417_fill(i, len(heights[0])-1, false, heights, reachability, result)
	}
	for i := 0; i < len(heights[0]); i++ {
		result = lc_0417_fill(len(heights)-1, i, false, heights, reachability, result)
	}
	return result
}

func lc_0417_fill(i, j int, pacific bool, heights [][]int, reachability [][][3]bool, result [][]int) [][]int {
	n := 0
	if !pacific {
		n = 1
	}
	reachability[i][j][n] = true
	if reachability[i][j][0] && reachability[i][j][1] && !reachability[i][j][2] {
		result = append(result, []int{i, j})
		reachability[i][j][2] = true
	}
	if i != 0 && !reachability[i-1][j][n] && heights[i-1][j] >= heights[i][j] {
		result = lc_0417_fill(i-1, j, pacific, heights, reachability, result)
	}
	if i != len(heights)-1 && !reachability[i+1][j][n] && heights[i+1][j] >= heights[i][j] {
		result = lc_0417_fill(i+1, j, pacific, heights, reachability, result)
	}
	if j != 0 && !reachability[i][j-1][n] && heights[i][j-1] >= heights[i][j] {
		result = lc_0417_fill(i, j-1, pacific, heights, reachability, result)
	}
	if j != len(heights[0])-1 && !reachability[i][j+1][n] && heights[i][j+1] >= heights[i][j] {
		result = lc_0417_fill(i, j+1, pacific, heights, reachability, result)
	}
	return result
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-islands
func Test_lc_0200_number_of_islands(t *testing.T) {
	cases := []struct {
		grid   [][]byte
		result int
	}{
		{
			[][]byte{
				{'1', '1', '0'},
				{'0', '0', '0'},
				{'1', '0', '1'},
			},
			3,
		},
		{
			[][]byte{
				{'1', '1', '0'},
				{'1', '0', '0'},
				{'1', '1', '1'},
			},
			1,
		},
		{
			[][]byte{
				{'0', '1', '0'},
				{'1', '0', '1'},
				{'0', '1', '0'},
			},
			4,
		},
		{
			[][]byte{
				{'0', '1', '1'},
				{'1', '0', '1'},
				{'1', '1', '0'},
			},
			2,
		},
		{[][]byte{{'0'}}, 0},
		{[][]byte{{'1'}}, 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0200_number_of_islands(c.grid))
	}
}

// Создаём рекурсивную функцию цепной закраски. Она валидирует элемент матрицы по индексам и закрашивает его,
// если он является незанятой землёй. Далее она пытается рекурсивно вызвать функцию закраски для своих соседей
// по горизонтали и вертикали. Таким образом, вызвав функцию закраски для одного элемента, мы закрасим сразу весь остров.
// Итого: итерируемся по матрице и при нахождении незанятой земли вызываем закраску и увеличиваем счётчик островов.
func lc_0200_number_of_islands(grid [][]byte) int {
	islands := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				islands++
				fill(i, j, grid)
			}
		}
	}
	return islands
}

func fill(i, j int, grid [][]byte) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	fill(i+1, j, grid)
	fill(i-1, j, grid)
	fill(i, j+1, grid)
	fill(i, j-1, grid)
}

package main

import (
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Route256
// «Электронная таблица»
// Задана прямоугольная таблица NxM из чисел. Требуется обработать k запросов: примените стабильную сортировку по неубыванию к столбцу i.
// Относительное расположение строк при одинаковых значениях в столбцах должно быть сохранено.
func Test_r256_spreadsheet(t *testing.T) {
	testFlow := func(f func(table [][]int, clicks []int)) {
		cases := []struct {
			initTable, sortedTable [][]int
			clicks                 []int
		}{
			{
				[][]int{
					{1, 3},
					{2, 2},
					{3, 1},
				},
				[][]int{
					{3, 1},
					{2, 2},
					{1, 3},
				},
				[]int{1, 1, 1, 1},
			},
			{
				[][]int{
					{1, 3},
					{2, 2},
					{3, 1},
				},
				[][]int{
					{1, 3},
					{2, 2},
					{3, 1},
				},
				[]int{1, 1, 0, 0},
			},
			{
				[][]int{
					{3, 4, 1},
					{2, 2, 5},
					{2, 4, 2},
					{2, 2, 1},
				},
				[][]int{
					{2, 2, 1},
					{3, 4, 1},
					{2, 4, 2},
					{2, 2, 5},
				},
				[]int{1, 0, 2},
			},
			{
				[][]int{
					{100},
					{9},
					{10},
				},
				[][]int{
					{9},
					{10},
					{100},
				},
				[]int{0, 0},
			},
			{
				[][]int{
					{2, 11, 72},
					{99, 11, 13},
					{2, 8, 13},
				},
				[][]int{
					{2, 8, 13},
					{2, 11, 72},
					{99, 11, 13},
				},
				[]int{1, 2, 1, 0, 1},
			},
		}
		for _, c := range cases {
			f(c.initTable, c.clicks)
			assert.Equal(t, c.sortedTable, c.initTable)
		}
	}
	testFlow(r256_spreadsheet_1)
	testFlow(r256_spreadsheet_2)
	testFlow(r256_spreadsheet_3)
}

// Решение через сортировку пузырьком.
func r256_spreadsheet_1(table [][]int, clicks []int) {
	for _, col := range clicks {
		for {
			swap := false
			for i := 0; i < len(table)-1; i++ {
				if table[i][col] > table[i+1][col] {
					table[i], table[i+1] = table[i+1], table[i]
					swap = true
				}
			}
			if !swap {
				break
			}
		}
	}
}

// Решение встроенными функциями.
func r256_spreadsheet_2(table [][]int, clicks []int) {
	for _, col := range clicks {
		sort.SliceStable(
			table,
			func(i, j int) bool {
				return table[i][col] < table[j][col]
			},
		)
	}
}

// Решение встроенными функциями.
func r256_spreadsheet_3(table [][]int, clicks []int) {
	for _, col := range clicks {
		slices.SortStableFunc(
			table,
			func(a, b []int) int {
				switch {
				case a[col] > b[col]:
					return 1
				case a[col] < b[col]:
					return -1
				default:
					return 0
				}
			},
		)
	}
}

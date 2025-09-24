package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0056_merge_intervals(t *testing.T) {
	cases := []struct {
		input, output [][]int
	}{
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {1, 6}}, [][]int{{1, 6}}},
		{[][]int{{1, 4}, {0, 0}}, [][]int{{1, 4}, {0, 0}}},
		{[][]int{{1, 4}, {1, 2}, {-5, 2}, {-10, 10}, {1, 2}, {3, 3}, {4, 4}}, [][]int{{-10, 10}}},
	}
	for _, c := range cases {
		assert.ElementsMatch(t, c.output, lc_0056_merge_intervals(c.input))
		assert.ElementsMatch(t, c.output, lc_0056_merge_intervals_bf(c.input))
	}
}

// Делаем сортировку входных данных по первому элементу каждого интервала.
// Далее проходимся по сортированному массиву и пробуем слить интервал со следующим, если они пересекаются.
// Если интервал нельзя слить со следующим, мы добавляем его в массив результатов.
// Также в самом начале добавляем в массив результатов последний интервал,
// так как у него не будет "шанса" на попытку слияния со следующим интервалом.
func lc_0056_merge_intervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	result := [][]int{intervals[len(intervals)-1]}
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1][0] <= intervals[i][1] {
			intervals[i+1][0] = intervals[i][0]
			if intervals[i+1][1] < intervals[i][1] {
				intervals[i+1][1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}

// Решение полным перебором. Квадратичная сложность по производительности.
// Идём по массиву и пытаемся слить интервал с каким-либо впереди. Если получилось, переходим к следующему.
// Если не получилось, добавляем интервал в массив результатов.
func lc_0056_merge_intervals_bf(intervals [][]int) [][]int {
	result := [][]int{}
loop:
	for i := range intervals {
		for j := i + 1; j < len(intervals); j++ {
			switch {
			case intervals[i][0] >= intervals[j][0] && intervals[i][0] <= intervals[j][1]:
				if intervals[i][1] > intervals[j][1] {
					intervals[j][1] = intervals[i][1]
				}
			case intervals[i][1] >= intervals[j][0] && intervals[i][1] <= intervals[j][1]:
				if intervals[i][0] < intervals[j][0] {
					intervals[j][0] = intervals[i][0]
				}
			case intervals[i][0] < intervals[j][0] && intervals[i][1] > intervals[j][1]:
				intervals[j][0], intervals[j][1] = intervals[i][0], intervals[i][1]
			default:
				continue
			}
			continue loop
		}
		result = append(result, intervals[i])
	}
	return result
}

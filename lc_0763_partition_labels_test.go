package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/partition-labels/
func Test_lc_0763_partition_labels(t *testing.T) {
	cases := []struct {
		input  string
		output []int
	}{
		{"ababccddefef", []int{4, 2, 2, 4}},
		{"ababccddefefb", []int{13}},
		{"aaaaaaaab", []int{8, 1}},
		{"a", []int{1}},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0763_partition_labels(c.input))
	}
}

// Формируем интервалы вхождения для каждого симовла.
// Собираем интервалы так, чтобы они были отсортированы в порядке возрастания левой	границы.
// Далее итерируемся по массиву интервалов и сливаем те соседние, которые пересекаются.
// В результат выводим длины получившихся после слияния интервалов.
func lc_0763_partition_labels(s string) []int {
	symbols, intervals := map[byte]int{}, [][2]int{}
	for i := range s {
		number, ok := symbols[s[i]]
		if !ok {
			intervals = append(intervals, [2]int{i, i})
			symbols[s[i]] = len(intervals) - 1
			continue
		}
		intervals[number][1] = i
	}
	results := []int{}
	l, r := 0, 0
	for i := range intervals {
		if intervals[i][0] > r {
			results = append(results, r-l+1)
			l, r = intervals[i][0], intervals[i][1]
		} else if intervals[i][1] > r {
			r = intervals[i][1]
		}
	}
	results = append(results, r-l+1)
	return results
}

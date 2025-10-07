package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sliding-window-median
func Test_lc_0480_sliding_window_median(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		result []float64
	}{
		{[]int{1, 2, 3, 4, 2, 3, 1, 4, 2}, 3, []float64{2, 3, 3, 3, 2, 3, 2}},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []float64{1, -1, -1, 3, 5, 6}},
		{[]int{1}, 1, []float64{1}},
		{[]int{1, 2}, 1, []float64{1, 2}},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0480_sliding_window_median(c.nums, c.k))
	}
}

// Создаём отдельный массив под окно. Итерируемся по входным данным и поддерживаем в массиве только необходимые значения,
// то есть удаляем те, которые вышли из окна, и добавляем те, которые вошли. Самое главное — это держать массив в отсортированном состоянии
// при добавлении элементов: тогда мы сможем сразу считывать медиану из его середины.
func lc_0480_sliding_window_median(nums []int, k int) []float64 {
	window, result := []int{}, []float64{}
	var saveMedian func()
	if k%2 == 0 {
		saveMedian = func() { result = append(result, (float64(window[k/2])+float64(window[(k/2)-1]))/2) }
	} else {
		saveMedian = func() { result = append(result, float64(window[k/2])) }
	}
	for i := range nums {
		index := sort.SearchInts(window, nums[i])
		window = append(window, 0)
		copy(window[index+1:], window[index:])
		window[index] = nums[i]
		if i >= k {
			index := sort.SearchInts(window, nums[i-k])
			window = append(window[:index], window[index+1:]...)
		}
		if i >= k-1 {
			saveMedian()
		}
	}
	return result
}

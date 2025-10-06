package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sliding-window-maximum
func Test_lc_0239_sliding_window_maximum(t *testing.T) {
	cases := []struct {
		nums, result []int
		k            int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, []int{3, 3, 5, 5, 6, 7}, 3},
		{[]int{1}, []int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5}, 2},
		{[]int{3, 4, 5}, []int{5}, 3},
		{[]int{1, -1}, []int{1, -1}, 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0239_sliding_window_maximum_1(c.nums, c.k))
		assert.Equal(t, c.result, lc_0239_sliding_window_maximum_2(c.nums, c.k))
	}
}

// Используем дек (двусторонняя очередь).
// Итерируемся по массиву. Вставляем текущий элемент в конец очереди. При этом, если перед ним есть меньшие элементы, удаляем их.
// Как только при итерациях достигнем размера окна, добавляем к результату первый элемент очереди.
// Также удаляем первый элемент, если левая граница окна с него сдвинулась.
func lc_0239_sliding_window_maximum_1(nums []int, k int) []int {
	result, queue := []int{}, []int{}
	for i := range nums {
		for len(queue) != 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		if i >= k-1 {
			result = append(result, queue[0])
			if nums[i-k+1] == queue[0] {
				queue = queue[1:]
			}
		}
	}
	return result
}

// Двигаем окно по массиву. Определяем локальный максимум на первой итерации и храним количество его появлений.
// При движении окна, если встречаем элемент больше максимума, запоминаем его и перезапускаем счётчик.
// Если ушедший из окна элемент был равен максимуму, уменьшаем счётчик.
// Если счётчик стал нулём, то производим новый поиск максимума в рамках окна.
// Сохраняем локальный максимум на каждой итерации.
func lc_0239_sliding_window_maximum_2(nums []int, k int) []int {
	max, counter, result := 10001, 1, []int{}
	for i := k - 1; i < len(nums); i++ {
		switch {
		case nums[i] == max:
			counter++
		case nums[i] > max:
			max = nums[i]
			counter = 1
		case i == k-1 || nums[i-k] == max:
			if counter == 1 {
				max = nums[i]
				for j := i - k + 1; j < i; j++ {
					switch {
					case nums[j] == max:
						counter++
					case nums[j] > max:
						max = nums[j]
						counter = 1
					}
				}
			} else {
				counter--
			}
		}
		result = append(result, max)
	}
	return result
}

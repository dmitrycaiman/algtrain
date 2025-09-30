package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_1493_longest_subarray_of_ones_after_deleting_one_element(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{[]int{1, 1, 1, 1, 1, 1}, 5},
		{[]int{0, 0, 0}, 0},
		{[]int{0, 1, 0, 0, 1, 1, 0, 1, 1}, 4},
		{[]int{1, 0, 0, 0, 1, 1}, 2},
		{[]int{1, 0, 0, 0, 1, 1, 0, 0}, 2},
		{[]int{1, 1}, 1},
		{[]int{0, 1, 1}, 2},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_1493_longest_subarray_of_ones_after_deleting_one_element(c.nums))
	}
}

// Идём двумя указателями по массиву. Первый останавливается в начале участка единиц, второй в конце. Считаем длины участков.
func lc_1493_longest_subarray_of_ones_after_deleting_one_element(nums []int) int {
	lastInterval, maxSum := 0, 0
loop:
	for i := 0; i < len(nums); {
		// Начинаем с поиска единицы.
		if nums[i] == 1 {
			for j := i + 1; j < len(nums); j++ {
				// Встретили единицу, ищем ноль.
				if nums[j] == 0 {
					// По найденному нулю определяем длину участка. Складываем его с последней длиной, если она не была занулена.
					interval := j - i
					if sum := interval + lastInterval; sum > maxSum {
						maxSum = sum
					}
					lastInterval = interval
					i = j + 1
					continue loop
				}
			}
			// Если не встретили ни одного нуля вообще, то удаляем одну из единиц и выходим.
			if i == 0 {
				return len(nums) - 1
			}
			// Если входной массив закончился единицей, попадаем сюда.
			if sum := len(nums) - i + lastInterval; sum > maxSum {
				maxSum = sum
				break
			}
		}
		i++
		// Если нулей между участками больше одного, попадаем сюда. Сбрасываем последний участок.
		lastInterval = 0
	}
	return maxSum
}

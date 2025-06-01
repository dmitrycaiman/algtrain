package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0033_search_in_rotated_sorted_array(t *testing.T) {
	cases := [][]int{
		{1, 2, 4, 5, 6, 7, 0},
		{2, 4, 5, 6, 7, 0, 1},
		{4, 5, 6, 7, 0, 1, 2},
		{5, 6, 7, 0, 1, 2, 4},
		{6, 7, 0, 1, 2, 4, 5},
		{7, 0, 1, 2, 4, 5, 6},
		{0, 1, 2, 4, 5, 6, 7},
		{7, 1, 5},
		{5, 7, 1},
		{1, 5, 7},
		{1, 2},
		{2, 1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprint(c), func(t *testing.T) {
			for i, v := range c {
				assert.Equal(t, i, lc_0033_search_in_rotated_sorted_array(c, v))
			}
			assert.Equal(t, -1, lc_0033_search_in_rotated_sorted_array(c, 10))
			assert.Equal(t, -1, lc_0033_search_in_rotated_sorted_array(c, -1))
		},
		)
	}
}

// Назовём свигающееся при вращении начало массива "хвост".
// 1. Вычисляем индекс среднего элемента между границ массива.
// 2. Определяем положение хвоста, слева или справа от среднего (по величинам границ и среднего):
// если левая граница больше других, то хвост слева; если средний больше, то хвост справа; иначе массив отсортирован.
// 3. Определяем, потенциально принадлежит ли элемент половине без хвоста (т.е. отсортированной половине):
// если нет, то он принадлежит половине с хвостом или не существует.
// 5. Сдвигаем границы в сторону половины, в которой предположительно есть элемент.
// 6. Повторить с п. 1, пока не сократим область поиска до 3-х элементов: в этом случае вычисляем наличие целевого элемента сранением.
// Если область поиска оказалась отсортированной, действуем как при бинарном поиске.
func lc_0033_search_in_rotated_sorted_array(nums []int, target int) int {
	a, b, turnLeft := 0, len(nums)-1, false
	for {
		m := a + (b-a)/2
		switch {
		case nums[a] == target:
			return a
		case b <= a:
			return -1
		case nums[b] == target:
			return b
		case nums[m] == target:
			return m
		case nums[a] > nums[m] && nums[a] > nums[b]:
			turnLeft = !(target > nums[m] && target < nums[b])
		case nums[m] > nums[a] && nums[m] > nums[b]:
			turnLeft = target > nums[a] && target < nums[m]
		default:
			turnLeft = nums[m] > target
		}
		if turnLeft {
			b = m - 1
		} else {
			a = m + 1
		}
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0704_binary_search(t *testing.T) {
	cases := []struct {
		nums           []int
		target, result int
	}{
		{[]int{0, 3, 4, 7, 8, 9}, 3, 1},
		{[]int{0, 3, 4, 7, 8, 9}, 3, 1},
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{0}, 6, -1},
		{[]int{1, 2}, 2, 1},
		{[]int{1, 2}, 1, 0},
		{[]int{1, 2}, 3, -1},
		{[]int{1, 2}, 0, -1},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0704_binary_search_iter(c.nums, c.target))
		assert.Equal(t, c.result, lc_0704_binary_search_recur(c.nums, c.target))
		assert.Equal(t, c.result, lc_0704_binary_search(c.nums, c.target))
	}
}

// Бинарный поиск методом двух меток по краям массива. Сдвигаем метку в зависимости от значения среднего элемента между метками.
func lc_0704_binary_search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for {
		length := r - l + 1
		if length == 0 {
			return -1
		}
		mid := l + length/2
		val := nums[mid]
		switch {
		case val > target:
			r = mid - 1
		case val < target:
			l = mid + 1
		default:
			return mid
		}
	}
}

// Устанавливаем метки в индексах начала и конца массива. Вычисляем расстояние между концами, делим его пополам и прибавляем
// получившееся значение к левому индексу: получаем индекс среднего элемента рассматриваемого промежутка. По величине среднего элемента
// определяем, в какую сторону нам следует сдвинуть одну из меток и повторяем процедуру заново. Выходим из цикла, когда метки сойдутся
// или когда средний элемент окажется искомым.
func lc_0704_binary_search_iter(nums []int, target int) int {
	a, b := 0, len(nums)-1
	for {
		if b <= a {
			if nums[a] == target {
				return a
			} else {
				return -1
			}
		}
		i := a + (b-a)/2
		switch {
		case nums[i] > target:
			b = i - 1
		case nums[i] < target:
			a = i + 1
		default:
			return i
		}
	}
}

// Рекурсивный вариант бинарного поиска. Вычисляем средний элемент входного массива и сравниваем его с искомым.
// В зависимости от результата сравнения сдвигаем границы массива (нужно учесть уменьшение номера границы при сдвиге вправо).
// Если средний элемент остался один в массиве, сверяем его с искомым и выдаём результат. Также средний может сразу оказаться искомым.
func lc_0704_binary_search_recur(nums []int, target int) int {
	i := (len(nums) - 1) / 2
	switch {
	case len(nums) == 0:
		return -1
	case nums[i] > target:
		return lc_0704_binary_search_recur(nums[:i], target)
	case nums[i] < target:
		res := lc_0704_binary_search_recur(nums[i+1:], target)
		if res != -1 {
			res += i + 1
		}
		return res
	default:
		if nums[i] == target {
			return i
		} else {
			return -1
		}
	}
}

package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/3sum
func Test_lc_0015_3sum(t *testing.T) {
	cases := []struct {
		input  []int
		output [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{-1}, [][]int{}},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0015_3sum(c.input))
	}
}

// Сортируем массив по возрастанию.
// Фиксируем первый элемент массива. Сейчас целевая сумма — отрицательное значение первого элемента.
// Ставим метки на следующий элемент и на конец. Двигаем метки навстречу.
// Если элементы на метках в сумме больше целевой суммы, сдвигаем левую метку, если меньше — правую.
// Если попали в целевую сумму, то сохраняем значения фиксированного элемента и элементов меток.
// Далее фиксируем следующий элемент и повторяем действия.
// Пропускаем последовательные повторения элементов при всех сдвигах для избежания дублей в результатах.
func lc_0015_3sum(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}
	for i, target := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			switch {
			case k != len(nums)-1 && nums[k] == nums[k+1]:
				k--
				continue
			case j != i+1 && nums[j] == nums[j-1]:
				j++
				continue
			}
			sum := nums[j] + nums[k]
			switch {
			case sum == -target:
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			case sum < -target:
				j++
			default:
				k--
			}
		}
	}
	return result
}

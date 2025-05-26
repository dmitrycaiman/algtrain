package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/two-sum/description/
func Test_lc_0001_two_sum(t *testing.T) {
	cases := []struct {
		nums, output []int
		target       int
	}{
		{nums: []int{1, 4, 5, 3, 7, 1}, target: 7, output: []int{1, 3}},
		{nums: []int{1, 4}, target: 5, output: []int{0, 1}},
		{nums: []int{8, 8, 8, 8, 8, 9, 1}, target: 10, output: []int{5, 6}},
	}
	for _, c := range cases {
		assert.ElementsMatch(t, c.output, lc_0001_two_sum_map(c.nums, c.target))
		assert.ElementsMatch(t, c.output, lc_0001_two_sum_brute(c.nums, c.target))
	}
}

// Создаём хеш-таблицу с целочисленными ключами и значениями. Ключ будет обозначать элемент входного массива,
// а значение — позиция этого элемента во входном массиве.
// Итерируемся по входному массиву. Вычисляем разность target и текущего элемента. Проверяем, есть ли в хеш-таблице значение по ключу, равному этой разности.
// Если есть, то возвращаем индекс текущего элемента и значение по ключу, т.е. индекс "недостающего" элемента.
// Если нет, то добавляем текущий элемент в хеш-таблицу по описанной схеме и продолжаем итерирование.
func lc_0001_two_sum_map(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}

func lc_0001_two_sum_brute(nums []int, target int) []int {
	for i := range len(nums) - 1 {
		for j := i + 1; j < len(nums); j++ {
			switch {
			case nums[i]+nums[j] == target:
				return []int{i, j}
			}
		}
	}
	return nil
}

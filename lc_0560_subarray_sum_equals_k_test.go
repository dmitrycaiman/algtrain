package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// #yandex
// https://leetcode.com/problems/subarray-sum-equals-k
func Test_lc_0560_subarray_sum_equals_k(t *testing.T) {
	cases := []struct {
		nums      []int
		k, result int
	}{
		{[]int{1, 1, 2, 3, -1}, 2, 3},
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1}, 0, 0},
		{[]int{-1, -1, 1}, 0, 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0560_subarray_sum_equals_k(append([]int(nil), c.nums...), c.k))
		assert.Equal(t, c.result, lc_0560_subarray_sum_equals_k_2(append([]int(nil), c.nums...), c.k))
	}
}

// 26/31, 9/79
// Изменяем массив так, чтобы каждый элемент был суммой всех предыдущих. Также формируем хеш-таблицу частот обновлённых элементов.
// Таким образом каждый элемент будет обозначать сумму подмассива с границами от нулевого элемента до себя включительно.
// Проверяем хеш-таблицу по ключу k и прибавляем значение к результату.
// Далее, чтобы рассмотреть все подмассивы, но уже без первого элемента, просто увеличим k на величину первого элемента и,
// уменьшив частоту первого элемента, повторяем проверку по ключу k.
// Так проходимся по всему обновлённому массиву, уменьшая частоты и увеличивая k, что есть по сути перебор сумм всех подмассивов.
func lc_0560_subarray_sum_equals_k(nums []int, k int) int {
	sums, result := map[int]int{nums[0]: 1}, 0
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
		sums[nums[i]]++
	}
	result += sums[k]
	for _, v := range nums {
		sums[v]--
		if sums[k+v] > 0 {
			result += sums[k+v]
		}
	}
	return result
}

// Полный перебор.
func lc_0560_subarray_sum_equals_k_2(nums []int, k int) int {
	result := 0
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				result++
			}
		}
	}
	return result
}

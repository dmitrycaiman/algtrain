package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/intersection-of-two-arrays
func Test_lc_0349_intersection_of_two_arrays(t *testing.T) {
	cases := []struct{ nums1, nums2, result []int }{
		{[]int{1, 2, 1, 3, 5, 6, 6, 4, 3}, []int{1, 1, 1, 1, 1, 1, 1, 2, 3, 6, 7, 8, 9, 9}, []int{1, 2, 3, 6}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 3, 3}, []int{3}, []int{3}},
	}
	for _, c := range cases {
		assert.ElementsMatch(t, c.result, lc_0349_intersection_of_two_arrays(c.nums1, c.nums2))
	}
}

// 100/33
// Выбираем наименьший массив и складываем его значения в хеш-таблицу.
// Проходимся по второму массиву. Если встречаем равный элемент в хеш-таблице,
// то складываем его в результат и удаляем из хеш-таблицы (чтобы не добавлять его несколько раз при повторах).
func lc_0349_intersection_of_two_arrays(nums1 []int, nums2 []int) []int {
	m, result := map[int]struct{}{}, []int{}
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			result = append(result, v)
			delete(m, v)
		}
	}
	return result
}

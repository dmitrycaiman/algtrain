package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/intersection-of-two-arrays-ii
func Test_lc_0350_intersection_of_two_arrays_ii(t *testing.T) {
	cases := []struct{ nums1, nums2, result []int }{
		{[]int{2, 6, 4, 4, 1, 1, 9}, []int{1, 1, 1, 1, 9, 9, 6}, []int{1, 1, 6, 9}},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0350_intersection_of_two_arrays_ii(c.nums1, c.nums2))
	}
}

// Сортируем оба массива. Ставим каждому в начало метку.
// Если элементы на метках равны, добавляем данный элемент в результат.
// Если нет, то метка с меньшим числом двигается вперёд, т.е. догоняет другую метку.
// Заканчиваем, если хоть одна метка дошла до конца.
func lc_0350_intersection_of_two_arrays_ii(nums1 []int, nums2 []int) []int {
	result := []int{}
	slices.Sort(nums1)
	slices.Sort(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		switch {
		case nums1[i] == nums2[j]:
			result = append(result, nums1[i])
			i++
			j++
		case nums1[i] < nums2[j]:
			i++
		default:
			j++
		}
	}
	return result
}

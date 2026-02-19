package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/first-missing-positive
func Test_lc_0041_first_missing_positive(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, -1, 3}, 4},
		{[]int{2, 2, -1, 3}, 1},
		{[]int{3, 2, 1, 5}, 4},
		{[]int{1, 2, 0}, 3},
		{[]int{13, 4, -1, 1}, 2},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0041_first_missing_positive(c.input))
	}
}

// Краевой случай: числа начинаются с единицы и идут подряд до конца массива. Если так, то каждое число имеет своё определённое место в нём.
// Первым проходом ищем такие числа, то есть в пределах от 1 до длины массива, и ставим их на свои места.
// Вторым проходом смотрим, на каком из мест образовалась "дыра" (то есть числа для места не нашлось). Это и будет ответом.
// Если дыры не нашли, то все числа на своих местах, и ответом будет следующее, то есть длина массива + 1.
func lc_0041_first_missing_positive(nums []int) int {
	for i, v1 := range nums {
		for {
			if i+1 != v1 && v1 > 0 && v1 <= len(nums) {
				v2 := nums[v1-1]
				if v1 == v2 {
					break
				}
				nums[v1-1], nums[i], v1 = v1, v2, v2
				continue
			}
			break
		}
	}
	for i, v := range nums {
		if i+1 != v {
			return i + 1
		}
	}
	return len(nums) + 1
}

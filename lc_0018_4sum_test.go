package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/4sum
func Test_lc_0018_4sum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		result [][]int
	}{
		{[]int{1, 2, 2, 3, 4, 1}, 6, [][]int{{1, 2, 2, 1}}},
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{[]int{-2, -1, -1, 1, 1, 2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}}},
		{[]int{-5, 5, 4, -3, 0, 0, 4, -2}, 4, [][]int{{-5, 0, 4, 5}, {-3, -2, 4, 5}}},
		{[]int{0, 2, 2, 2, 10, -3, -9, 2, -10, -4, -9, -2, 2, 8, 7}, 6, [][]int{{-10, -2, 8, 10}, {-9, -3, 8, 10}, {-9, -2, 7, 10}, {-9, 0, 7, 8}, {-4, -2, 2, 10}, {-4, 0, 2, 8}, {-3, 0, 2, 7}, {0, 2, 2, 2}}},
	}
	for _, c := range cases {
		assert.True(t, CheckEqualityWithFunc(lc_0018_4sum(c.nums, c.target), c.result, func(a, b []int) bool { return CheckEquality(a, b) }))
	}
}

// 1319ms (5.03%), 10.94MB (5.03%)
// Рекурсивный проход по возможным кейсам. В каждом кейсе делаем выбор: берём текущий элемент в набор или нет.
// Рекурсия заканчивается, если либо нечем наполнить массив текущего результата, либо он уже полон, но нужная сумма не набрана.
// Самое тяжёлое это определить, был ли текущий результат уже получен.
// Данное решение основано на:
// 1. Хранении структуры с сортированными значениями в мапе.
// 2. Удалении значений с более чем 4-мя копиями.
// Это позволило пройти тесты, но на самом низком уровне. Решение ТРЕБУЕТ ДОРАБОТКИ.
func lc_0018_4sum(nums []int, target int) [][]int {
	return r(prepare(nums), target, []int{}, [][]int{}, map[quad]struct{}{})
}

type quad struct{ a, b, c, d int }

func r(nums []int, target int, cur []int, result [][]int, m map[quad]struct{}) [][]int {
	switch {
	case target == 0 && len(cur) == 4:
		q := quad{cur[0], cur[1], cur[2], cur[3]}
		if _, ok := m[q]; ok {
			return result
		}
		m[q] = struct{}{}
		cp := make([]int, len(cur))
		copy(cp, cur)
		return append(result, cp)
	case len(cur) == 4 || len(nums) < 4-len(cur):
		return result
	}
	return r(nums[1:], target, cur, r(nums[1:], target-nums[0], append(cur, nums[0]), result, m), m)
}

func prepare(in []int) []int {
l:
	for {
		shift := false
		for i := range len(in) - 1 {
			if in[i] > in[i+1] {
				in[i], in[i+1] = in[i+1], in[i]
				shift = true
			}
		}
		if !shift {
			break l
		}
	}
	output, cur, rep := []int{}, in[0], 1
	for _, v := range in {
		switch {
		case v == cur && rep <= 4:
			output = append(output, v)
			rep++
		case v != cur:
			cur = v
			output = append(output, v)
			rep = 1
		}
	}
	return output
}

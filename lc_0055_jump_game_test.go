package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0055_jump_game(t *testing.T) {
	cases := []struct {
		input  []int
		result bool
	}{
		{[]int{1, 1, 1, 1}, true},
		{[]int{1, 1, 0, 1}, false},
		{[]int{3, 0, 0, 0}, true},
		{[]int{2, 0, 0, 0}, false},
		{[]int{3, 3, 0, 1, 1, 0}, true},
		{[]int{3}, true},
		{[]int{0}, true},
		{[]int{1, 10, 0, 0}, true},
		{[]int{1, 2, 0, 3, 0, 0, 4, 0, 0, 0, 5, 0, 0, 0, 0, 0}, true},
		{[]int{1, 2, 0, 3, 0, 0, 8, 0, 0, 0, 4, 0, 0, 0, 0, 0}, false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0055_jump_game(c.input))
		assert.Equal(t, c.result, lc_0055_jump_game_2(c.input))

	}
}

func lc_0055_jump_game(nums []int) bool {
	return lc_0055_check(nums, len(nums)-1)
}

// Обходим массив рекурсивно с конца. Ищем ближайшую точку слева, из которой достижима текущая точка, и рекурсируем.
// Останавливаемся с успехом, когда текущая точка будет левой границей.
func lc_0055_check(nums []int, pos int) bool {
	if pos == 0 {
		return true
	}
	for i := 1; i <= pos; i++ {
		if nums[pos-i] >= i {
			return lc_0055_check(nums, pos-i)
		}
	}
	return false
}

// Обходим массив рекурсивно с начала, прыгая из текущей точки как можно дальше и стараясь попасть на правую границу или за неё.
// Посещённые точки запоминаем. Не всегда эффективно.
func lc_0055_jump_game_2(nums []int) bool {
	return lc_0055_check_2(nums, 0, map[int]struct{}{})
}

func lc_0055_check_2(nums []int, pos int, visited map[int]struct{}) bool {
	if pos >= len(nums)-1 {
		return true
	}
	visited[pos] = struct{}{}
	for i := pos + nums[pos]; i > pos; i-- {
		if _, ok := visited[i]; !ok && lc_0055_check_2(nums, i, visited) {
			return true
		}
	}
	return false
}

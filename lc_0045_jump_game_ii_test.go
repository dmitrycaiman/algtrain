package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/jump-game-ii
func Test_lc_0045_jump_game_ii(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 1, 1, 1}, 3}, {[]int{1}, 0}, {[]int{2, 30, 1, 1, 4}, 2}, {[]int{2, 40, 0, 1, 4}, 2}, {[]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}, 3},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0045_jump_game_ii_routes(c.input))
		assert.Equal(t, c.output, lc_0045_jump_game_ii_dir(c.input))
		assert.Equal(t, c.output, lc_0045_jump_game_ii_rec(c.input))
		assert.Equal(t, c.output, lc_0045_jump_game_ii_map(c.input))
		assert.Equal(t, c.output, lc_0045_jump_game_ii_slice(c.input))
	}
}

// 10 ms, 8.13 MB
// Решение через простую рекурсию.
func lc_0045_jump_game_ii_rec(nums []int) int {
	jumpIndex, lastIndex := -1, len(nums)-1
	for i := range lastIndex {
		if nums[i] >= lastIndex-i {
			jumpIndex = i
			break
		}
	}
	switch jumpIndex {
	case -1:
		return 0
	case 0, 1:
		return jumpIndex + 1
	default:
		return lc_0045_jump_game_ii_rec(nums[:jumpIndex+1]) + 1
	}
}

// 13 ms, 8.06 MB
// Решение через прямой расчёт.
func lc_0045_jump_game_ii_dir(nums []int) int {
	target, i, res := len(nums)-1, 0, 0
	if target == 0 {
		return 0
	}
	for {
		if nums[i] >= target-i {
			res++
			if i == 0 {
				return res
			}
			target, i = i, 0
		} else {
			i++
		}
	}
}

// 20 ms, 8.01 MB
// Решение через аналитический слайс, элементы которого обозначают тот индекс, из которого наиболее выгодно можно прыгнуть в текущий.
func lc_0045_jump_game_ii_routes(nums []int) (res int) {
	routes, mark := make([]int, len(nums)), 0
	for i := range len(routes) - 1 {
		for j := 1; i+j < len(nums) && j <= nums[i]; j++ {
			if i+j > mark {
				routes[i+j] = i
				mark++
			}
		}
	}
	for {
		if mark == 0 {
			return res
		}
		mark = routes[mark]
		res++
	}
}

// 106 ms, 11.33 MB
// Перебор с мемоизацией через слайс.
func lc_0045_jump_game_ii_slice(nums []int) (res int) {
	cache := make([]int, len(nums))
	for i := range cache {
		cache[i] = -1
	}
	return lc_0045_slice(0, nums, cache)
}

func lc_0045_slice(pos int, nums []int, cache []int) (best int) {
	if cached := cache[pos]; cached != -1 {
		return cached
	}
	defer func() { cache[pos] = best }()
	for i := 1; i <= nums[pos] && i < len(nums); i++ {
		if pos+i == len(nums)-1 {
			return 1
		}
		res := lc_0045_slice(pos+i, nums, cache)
		if res != 0 && (res < best || best == 0) {
			best = res + 1
		}
	}
	return best
}

// 410 ms, 11.08 MB
// Перебор с мемоизацией через мапу.
func lc_0045_jump_game_ii_map(nums []int) int {
	cache := map[int]int{}
	return lc_0045_map(0, nums, cache)
}

func lc_0045_map(pos int, nums []int, cache map[int]int) (best int) {
	if cached, ok := cache[pos]; ok {
		return cached
	}
	defer func() { cache[pos] = best }()
	for i := 1; i <= nums[pos] && i < len(nums); i++ {
		if pos+i == len(nums)-1 {
			return 1
		}
		res := lc_0045_map(pos+i, nums, cache)
		if res != 0 && (res < best || best == 0) {
			best = res + 1
		}
	}
	return best
}

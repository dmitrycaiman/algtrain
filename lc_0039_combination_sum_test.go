package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/combination-sum
func Test_lc_0039_combination_sum(t *testing.T) {
	cases := []struct {
		candidates []int
		target     int
		result     [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{1, 2, 3}, 3, [][]int{{1, 1, 1}, {1, 2}, {3}}},
	}
	for _, c := range cases {
		assert.ElementsMatch(t, c.result, lc_0039_combination_sum(c.candidates, c.target))
	}
}

// Сортируем массив кандидатов по возрастанию.
// Рекурсивно перебираем все возможные комбинации с повторами, отслеживаем сумму.
func lc_0039_combination_sum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	return lc_0039(candidates, []int{}, target, 0, [][]int{})
}

func lc_0039(candidates, result []int, target, sum int, storage [][]int) [][]int {
	for i, v := range candidates {
		currentSum := sum + v
		if currentSum <= target {
			currentResult := make([]int, len(result))
			copy(currentResult, result)
			currentResult = append(currentResult, v)
			if currentSum == target {
				storage = append(storage, currentResult)
			} else {
				storage = lc_0039(candidates[i:], currentResult, target, currentSum, storage)
			}
		} else {
			break
		}
	}
	return storage
}

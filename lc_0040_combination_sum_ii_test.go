package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/combination-sum-ii
func Test_lc_0040_combination_sum_ii(t *testing.T) {
	cases := []struct {
		candidates []int
		target     int
		result     [][]int
	}{
		{[]int{2, 3, 1, 1, 6}, 6, [][]int{{1, 2, 3}, {6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}
	for _, c := range cases {
		assert.ElementsMatch(t, c.result, lc_0040_combination_sum_ii(c.candidates, c.target))
	}
}

// Сортируем массив кандидатов по возрастанию.
// Рекурсивно перебираем все возможные комбинации без повторов, отслеживаем сумму.
func lc_0040_combination_sum_ii(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	return lc_0040(candidates, []int{}, target, 0, [][]int{})
}

func lc_0040(candidates, result []int, target, sum int, storage [][]int) [][]int {
	last, currentSum := 0, 0
	for i, v := range candidates {
		if v == last {
			continue
		}
		last, currentSum = v, sum+v
		if currentSum <= target {
			currentResult := make([]int, len(result)+1)
			copy(currentResult, result)
			currentResult[len(currentResult)-1] = v
			if currentSum == target {
				storage = append(storage, currentResult)
			} else {
				storage = lc_0040(candidates[i+1:], currentResult, target, currentSum, storage)
			}
		} else {
			break
		}
	}
	return storage
}

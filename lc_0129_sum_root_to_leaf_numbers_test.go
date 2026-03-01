package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-root-to-leaf-numbers
func Test_lc_0129_sum_root_to_leaf_numbers(t *testing.T) {
	cases := []struct {
		tree   string
		result int
	}{
		{"1,2,3,4", 124 + 13},
		{"1,2,3", 12 + 13},
		{"4,9,0,5,1", 495 + 491 + 40}}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0129_sum_root_to_leaf_numbers(NewTree(c.tree)))
	}
}

// В каждом встретившемся узле вызываем рекурсивную функцию. Функция получает предыдущий результат, умножает его на 10 (т.е. сдвигает разряд)
// и прибавляет значение узла. Далее вызывает рекурсию в своих детях с модифицированным результатом и складывает получившееся.
// Если детей нет, то отдаём уже модифицированный результат.
func lc_0129_sum_root_to_leaf_numbers(root *TreeNode) int { return lc_0129(0, root) }

func lc_0129(result int, node *TreeNode) int {
	sum, result := 0, result*10+node.Val
	if node.Left != nil {
		sum += lc_0129(result, node.Left)
	}
	if node.Right != nil {
		sum += lc_0129(result, node.Right)
	}
	if sum == 0 {
		return result
	}
	return sum
}

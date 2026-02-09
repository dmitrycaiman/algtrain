package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/path-sum
func Test_lc_0112_path_sum(t *testing.T) {
	cases := []struct {
		scheme    string
		targetSum int
		result    bool
	}{
		{"5,1,2,10", 7, true},
		{"5,1,2,10", 10, false},
		{"1,2,3,4,5", 8, true},
		{"", 0, false},
		{"5,4,8,11,null,13,4,7,2,null,null,null,1", 22, true},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0112_path_sum(NewTree(c.scheme), c.targetSum))
	}
}

// Рекурсивно обходим дерево. Если пришли в лист и набрали сумму, то выходим с успехом.
// Иначее пытаемся сходит в обе ветви и повторить эту проверку.
func lc_0112_path_sum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil && lc_0112_path_sum(root.Left, targetSum) {
		return true
	}
	return root.Right != nil && lc_0112_path_sum(root.Right, targetSum)
}

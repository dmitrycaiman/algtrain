package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/symmetric-tree
func Test_lc_0101_symmetric_tree(t *testing.T) {
	cases := []struct {
		tree   string
		result bool
	}{
		{"1,2,3,4", false},
		{"1,2,2", true},
		{"1,2,2,4,3,3,4", true},
		{"1,2,2,4,3,3", false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0101_symmetric_tree(NewTree(c.tree)))
	}
}

// Симметричность проверяем рекурсивно через равенство противолежащих узлов (левый с правым, правый с левым).
func lc_0101_symmetric_tree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return lc_0101_symmetric_tree_check(root.Left, root.Right)
}

func lc_0101_symmetric_tree_check(t1, t2 *TreeNode) bool {
	switch {
	case t1 == nil && t2 == nil:
		return true
	case t1 == nil || t2 == nil || t1.Val != t2.Val:
		return false
	default:
		return lc_0101_symmetric_tree_check(t1.Right, t2.Left) && lc_0101_symmetric_tree_check(t1.Left, t2.Right)
	}
}

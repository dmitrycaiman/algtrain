package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/same-tree
func Test_lc_0100_same_tree(t *testing.T) {
	cases := []struct {
		t1, t2 string
	}{
		{"1,2,3,4", "1,2,3,4"},
		{"1,2,3,4", "1,2,3"},
		{"1", "1"},
		{"1,2,3,4", "1,2,4,3"},
	}
	for _, c := range cases {
		assert.Equal(t, c.t1 == c.t2, lc_0100_same_tree(NewTree(c.t1), NewTree(c.t2)))
	}
}

// Рекурсивное сравнение каждого узла с выходом при первом же неравенстве.
func lc_0100_same_tree(p *TreeNode, q *TreeNode) bool {
	switch {
	case p == nil && q == nil:
		return true
	case p == nil || q == nil || p.Val != q.Val:
		return false
	default:
		return lc_0100_same_tree(p.Left, q.Left) && lc_0100_same_tree(p.Right, q.Right)
	}
}

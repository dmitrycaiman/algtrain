package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/balanced-binary-tree
func Test_lc_0110_balanced_binary_tree(t *testing.T) {
	cases := []struct {
		tree   string
		result bool
	}{
		{"1", true},
		{"1,2", true},
		{"1,2,null,3", false},
		{"1,2,3", true},
		{"1,2,3,4,5", true},
		{"1,2,3,4,5,null,null", true},
		{"1,2,3,4,5,null,null,6", false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0110_balanced_binary_tree(NewTree(c.tree)))
	}
}

// Рекурсивно проверяем баланс дерева. Считаем длины поддеревьев и сигнализируем о небалансе вверх по стеку вызовов.
// При получении сигнала дальнейших вычислений не проводим.
func lc_0110_balanced_binary_tree(root *TreeNode) bool {
	return lc_0110_check(root, 0) != -1
}

func lc_0110_check(node *TreeNode, counter int) int {
	if node == nil {
		return counter
	}
	counter++
	l1 := lc_0110_check(node.Left, counter)
	if l1 == -1 {
		return -1
	}
	l2 := lc_0110_check(node.Right, counter)
	if l2 == -1 {
		return -1
	}
	if l1 < l2 {
		l1, l2 = l2, l1
	}
	if l1-l2 > 1 {
		return -1
	}
	return l1
}

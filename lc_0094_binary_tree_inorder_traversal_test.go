package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-inorder-traversal
func Test_lc_0094_binary_tree_inorder_traversal(t *testing.T) {
	cases := []struct {
		input  string
		result []int
	}{
		{"1,2,3,4,5,null,8,null,null,6,7,9", []int{4, 2, 6, 5, 7, 1, 3, 9, 8}},
		{"1,null,2,3", []int{1, 3, 2}},
		{"1", []int{1}},
		{"", []int{}},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0094_binary_tree_inorder_traversal(NewTree(c.input)))
	}
}

// Пытаемся сразу пойти в левый узел, а текущий кладём в стек. Когда доходим до листа, то записываем его значение
// и берём следующий узел из стека. При этом у взятого из стека узла также запоминаем значение
// и налево из него уже не ходим, только направо. Придя направо, всё повторяем до тех пор, пока уже некуда будет идти.
func lc_0094_binary_tree_inorder_traversal(root *TreeNode) []int {
	stack, result, fromStack := []*TreeNode{}, []int{}, false
	for {
		switch {
		case root != nil && root.Left != nil && !fromStack:
			stack = append(stack, root)
			root = root.Left
		case root != nil:
			result, fromStack = append(result, root.Val), false
			switch {
			case root.Right != nil:
				root = root.Right
			case len(stack) != 0:
				root = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fromStack = true
			default:
				return result
			}
		default:
			return result
		}
	}
}

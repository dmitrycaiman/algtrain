package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/invert-binary-tree
func Test_lc_0226_invert_binary_tree(t *testing.T) {
	cases := []struct {
		input, output string
	}{
		{"4,2,7,1,3,6,9", "4,7,2,9,6,3,1"},
		{"2,3,null,1", "2,null,3,null,1"},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0226_invert_binary_tree(NewTree(c.input)).String())
	}
}

// Обходим дерево preorder. Меняем узлы местами, когда не идём налево или когда в листе.
func lc_0226_invert_binary_tree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack, fromStack, node := []*TreeNode{}, false, root
	for {
		switch {
		case !fromStack && node.Left != nil:
			stack = append(stack, node)
			node = node.Left
		case node.Left == nil && node.Right == nil:
			if len(stack) == 0 {
				return root
			}
			node, fromStack, stack = stack[len(stack)-1], true, stack[:len(stack)-1]
		default:
			node.Left, node.Right, fromStack = node.Right, node.Left, false
			switch {
			case node.Left != nil:
				node = node.Left
			case len(stack) == 0:
				return root
			default:
				node, fromStack, stack = stack[len(stack)-1], true, stack[:len(stack)-1]
			}
		}
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/validate-binary-search-tree
func Test_lc_0098_validate_binary_search_tree(t *testing.T) {
	cases := []struct {
		root   string
		result bool
	}{
		{"2,1,3", true},
		{"5,1,8,null,null,7,10", true},
		{"5,4,6,null,null,3,7", false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0098_validate_binary_search_tree(NewTree(c.root)))
	}
}

// Делаем preorder-обход дерева. Проверяем значения узлов, когда идём не в левый узел.
// Стоит отметить, что при переходе в правый узел нужно устанавливать минимум для остальных встречающихся далее узлов.
func lc_0098_validate_binary_search_tree(root *TreeNode) bool {
	for stack, fromStack, min := []*TreeNode{}, false, -3_000_000_000; ; {
		switch {
		case !fromStack && root.Left != nil:
			stack = append(stack, root)
			root = root.Left
		case root.Right != nil:
			if root.Val >= root.Right.Val {
				return false
			}
			root, fromStack, min = root.Right, false, root.Val
		default:
			if len(stack) == 0 {
				return true
			}
			if root.Val >= stack[len(stack)-1].Val || root.Val <= min {
				return false
			}
			root, fromStack, stack = stack[len(stack)-1], true, stack[:len(stack)-1]
		}
	}
}

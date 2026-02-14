package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-preorder-traversal
func Test_lc_0144_binary_tree_preorder_traversal(t *testing.T) {
	cases := []struct {
		input  string
		result []int
	}{
		{"1,null,2,3", []int{1, 2, 3}},
		{"1,2,3,4,5,6,7,8,9", []int{1, 2, 4, 8, 9, 5, 3, 6, 7}},
		{"1", []int{1}},
		{"", []int{}},
		{"1,null,3,4", []int{1, 3, 4}},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0144_binary_tree_preorder_traversal(NewTree(c.input)))
	}
}

// Заходим в узел. Если он не пустой, то записываем его значение. Кладём правый узел в стек и идём в левый.
// На следующей итерации, когда в левый узел идти уже невозможно, берём узел из стека и делаем то же самое.
// Выходим с результатом, когда не будет левых узлов и стек будет пуст.
func lc_0144_binary_tree_preorder_traversal(root *TreeNode) []int {
	stack, result := []*TreeNode{}, []int{}
	for {
		switch {
		case root != nil:
			result = append(result, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root = root.Left
		case len(stack) != 0:
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			return result
		}
	}
}

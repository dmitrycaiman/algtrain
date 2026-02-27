package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal
func Test_lc_0103_binary_tree_zigzag_level_order_traversal(t *testing.T) {
	cases := []struct {
		tree   string
		output [][]int
	}{
		{"1,2,3,4,5,6", [][]int{{1}, {3, 2}, {4, 5, 6}}},
		{"1,2,null,4,5,6,7", [][]int{{1}, {2}, {4, 5}, {7, 6}}},
		{"3,9,20,null,null,15,7", [][]int{{3}, {20, 9}, {15, 7}}},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0103_binary_tree_zigzag_level_order_traversal(NewTree(c.tree)))
	}
}

// На каждом уровне собираем массив следующих узлов.
// К примеру, на первом уровне будет корень.
// На втором будет массив из двух его детей, слева направо.
// На третьем будут дети этих двух узлов, вытащенные так же слева направо, и так далее.
// Рекурсивно собираем эти массивы, на каждом подходе рекурсии вытаскиваем значения то в одном направлении, то в другом.
// Как только не получится собрать следующий уровень, возвращаемся из рекурсии с результатами.
func lc_0103_binary_tree_zigzag_level_order_traversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	return lc_0103([]*TreeNode{root}, [][]int{}, true)
}

func lc_0103(nodes []*TreeNode, result [][]int, left bool) [][]int {
	level, nextNodes := make([]int, len(nodes)), []*TreeNode{}
	if left {
		for i := 0; i < len(nodes); i++ {
			level[i] = nodes[i].Val
		}
	} else {
		for i := len(nodes) - 1; i >= 0; i-- {
			level[len(nodes)-1-i] = nodes[i].Val
		}
	}
	result = append(result, level)
	for _, v := range nodes {
		if v.Left != nil {
			nextNodes = append(nextNodes, v.Left)
		}
		if v.Right != nil {
			nextNodes = append(nextNodes, v.Right)
		}
	}
	if len(nextNodes) == 0 {
		return result
	}
	return lc_0103(nextNodes, result, !left)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0113_path_sum_ii(t *testing.T) {
	cases := []struct {
		root   string
		target int
		output [][]int
	}{
		{"", 0, [][]int{}},
		{"", 1000, [][]int{}},
		{"1000", 1000, [][]int{{1000}}},
		{"500,500,300,0,null,1,200", 1000, [][]int{{500, 500, 0}, {500, 300, 200}}},
		{"500,-500,700,1000,null,1,-200", 1000, [][]int{{500, -500, 1000}, {500, 700, -200}}},
		{"5,4,8,11,null,13,4,7,2,null,null,5,1", 22, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
	}
	for _, c := range cases {
		assert.ElementsMatch(t, c.output, lc_0113_path_sum_ii(NewTree(c.root), c.target))
	}
}

func lc_0113_path_sum_ii(root *TreeNode, targetSum int) [][]int {
	return lc_0113_path_sum_ii_rec(root, 0, targetSum, []int{}, [][]int{})
}

func lc_0113_path_sum_ii_rec(node *TreeNode, sum, targetSum int, path []int, results [][]int) [][]int {
	if node == nil {
		return results
	}
	sum += node.Val
	path = append(path, node.Val)
	switch {
	case node.Left == nil && node.Right == nil:
		if sum == targetSum {
			results = append(results, path)
		}
		return results
	case node.Left == nil && node.Right != nil:
		return lc_0113_path_sum_ii_rec(node.Right, sum, targetSum, path, results)
	case node.Left != nil && node.Right == nil:
		return lc_0113_path_sum_ii_rec(node.Left, sum, targetSum, path, results)
	}
	leftPath := make([]int, len(path))
	copy(leftPath, path)
	return lc_0113_path_sum_ii_rec(
		node.Left,
		sum,
		targetSum,
		leftPath,
		lc_0113_path_sum_ii_rec(node.Right, sum, targetSum, path, results),
	)
}

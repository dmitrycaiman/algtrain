package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(scheme string) *TreeNode {
	nodes := []*TreeNode{}
	for _, s := range strings.Split(scheme, ",") {
		val, err := strconv.Atoi(s)
		if err != nil {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, &TreeNode{Val: val})
		}
	}
	if len(nodes) == 0 {
		return nil
	}
	for i, j := 0, 1; j < len(nodes); {
		if nodes[i] != nil {
			nodes[i].Left = nodes[j]
			if j+1 < len(nodes) {
				nodes[i].Right = nodes[j+1]
			}
			i, j = i+1, j+2
		} else {
			i++
		}
	}
	return nodes[0]
}

func (slf *TreeNode) String() string {
	result, nodes, counter, i := "", []*TreeNode{slf}, 0, 0
	for counter != slf.NodesCount() {
		if nodes[i] != nil {
			counter++
			nodes = append(nodes, nodes[i].Left, nodes[i].Right)
			result += fmt.Sprint(",", nodes[i].Val)
		} else {
			result += ",null"
		}
		i++
	}
	return result
}
func (slf *TreeNode) NodesCount() int { return slf.num(0) }

func (slf *TreeNode) num(n int) int {
	n++
	switch {
	case slf == nil:
		return 0
	case slf.Left == nil && slf.Right == nil:
		return n
	case slf.Left == nil && slf.Right != nil:
		return slf.Right.num(n)
	case slf.Left != nil && slf.Right == nil:
		return slf.Left.num(n)
	}
	return slf.Left.num(slf.Right.num(n))
}

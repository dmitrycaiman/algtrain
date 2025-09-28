package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0108_convert_sorted_array_to_binary_search_tree(t *testing.T) {
	cases := []struct {
		input  []int
		output []string
	}{
		{[]int{-10, -3, 0, 5, 9}, []string{"0", "-3", "9", "-10", "null", "5", "null"}},
		{[]int{1, 2, 3, 4, 5}, []string{"3", "2", "5", "1", "null", "4", "null"}},
		{[]int{1, 2, 3}, []string{"2", "1", "3"}},
		{[]int{1, 2, 3, 4}, []string{"3", "2", "4", "1", "null"}},
	}
	for _, v := range cases {
		root := lc_0108_convert_sorted_array_to_binary_search_tree(v.input)
		var dst *[]string = &[]string{fmt.Sprint(root.Val)}
		unpackTree([]*TreeNode{root}, dst)
		assert.Equal(t, v.output, *dst)
	}

}

// Исходный массив есть ветка. Ломаем ветку пололам, т.е. вычисляем средний элемент целочисленным делением длины на 2. Присваиваем значение
// среднего элемента текущему узлу. Берём обломки ветки. Левый отдаём левому потомку, правый — правому, и рекурсируем. На краевых случаях,
// когда длина полученной ветки 3 и менее элемента, расставляем значения родителя и потомков сразу.
func lc_0108_convert_sorted_array_to_binary_search_tree(nums []int) *TreeNode {
	root := &TreeNode{}
	f(root, nums)
	return root
}

func unpackTree(src []*TreeNode, dst *[]string) {
	newSrc := []*TreeNode{}
	for _, v := range src {
		tmp := *dst
		if v.Left != nil {
			tmp = append(tmp, fmt.Sprint(v.Left.Val))
			if v.Left.Left != nil || v.Left.Right != nil {
				newSrc = append(newSrc, v.Left)
			}
		} else {
			tmp = append(tmp, "null")
		}
		if v.Right != nil {
			tmp = append(tmp, fmt.Sprint(v.Right.Val))
			if v.Right.Left != nil || v.Right.Right != nil {
				newSrc = append(newSrc, v.Right)
			}
		} else {
			tmp = append(tmp, "null")
		}
		*dst = tmp
	}
	if len(newSrc) != 0 {
		unpackTree(newSrc, dst)
	}
}

func f(parent *TreeNode, input []int) {
	switch len(input) {
	case 3:
		parent.Val = input[1]
		parent.Left = &TreeNode{Val: input[0]}
		parent.Right = &TreeNode{Val: input[2]}
	case 2:
		parent.Val = input[1]
		parent.Left = &TreeNode{Val: input[0]}
	case 1:
		parent.Val = input[0]
	case 0:
	default:
		mid := len(input) / 2
		parent.Val = input[mid]
		parent.Left, parent.Right = &TreeNode{}, &TreeNode{}
		f(parent.Left, input[:mid])
		f(parent.Right, input[mid+1:])
	}
}

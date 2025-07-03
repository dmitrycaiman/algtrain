package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-two-sorted-lists
func Test_lc_0021_merge_two_sorted_lists(t *testing.T) {
	cases := []struct {
		list1, list2, result string
	}{
		{"1,2,5", "3,4", "1,2,3,4,5"},
		{"1,2,5", "0,2,3,4,5", "0,1,2,2,3,4,5,5"},
		{"", "0", "0"},
		{"", "", ""},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0021_merge_two_sorted_lists(NewList(c.list1), NewList(c.list2)).String())
	}
}

// 0ms (100%), 4.38MB (83.3%)
// Ставим в начало каждого листа метку. Сравниваем узлы, на которых стоят метки. Узел с меньшим значением
// привязываем к результирующему списку и сдвигаем метку на его потомка. Как только один лист кончается, добавляем все элементы другого.
func lc_0021_merge_two_sorted_lists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	next := result
	for {
		switch {
		case (list1 != nil && list2 != nil && list1.Val <= list2.Val) || (list1 != nil && list2 == nil):
			next.Next = list1
			list1 = list1.Next
		case (list1 != nil && list2 != nil && list1.Val > list2.Val) || (list1 == nil && list2 != nil):
			next.Next = list2
			list2 = list2.Next
		default:
			return result.Next
		}
		next = next.Next
	}
}

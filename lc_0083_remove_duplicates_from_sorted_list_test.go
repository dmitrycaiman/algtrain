package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-list
func Test_lc_0083_remove_duplicates_from_sorted_list(t *testing.T) {
	cases := []struct {
		input, output string
	}{
		{"1,2,2,2,3,4", "1,2,3,4"},
		{"1,1,2,2,2,3,4,4", "1,2,3,4"},
		{"1,1,1,1", "1"},
		{"", ""},
	}
	for _, c := range cases {
		assert.Equal(t, NewList(c.output), lc_0083_remove_duplicates_from_sorted_list(NewList(c.input)))
	}
}

// Сохраняем корень. Итерируемся по списку. Если встречаем равный узел, пропускаем его, присоединяясь к следующему.
func lc_0083_remove_duplicates_from_sorted_list(head *ListNode) *ListNode {
	root := &ListNode{Val: -101, Next: head}
	head = root
	for {
		switch {
		case head.Next == nil:
			return root.Next
		case head.Val == head.Next.Val:
			head.Next = head.Next.Next
		default:
			head = head.Next
		}
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/add-two-numbers
func Test_lc_0002_add_two_numbers(t *testing.T) {
	cases := []struct {
		l1, l2, result string
	}{
		{"1,2,3", "4,5,6", "5,7,9"},
		{"9,9,9,9,9,9,9", "9,9,9,9", "8,9,9,9,0,0,0,1"},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0002_add_two_numbers(NewList(c.l1), NewList(c.l2)).String())
	}
}

// Производим поразрядное сложение соответствующих значений списков, как "в столбик".
func lc_0002_add_two_numbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	last := result
	overflow := 0
	for l1 != nil || l2 != nil {
		sum := overflow
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		last.Next = &ListNode{Val: sum % 10}
		last = last.Next
		overflow = sum / 10
	}
	if overflow != 0 {
		last.Next = &ListNode{Val: overflow}
	}
	return result.Next
}

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/add-two-numbers

func Test_lc_0002_add_two_numbers(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 7,
		},
	}
	l3 := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l4 := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	l5 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	assert.True(t, lc_0002_CheckEquality(l3, lc_0002_add_two_numbers(l1, l2)))
	assert.True(t, lc_0002_CheckEquality(l5, lc_0002_add_two_numbers(l3, l4)))
	assert.True(t, lc_0002_CheckEquality(l3, lc_0002_add_two_numbers(nil, l3)))
	assert.True(t, lc_0002_CheckEquality(l5, lc_0002_add_two_numbers(l5, nil)))
	assert.Nil(t, lc_0002_add_two_numbers(nil, nil))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func lc_0002_add_two_numbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	overflow := false
	a, b := 0, 0
	for {
		switch {
		case l1 != nil && l2 != nil:
			a, b = l1.Val, l2.Val
			l1, l2 = l1.Next, l2.Next
		case l1 != nil && l2 == nil:
			a, b = l1.Val, 0
			l1 = l1.Next
		case l1 == nil && l2 != nil:
			a, b = 0, l2.Val
			l2 = l2.Next
		case l1 == nil && l2 == nil:
			if overflow {
				current.Next = &ListNode{Val: 1}
			}
			return result.Next
		}
		current.Next = &ListNode{}
		current = current.Next
		if overflow {
			overflow = false
			a++
		}
		current.Val = a + b
		if current.Val >= 10 {
			overflow = true
			current.Val -= 10
		}
	}
}

func lc_0002_CheckEquality(l1, l2 *ListNode) bool {
	for {
		switch {
		case l1 == nil && l2 != nil:
		case l1 != nil && l2 == nil:
		case l1 == nil && l2 == nil:
			return true
		case l1.Val == l2.Val:
			l1 = l1.Next
			l2 = l2.Next
			continue
		}
		return false
	}
}

func (slf *ListNode) String() (result string) {
	result += "["
	current := slf
	for current != nil {
		result += fmt.Sprint(current.Val, ",")
		current = current.Next
	}
	return result[:len(result)-1] + "]"
}

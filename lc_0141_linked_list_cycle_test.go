package main

import (
	"testing"
)

func Test_lc_0141_linked_list_cycle(t *testing.T) {
	lc_0141_linked_list_cycle(nil)
	lc_0141_linked_list_cycle_naive(nil)
}

func lc_0141_linked_list_cycle(head *ListNode) bool {
	tort, hare := head, head
	for {
		if hare == nil || hare.Next == nil || hare.Next.Next == nil {
			return false
		}
		hare = hare.Next.Next.Next
		tort = tort.Next
		if hare == tort {
			return true
		}
	}
}

func lc_0141_linked_list_cycle_naive(head *ListNode) bool {
	m := map[*ListNode]struct{}{}
	for {
		if head != nil {
			if _, ok := m[head]; ok {
				return true
			}
			m[head] = struct{}{}
			head = head.Next
		} else {
			break
		}
	}
	return false
}

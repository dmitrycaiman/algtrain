package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Unpack формирует строку связного списка от текущего узла до хвоста либо до позиции циклирования.
func (slf *ListNode) String() (result string) {
	m, res, next := map[*ListNode]struct{}{}, "[", slf
	for {
		if _, ok := m[next]; ok {
			return res + fmt.Sprintf("(%v)]", next.Val)
		}
		m[next] = struct{}{}
		res += fmt.Sprintf("%v,", next.Val)
		if next.Next == nil {
			return res[:len(res)-1] + "]"
		}
		next = next.Next
	}
}

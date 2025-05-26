package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(scheme string) *ListNode {
	m, root := map[int]*ListNode{}, &ListNode{}
	last := root
	for _, v := range strings.Split(scheme, ",") {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		if existingNode, ok := m[n]; ok {
			last.Next = existingNode
			return root.Next
		}
		m[n] = &ListNode{Val: n}
		last.Next = m[n]
		last = last.Next
	}
	return root.Next
}

// Unpack формирует строку связного списка от текущего узла до хвоста либо до позиции циклирования.
func (slf *ListNode) String() (result string) {
	if slf == nil {
		return ""
	}
	m, res, next := map[*ListNode]struct{}{}, "", slf
	for {
		if _, ok := m[next]; ok {
			return res + fmt.Sprint(next.Val)
		}
		m[next] = struct{}{}
		res += fmt.Sprintf("%v,", next.Val)
		if next.Next == nil {
			return res[:len(res)-1]
		}
		next = next.Next
	}
}

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

// NewList создаёт связный список согласно схеме. Узлы могут повторяться.
func NewList(scheme string) *ListNode {
	root, _ := NewListWithTail(scheme)
	return root
}

func NewListWithTail(scheme string) (*ListNode, *ListNode) {
	root := &ListNode{}
	last := root
	for _, v := range strings.Split(scheme, ",") {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, nil
		}
		last.Next = &ListNode{Val: n}
		last = last.Next
	}
	return root.Next, last
}

// NewListWithCycle создаёт связный список согласно схеме. Зацикливание производится по указанной позиции.
func NewListWithCycle(scheme string, pos int) *ListNode {
	if pos < 0 {
		return NewList(scheme)
	}
	root := &ListNode{}
	last, length := root, 0
	for _, v := range strings.Split(scheme, ",") {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		last.Next = &ListNode{Val: n}
		last = last.Next
		length++
	}
	if pos <= length {
		cyclePos := root.Next
		for range pos {
			cyclePos = cyclePos.Next
		}
		last.Next = cyclePos
	}
	return root.Next
}

// String формирует строку связного списка от текущего узла до хвоста либо до позиции циклирования.
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

// Reverse рекурсивно производит разворот связного списка и возвращает новый корень.
func (slf *ListNode) Reverse() *ListNode {
	if slf != nil && slf.Next != nil {
		tail := slf.Next.Reverse()
		slf.Next.Next = slf
		if tail == nil {
			tail = slf.Next
		}
		slf.Next = nil
		return tail
	}
	return slf
}

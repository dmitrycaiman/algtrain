package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0141_linked_list_cycle(t *testing.T) {
	noCycle, withCycle := lc_0141_generateTestCase(10, -1), lc_0141_generateTestCase(10, 9)
	assert.True(t, lc_0141_linked_list_cycle(withCycle))
	assert.True(t, lc_0141_linked_list_cycle_naive(withCycle))
	assert.False(t, lc_0141_linked_list_cycle(noCycle))
	assert.False(t, lc_0141_linked_list_cycle_naive(noCycle))
}

// Решение по типу "Черепаха и Заяц". Суть в том, чтобы запустить по списку две метки с разной скоростью и ждать либо момента их встречи
// (значит, в списке есть цикл), либо встречи более быстрой меткой узла без дальнейших связей (значит, цикла нет).
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

// Наивное решение перебором с сохранением указателей на уже просмотренные узлы в хеш-таблице.
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

// lc_0141_generateTestCase генерирует связный список длины length с позицией циклирования pos. Возможно зацикливание хвоста на себя.
func lc_0141_generateTestCase(length, pos int) *ListNode {
	root := &ListNode{Val: 0}
	next := root
	var ins *ListNode
	for i := range length {
		if i == pos {
			ins = next
		}
		if i != length-1 {
			next.Next = &ListNode{Val: i + 1}
			next = next.Next
		}
	}
	if ins != nil {
		next.Next = ins
	}
	return root
}

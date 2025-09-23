package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0141_linked_list_cycle(t *testing.T) {
	cases := []struct {
		list     string
		cyclePos int
		cycled   bool
	}{
		{"1,2,3", 0, true},
		{"1,2,3", 1, true},
		{"1,2,3", 2, true},
		{"1,2,3", 3, false},
	}
	for _, c := range cases {
		assert.Equal(t, c.cycled, lc_0141_linked_list_cycle(NewListWithCycle(c.list, c.cyclePos)))
		assert.Equal(t, c.cycled, lc_0141_linked_list_cycle_naive(NewListWithCycle(c.list, c.cyclePos)))
	}
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
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-linked-list/
func Test_lc_0206_reverse_linked_list(t *testing.T) {
	cases := []struct{ input, output string }{
		{"1,2,3,4", "4,3,2,1"},
		{"", ""},
		{"1", "1"},
		{"1,2", "2,1"},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0206_reverse_linked_list_recur(NewList(c.input)).String())
		assert.Equal(t, c.output, lc_0206_reverse_linked_list_iter(NewList(c.input)).String())
	}
}

func lc_0206_reverse_linked_list_recur(head *ListNode) *ListNode {
	return head.Reverse()
}

// Кратко
// Для того, чтобы "развернуть" связный список, нужно содержать три переменные: предыдущий, текущий и следующий узлы.
// Сохраняем в следующий то, на что указывает текущий, и записываем предыдущий в текущий, затем сдвигаем позиции, пока текущий не будет пуст.
//
// Подробно
// Начальные условия: предыдущий узел пуст, текущий есть исходный корень.
// 1. Проверяем, является ли текущий узел пустым. Если да, то возвращаем предыдущий как последний обработанный (и, следовательно, корень).
// 2. Запоминаем, куда указывает текущий узел (т.е. сохраняем следующий узел).
// 3. Привязываем текущий узел к предыдущему.
// 4. Делаем текущий узел предыдущим и следующий текущим.
// 5. Повторяем процедуру.
func lc_0206_reverse_linked_list_iter(head *ListNode) *ListNode {
	var prev *ListNode
	for {
		if head == nil {
			return prev
		}
		head, head.Next, prev = head.Next, prev, head
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-nth-node-from-end-of-list
func Test_lc_0019_remove_nth_node_from_end_of_list(t *testing.T) {
	cases := []struct {
		input  string
		n      int
		output string
	}{
		{"1,2,3,4", 1, "1,2,3"},
		{"1,2,3,4", 2, "1,2,4"},
		{"1,2,3,4", 4, "2,3,4"},
		{"1", 1, ""},
		{"1,2", 1, "1"},
	}
	for _, c := range cases {
		assert.Equal(t, NewList(c.output).String(), lc_0019_remove_nth_node_from_end_of_list(NewList(c.input), c.n).String())
	}
}

// Итерируем по списку 2 метки — фронт и хвост. Отслеживаем дистанцию между ними. Начинаем двигать хвост, когда дистанция будет такая,
// чтобы хвост указывал на элемент, находящийся перед тем элементом, который нужно удалить.
// Как только фронт достигает конца списка:
// - если не набрали небоходимую для движения дистанцию, то возвращаем узел, следующий за головой (удаляем первый элемент списка).
// - в противном случае связываем хвост с узлом, идущим за следующим узлом (это учитывает случай, когда нужно удалить последний узел списка).
func lc_0019_remove_nth_node_from_end_of_list(head *ListNode, n int) *ListNode {
	for distance, front, tail := 0, head, head; ; distance++ {
		switch {
		case front.Next == nil:
			if distance < n {
				return tail.Next
			}
			tail.Next = tail.Next.Next
			return head
		case distance >= n:
			tail = tail.Next
		}
		front = front.Next
	}
}

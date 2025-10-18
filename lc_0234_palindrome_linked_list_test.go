package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/palindrome-linked-list
func Test_lc_0234_palindrome_linked_list(t *testing.T) {
	cases := []struct {
		head   string
		result bool
	}{
		{"1,2,3,2,1", true},
		{"1,2,2,1", true},
		{"1,2,2", false},
		{"1", true},
		{"1,2,1", true},
		{"1,2", false},
		{"1,3,2,4,3,2,1", false},
		{"1,2,1,3", false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0234_palindrome_linked_list(NewList(c.head)))
		assert.Equal(t, c.result, lc_0234_palindrome_linked_list_2(NewList(c.head)))
	}
}

// 100/83
// Первым проходом узнаём длину связного списка.
// Вторым проходом в левой половине прибавляем к счётчику произведение значения и квадрата расстояния до середины,
// в правой — то же самое, только убавляем. В итоге, если список есть палиндром, то счётчик станет равным нулю.
// Квадрат нужен для избежания коллизий.
func lc_0234_palindrome_linked_list(head *ListNode) bool {
	length, counter := 1, 0
	for node := head.Next; node != nil; node, length = node.Next, length+1 {
	}
	for i := 1; head != nil; head, i = head.Next, i+1 {
		switch {
		case i <= length/2:
			counter += head.Val * i * i
		case length%2 == 1 && i == 1+length/2:
		case i >= length/2:
			counter -= head.Val * (length - i + 1) * (length - i + 1)
		}
	}
	return counter == 0
}

// 62/83
// Сначала находим середину методом черепахи и зайца.
// Разворачиваем первую половину списка.
// Последовательно сравниваем получившийся список и вторую половину.
func lc_0234_palindrome_linked_list_2(head *ListNode) bool {
	a, b := head, head.Next
	for {
		switch {
		case b == nil:
		case b.Next == nil:
			b = a.Next
		case b.Next.Next == nil:
			b = a.Next.Next
		default:
			a, b = a.Next, b.Next.Next
			continue
		}
		break
	}
	var prev *ListNode
	for head != a {
		prev, head, head.Next = head, head.Next, prev
	}
	for a.Next = prev; b != nil; a, b = a.Next, b.Next {
		if a.Val != b.Val {
			return false
		}
	}
	return true
}

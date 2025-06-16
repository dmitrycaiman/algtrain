package main

import (
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
	assert.True(t, lc_0002_checkEquality(l3, lc_0002_add_two_numbers(l1, l2)))
	assert.True(t, lc_0002_checkEquality(l5, lc_0002_add_two_numbers(l3, l4)))
	assert.True(t, lc_0002_checkEquality(l3, lc_0002_add_two_numbers(nil, l3)))
	assert.True(t, lc_0002_checkEquality(l5, lc_0002_add_two_numbers(l5, nil)))
	assert.Nil(t, lc_0002_add_two_numbers(nil, nil))
}

// Производим поразрядное сложение соответствующих значений списов, как "в столбик".
func lc_0002_add_two_numbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{} // Корневой узел результата
	current := result     // Текущий заполняемый узел результата
	overflow := false     // Флаг переполнения предыдущего разряда
	a, b := 0, 0          // Слагаемые текущего разряда
	for {
		switch { // Проверяем на пустоту следующие соответствующие узлы. Если они непустые, то берём их значения как слагаемые.
		case l1 != nil && l2 != nil:
			a, b = l1.Val, l2.Val
			l1, l2 = l1.Next, l2.Next
		case l1 != nil && l2 == nil:
			a, b = l1.Val, 0
			l1 = l1.Next
		case l1 == nil && l2 != nil:
			a, b = 0, l2.Val
			l2 = l2.Next
		case l1 == nil && l2 == nil: // Если дальше "пути нет" (оба узла пустые), то учитываем флаг переполнения и выходим с результатом.
			if overflow {
				current.Next = &ListNode{Val: 1}
			}
			return result.Next
		}
		current.Next = &ListNode{}
		current = current.Next
		if overflow { // Если в предыдущем разряде было переполнение, то учитываем это в текущем значении одного из слагаемых.
			overflow = false
			a++
		}
		current.Val = a + b
		if current.Val >= 10 { // Если случилось переполнение текущего разряда, то запоминаем это и помещаем в текущее значение только остаток.
			overflow = true
			current.Val -= 10
		}
	}
}

func lc_0002_checkEquality(l1, l2 *ListNode) bool {
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

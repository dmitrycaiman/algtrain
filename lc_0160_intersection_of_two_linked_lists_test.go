package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/intersection-of-two-linked-lists
func Test_lc_0160_intersection_of_two_linked_lists(t *testing.T) {
	t.Run(
		"general case",
		func(t *testing.T) {
			list1, tail1 := NewListWithTail("1,2,3,4,5")
			list2, tail2 := NewListWithTail("1,2,3")
			list3 := NewList("1,2,3,4,5,6")
			tail1.Next = list3
			tail2.Next = list3

			assert.Equal(t, list3, lc_0160_intersection_of_two_linked_lists_1(list1, list2))
			assert.Equal(t, list3, lc_0160_intersection_of_two_linked_lists_2(list1, list2))
		},
	)
	t.Run(
		"equal tail",
		func(t *testing.T) {
			list1, tail1 := NewListWithTail("1,2,3,4,5")
			list2, tail2 := NewListWithTail("1,2,3")
			list3 := NewList("1")
			tail1.Next = list3
			tail2.Next = list3

			assert.Equal(t, list3, lc_0160_intersection_of_two_linked_lists_1(list1, list2))
			assert.Equal(t, list3, lc_0160_intersection_of_two_linked_lists_2(list1, list2))
		},
	)
	t.Run(
		"no intersection",
		func(t *testing.T) {
			list1 := NewList("1,2,3,4,5")
			list2 := NewList("1,2,3")

			assert.Nil(t, lc_0160_intersection_of_two_linked_lists_1(list1, list2))
			assert.Nil(t, lc_0160_intersection_of_two_linked_lists_2(list1, list2))
		},
	)
}

// Два указателя начинают проход с голов списоков. На каждой итерации происходит следующий переход:
// - если следующий узел существует, то переходим на него;
// - иначе переходим на голову противоположного списка относительно изначального (т.е. в изначальный список уже не возвращаемся).
// Таким образом, когда будут осуществлены оба перехода на противоположные списки, то оба указателя будут на одинаковом расстоянии
// от узла пересечения. Возвращаем один из указателей, когда оба указателя будут равны.
//
// Также при отсутствии узла пересечения обязательно возникнет ситуация, когда оба указателя равны nil. Он и будет возвращён.
//
// // Пример при наличии узла пересечения:
// //
// // 1. Указатели a и b на головах списков A и B.
// // --------------
// //    a
// // A: o-o-o-o-o-o
// // B:     o-o-o
// //        b
// // --------------
// // 2. Сделали 4 итерации: указатель b прошёл до конца списка B и перешёл на голову списка A.
// // --------------
// //    b       a
// // A: o-o-o-o-o-o
// // B:     o-o-o
// // --------------
// // 3. Сделали 6 итераций: указатель a прошёл до конца списка A и перешёл на голову списка B.
// // Теперь оба указателя на одинаковом расстоянии от точки пересечения.
// // --------------
// //        b
// // A: o-o-o-o-o-o
// // B:     o-o-o
// //        a
// // --------------
// // 3. Сделали 8 итераций: пришли в точку пересечения.
// // --------------
// //            b
// // A: o-o-o-o-o-o
// // B:     o-o-o
// // 			 a
// // --------------
func lc_0160_intersection_of_two_linked_lists_1(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}
	return l1
}

// Собираем элементы обоих список в хеш-таблицы. На каждой итерации проверяем, не встречалмя ли узел в таблице другого списка.
func lc_0160_intersection_of_two_linked_lists_2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := map[*ListNode]struct{}{}, map[*ListNode]struct{}{}
	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := b[headA]; ok {
				return headA
			}
			a[headA] = struct{}{}
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := a[headB]; ok {
				return headB
			}
			b[headB] = struct{}{}
			headB = headB.Next
		}
	}
	return nil
}

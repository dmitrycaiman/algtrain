package main

import (
	"testing"
)

// https://leetcode.com/problems/rotate-list
func Test_lc_0061_rotate_list(t *testing.T) {

}

// 0ms (100%), 4.33MB (81.01%)
// Вычисляем полную длину списка прямым проходом. В процессе этого получаем последний элемент.
// Вычисляем остаток от деления k на длину, что есть число вращений без циклов.
// Вычисляем прямым проходом элемент, стоящий перед элементом, порядковый номер которого равен числу вращений.
// Отрезаем список на этом элементе и склеиваем оставшуюся часть с началом.
func lc_0061_rotate_list(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	l, last, cut := 1, head, head
	for last.Next != nil {
		l++
		last = last.Next
	}
	if k%l == 0 {
		return head
	}
	for range l - k%l - 1 {
		cut = cut.Next
	}
	result := cut.Next
	cut.Next = nil
	last.Next = head
	return result
}

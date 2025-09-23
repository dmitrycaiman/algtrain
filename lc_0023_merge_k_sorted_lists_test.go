package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: написать тесты.
func Test_lc_0023_merge_k_sorted_lists(t *testing.T) {
	cases := []struct {
		lists  []string
		result string
	}{
		{[]string{"1,2,3", "0,1,3,4", "-5,-4,0,9", "1"}, "-5,-4,0,0,1,1,1,2,3,3,4,9"},
	}
	for _, c := range cases {
		lists := []*ListNode{}
		for _, list := range c.lists {
			lists = append(lists, NewList(list))
		}
		assert.Equal(t, c.result, lc_0023_merge_k_sorted_lists(lists).String())
	}
}

// Имеем слайс голов списков. Среди этих голов ищем самое маленькое значение проходом по слайсу.
// После этого проходимся по слайсу ещё раз и каждую голову с этим значением присоединяем к результирующему списку.
// При этом на место головы в слайсе ставим её потомка. Заканчиваем вычисления, когда весь список голов состоит из nil-элементов.
func lc_0023_merge_k_sorted_lists(lists []*ListNode) *ListNode {
	const inf = 10_001
	result := &ListNode{}
	last := result
	for {
		smallest := inf
		for _, v := range lists {
			if v != nil && v.Val < smallest {
				smallest = v.Val
			}
		}
		if smallest == inf {
			break
		}
		for i := range lists {
			if lists[i] != nil && lists[i].Val == smallest {
				last.Next, last = lists[i], lists[i]
				lists[i] = lists[i].Next
			}
		}
	}
	return result.Next
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/min-stack
func Test_lc_0155_min_stack(t *testing.T) {
	cases := []struct {
		actions func(s *MinStack)
		lastMin int
	}{
		{
			func(s *MinStack) {
				s.Push(-2)
				s.Push(0)
				s.Push(-3)
				s.GetMin()
				s.Pop()
				s.Top()
				s.GetMin()
			},
			-2,
		},
		{
			func(s *MinStack) {
				s.Push(-2)
				s.Push(0)
				s.Push(-3)
				s.GetMin()
				s.Pop()
				s.Top()
				s.GetMin()
				s.Push(-4)
			},
			-4,
		},
		{
			func(s *MinStack) {
				s.Push(-2)
				s.Pop()
				s.Push(-3)
				s.Pop()
				s.Push(3)
				s.Pop()
				s.Push(4)
				s.Push(5)
				s.Push(6)
				s.Push(-1)
				s.GetMin()
				s.Pop()
				s.Pop()
				s.Top()
			},
			4,
		},
	}
	for _, c := range cases {
		assert.Equal(t, c.lastMin, lc_0155_min_stack(c.actions))
	}
}

// 1ms, (89.59%), 9.40MB (34.12%)
// Стек храним в предаллоцированном массиве, вершину держим в памяти.
// Отношения "больше-меньше" между элементами храним в связном списке, где корень есть наименьший элемент.
// На Pop удаляем вершину стека и правим связный список (соединяем разрыв или приводим в порядок края).
// На Top отдаём вешину стека.
// На GetMin отдаём корень связного списка.
// На Push берём вершину стека и от неё ищем в связном списке место вставки (самое долгое). Как только нашли, вставляем и обновляем вершину.
func lc_0155_min_stack(actions func(s *MinStack)) int {
	s := NewMinStack()
	actions(&s)
	return s.min.val
}

type minStackElement struct {
	back, next *minStackElement
	val        int
}

type MinStack struct {
	stack []*minStackElement
	min   *minStackElement
	pos   int
}

func NewMinStack() MinStack { return MinStack{stack: make([]*minStackElement, 30000), pos: -1} }

func (slf *MinStack) Push(val int) {
	if slf.pos == -1 {
		slf.pos++
		e := &minStackElement{val: val}
		slf.stack[0] = e
		slf.min = e
		return
	}
	e := &minStackElement{val: val}
	tmp := slf.stack[slf.pos]
	if tmp.val > val {
		for e.next == nil {
			switch {
			case tmp.val >= val && tmp.back != nil:
				tmp = tmp.back
			case tmp.val >= val && tmp.back == nil:
				e.next = tmp
				tmp.back = e
				slf.min = e
			default:
				e.next = tmp.next
				e.back = tmp
				e.next.back = e
				e.back.next = e
			}
		}
	} else {
		for e.back == nil {
			switch {
			case tmp.val <= val && tmp.next != nil:
				tmp = tmp.next
			case tmp.val <= val && tmp.next == nil:
				e.back = tmp
				tmp.next = e
			default:
				e.next = tmp
				e.back = tmp.back
				e.next.back = e
				e.back.next = e
			}
		}
	}
	slf.pos++
	slf.stack[slf.pos] = e
}

func (slf *MinStack) Pop() {
	tmp := slf.stack[slf.pos]
	switch {
	case tmp.back != nil && tmp.next != nil:
		tmp.back.next = tmp.next
		tmp.next.back = tmp.back
	case tmp.back != nil:
		tmp.back.next = nil
	case tmp.next != nil:
		slf.min = tmp.next
		tmp.next.back = nil
	}
	slf.stack[slf.pos] = nil
	slf.pos--
}

func (slf *MinStack) Top() int { return slf.stack[slf.pos].val }

func (slf *MinStack) GetMin() int { return slf.min.val }

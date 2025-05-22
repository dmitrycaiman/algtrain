package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0020_valid_parentheses(t *testing.T) {
	cases := []struct {
		input  string
		result bool
	}{
		{"()", true}, {"(]", false}, {"())", false}, {")", false}, {"((", false}, {"{[]{{{()}}}()}", true},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0020_valid_parentheses(c.input))
	}
}

// Итерируемся по строке. Складываем открывающие скобки в стек. Если встречаем закрывающую скобку, то извлекаем элемент из стека: если он
// не соответствует ей по виду, делаем вывод о невалидности строки. В конце проверяем размер стека: он должен быть пуст, что значит, что
// каждая открывающая скобка нашла "свою" закрывающую.
func lc_0020_valid_parentheses(s string) bool {
	var stack []rune
	pop := func() rune {
		l := len(stack)
		if l == 0 {
			return '_'
		}
		output := stack[l-1]
		stack = stack[:l-1]
		return output
	}
	push := func(input rune) {
		stack = append(stack, input)
	}
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			push(v)
		default:
			if m[pop()] != v {
				return false
			}
		}
	}
	return len(stack) == 0
}

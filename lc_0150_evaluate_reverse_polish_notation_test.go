package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0150_evaluate_reverse_polish_notation(t *testing.T) {
	assert.Equal(t, 2, lc_0150_evaluate_reverse_polish_notation([]string{"2"}))
	assert.Equal(t, 3, lc_0150_evaluate_reverse_polish_notation([]string{"1", "2", "+"}))
	assert.Equal(t, -4, lc_0150_evaluate_reverse_polish_notation([]string{"1", "2", "3", "+", "-"}))
	assert.Equal(t, -5, lc_0150_evaluate_reverse_polish_notation([]string{"5", "1", "2", "*", "3", "-", "*"}))

}

// Итерируемся по массиву токенов. Если встречаем число, то складываем его в стек. Если встречаем оператор, то извлекаем из стека последние
// два значения и производим над ними действие оператора. При чём правым операндом будет первое извлечённое значение, левым — следующее.
// Таким образом любое верное выражение в виде польской обратной нотации будет вычислено.
func lc_0150_evaluate_reverse_polish_notation(tokens []string) int {
	s := []int{}
	pop := func() int {
		if len(s) == 0 {
			return 0
		}
		output := s[len(s)-1]
		s = s[:len(s)-1]
		return output
	}
	push := func(input int) { s = append(s, input) }
	solve := func(a, b int, op string) int {
		switch op {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			return a / b
		}
		return 0
	}
	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			b, a := pop(), pop()
			push(solve(a, b, v))
		default:
			n, _ := strconv.Atoi(v)
			push(n)
		}
	}
	return s[0]
}

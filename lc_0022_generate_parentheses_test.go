package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0022_generate_parentheses(t *testing.T) {
	assert.ElementsMatch(t, []string{"()"}, lc_0022_generate_parentheses(1))
	assert.ElementsMatch(t, []string{"()()", "(())"}, lc_0022_generate_parentheses(2))
	assert.ElementsMatch(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, lc_0022_generate_parentheses(3))
}

// Задача состоит в том, чтобы расставить закрывающие скобки между N открывающих. Ограничения:
// 1. После N-ой открывающей скобки может стоять от 0 до N закрывающих.
// 2. Особый случай: после последней открывающей скобки может стоять от 1 до N закрывающих.
// Формируем массивы, отображающие расположение закрывающих скобок в формате {c1, c2, ..., cN},
// где ci — количество закрывающих скобок после соответствующей открывающей. На основе массива формируем конечный результат.
func lc_0022_generate_parentheses(n int) []string {
	s := [][]int{{0}}
	// i — позиция скобки.
	for i := 1; i <= n; i++ {
		l := len(s)
		// j — выбор слайса.
		for j := range l {
			// k — число, которое будем прибавлять.
			if i != n {
				for k := 0; k <= i; k++ {
					c := make([]int, len(s[j]))
					copy(c, s[j])
					if c[0]+k > i {
						break
					}
					c[0] += k
					s = append(s, append(c, k))
				}
			} else {
				for k := 1; k <= i; k++ {
					c := make([]int, len(s[j]))
					copy(c, s[j])
					if c[0]+k != i {
						continue
					}
					c[0] += k
					s = append(s, append(c, k))
				}
			}
		}
		s = s[l:]
	}
	res := []string{}
	// i — выбор слайса.
	for i := range len(s) {
		v := ""
		// j — выбор позиции в слайсе.
		for j := 1; j < len(s[i]); j++ {
			v += "("
			// k — счётчик закрывающих скобок.
			for k := 0; k < s[i][j]; k++ {
				v += ")"
			}
		}
		res = append(res, v)
	}
	return res
}

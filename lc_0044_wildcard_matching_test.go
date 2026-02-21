package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/wildcard-matching
func Test_lc_0044_wildcard_matching(t *testing.T) {
	cases := []struct {
		s, p   string
		result bool
	}{
		{"aaa", "aaa*", true},
		{"aaa", "*a*a*a*", true},
		{"aaa", "aa?", true},
		{"aaa", "??", false},
		{"aaa", "*", true},
		{"aaa", "****", true},
		{"acccdcccc", "a*c*d*c???", true},
		{"a", "**?**", true},
		{"ab", "*a", false},
		{"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb", false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0044_wildcard_matching(c.s, c.p))
	}
}

// Ставим метки на начала строк и итерируемся.
// Если встречаем одинаковые буквы или знак "?", то переходим к следующим.
// Если встречаем "*", то:
//   - пробегаем в шаблоне дальше до следующей буквы (т.е. пропускаем идущие подряд "*"),
//     при этом если дошли до конца шаблона, то выходим с успехом, т.к. символ "*" стоял в конце.
//   - сравниваем букву строки с найденной буквой шаблона: если они совпадают, то пробуем рекурсивно решить такую подзадачу,
//     где действие "*" действительно закончилось на этом месте.
//   - если подзадача не была решена, то запоминаем неудачный результат подзадачи в хранилище (мемоизация) и пытаемся найти следующую букву строки,
//     совпадающую с буквой шаблона и так же пытаемся решить подзадачу.
//   - если подзадача была удачно решена, выходим с успехом.
//   - если не удалось найти подходящих букв строки для решения подзадачи, выходим с неудачей.
//   - при решении других подзадач сначала ищем решение в хранилище (убираем дублирование вычислений).
//
// Если получилось проитерироваться по всей строке и всему шаблону, выходим с успехом.
func lc_0044_wildcard_matching(s, p string) bool {
	return lc_0044(s, p, 0, 0, map[[2]int]struct{}{})
}

func lc_0044(s, p string, i, j int, storage map[[2]int]struct{}) bool {
	for i < len(s) && j < len(p) {
		switch {
		case s[i] == p[j] || p[j] == '?':
			i, j = i+1, j+1
		case p[j] == '*':
			for ; j < len(p); j++ {
				if p[j] != '*' {
					break
				}
			}
			if j == len(p) {
				return true
			}
			for ; i < len(s); i++ {
				if s[i] == p[j] || p[j] == '?' {
					_, ok := storage[[2]int{i, j}]
					switch {
					case !ok && lc_0044(s, p, i+1, j+1, storage):
						return true
					case !ok:
						storage[[2]int{i, j}] = struct{}{}
					}
				}
			}
			fallthrough
		default:
			return false
		}
	}
	for ; j < len(p); j++ {
		if p[j] != '*' {
			break
		}
	}
	return i == len(s) && j == len(p)
}

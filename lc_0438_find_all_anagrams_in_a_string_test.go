package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-all-anagrams-in-a-string
func Test_lc_0438_find_all_anagrams_in_a_string(t *testing.T) {
	cases := []struct {
		s, p   string
		result []int
	}{
		{"aabaaba", "aab", []int{0, 1, 2, 3, 4}},
		{"aabcaaba", "aab", []int{0, 4, 5}},
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abcacba", "abc", []int{0, 1, 3, 4}},
		{"a", "a", []int{0}},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0438_find_all_anagrams_in_a_string(c.s, c.p))
	}
}

// Формируем хеш-таблицу частот, которая говорит о количестве вхождений символов анаграммы в текующую подстроку.
// Ключ таблицы — символ, значение есть пара [i,j], где i — количество вхождений, j — необходимое количество вхождений
// (в анаграмму символ может входить несколько раз).
// Двигаем по строке окно длиной искомой анаграммы.
// Для символа, оставленного слева, уменьшаем количесво вхождений (но не менее, чем до нуля).
// Для символа, настигнутого справа, увеличиваем количество вхождений.
// Если символ не входит в анаграмму, ничего не делаем.
// После каждого движения окна проверяем таблицу частот: если i == j, то добавляем начапо подстроки в результат.
func lc_0438_find_all_anagrams_in_a_string(s string, p string) []int {
	result, symbols := []int{}, map[byte][]int{}
	for i := range p {
		if _, ok := symbols[p[i]]; !ok {
			symbols[p[i]] = make([]int, 2)
		}
		symbols[p[i]][1]++
	}
	check := func() bool {
		for _, v := range symbols {
			if v[0] != v[1] {
				return false
			}
		}
		return true
	}
	for i, j := -len(p), 0; j < len(s); i, j = i+1, j+1 {
		sym, ok := symbols[s[j]]
		if ok {
			sym[0]++
		}
		if i >= 0 {
			sym, ok = symbols[s[i]]
			if ok && sym[0] >= 0 {
				sym[0]--
			}
		}
		if check() {
			result = append(result, i+1)
		}
	}
	return result
}

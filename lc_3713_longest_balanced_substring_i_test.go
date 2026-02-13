package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-balanced-substring-i
func Test_lc_3713_longest_balanced_substring_i(t *testing.T) {
	cases := []struct {
		input  string
		result int
	}{
		{"abbac", 4},
		{"sdgfsgskokokokoxcgdghzzlpkolkolkolkolafafqcvcfh", 12},
		{"aaaaaaaaaaaaaaabcccccccccccccccczccccccz", 16},
		{"ab", 2},
		{"a", 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_3713_longest_balanced_substring_i(c.input))
	}
}

// Полный перебор всех возможных подстрок с проверкой каждой на сбалансированность.
// Включены эвристики пропуска таких подстрок, которые заведомо короче текущего результата.
// В качестве хеш-таблицы используем массив длиной 26, по количеству английских строчных букв.
func lc_3713_longest_balanced_substring_i(s string) int {
	max := 0
	for i := range s {
		b := [26]int{}
		if max >= len(s)-i { // Нет смысла смотреть меньшие подстроки.
			break
		}
		for j := i; j < len(s); j++ { // Сначала собираем информацию обо всей подстроке.
			b[s[j]-97]++
		}
	loop:
		for j := len(s); j > i; j-- { // Убираем по одному символу.
			if j != len(s) {
				b[s[j]-97]--
				if max >= j-i { // Нет смысла смотреть меньшие подстроки.
					break
				}
			}
			v1 := b[s[i]-97]
			for _, v2 := range b { // Проверяем на баланс.
				if v2 != 0 && v2 != v1 {
					continue loop
				}
			}
			if j-i > max { // Обновляем результат, если нужно.
				max = j - i
			}
		}
	}
	return max
}

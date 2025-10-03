package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-repeating-character-replacement
func Test_lc_0424_longest_repeating_character_replacement(t *testing.T) {
	cases := []struct {
		s         string
		k, result int
	}{
		{"AAA", 1, 3},
		{"AAABBBB", 1, 5},
		{"AAABBAABBB", 2, 7},
		{"AAABBAABBBB", 2, 8},
		{"ABCABCABCCC", 2, 6},
		{"AABABBACD", 1, 4},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0424_longest_repeating_character_replacement(c.s, c.k))
	}
}

// Скользящее окно. Формируем набор уникальных символов входной строки и применяем алгоритм для кажого из них.
// Суть: расширяем окно вправо и следим за тем, сколько в нём символов, не равных ключевому.
// Как только символов станет максимально возможное количество, окно окажется на грани валидности.
// Такое окно может расширяться только через ключевые символы. При встрече неключевого символа двигаем левую границу до тех пор,
// пока число неключевых символов в подстроке станет равно максимально возможному.
// На каждой итерации, когда окно валидно или находится на грани, считаем полученную длину. Возвращаем максимальную длину после всех проходов.
func lc_0424_longest_repeating_character_replacement(s string, k int) int {
	syms := map[byte]struct{}{}
	for i := range s {
		syms[s[i]] = struct{}{}
	}
	maxLength := 0
	for sym := range syms {
		wrongSymCounter := 0
		for i, j := 0, 0; j < len(s); {
			switch {
			case s[j] == sym:
			case wrongSymCounter == k:
				for {
					if s[i] != sym {
						wrongSymCounter--
						i++
						break
					}
					i++
				}
				continue
			case s[j] != sym:
				wrongSymCounter++
			}
			if length := j - i + 1; length > maxLength {
				maxLength = length
			}
			j++

		}
	}
	return maxLength
}

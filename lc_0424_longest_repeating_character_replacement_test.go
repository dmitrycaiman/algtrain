package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func lc_0424_longest_repeating_character_replacement(s string, k int) int {
	m, keys := map[byte]int{}, []byte{}
	for i := range s {
		if _, ok := m[s[i]]; !ok {
			keys = append(keys, s[i])
		}
		m[s[i]]++
	}
	slices.SortFunc(
		keys,
		func(a, b byte) int {
			switch {
			case m[a] < m[b]:
				return 1
			case m[a] > m[b]:
				return -1
			default:
				return 0
			}
		},
	)
	maxLength := 0
	for _, key := range keys {
		if m[key]+k <= maxLength {
			break
		}
		counter, length := k, 0
		for i, j := 0, 0; j < len(s); j++ {
			if s[j] != key {
				counter--
			}
			if counter == -1 {
				if length > maxLength {
					maxLength = length
				}
				if s[i] == key {
					for s[i] == key {
						i++
						length--
					}
				}
				i++
				length--
				counter++
			}
			length++
		}
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

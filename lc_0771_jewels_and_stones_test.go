package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/jewels-and-stones
func Test_lc_0771_jewels_and_stones(t *testing.T) {
	cases := []struct {
		jewels, stones string
		result         int
	}{
		{"abcABC", "AAABBVVGGDDcccab", 10},
		{"abcC", "VVGGDD", 0},
		{"a", "qqqqqqqqppppppalllfffffbnghj", 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0771_jewels_and_stones(c.jewels, c.stones))
	}
}

// Набираем хеш-таблицу драгоценностей и проверяем каждый элемент камня на то, что он в ней есть. Считаем их сумму.
func lc_0771_jewels_and_stones(jewels string, stones string) int {
	m, counter := make(map[byte]struct{}, len(jewels)), 0
	for i := range jewels {
		m[jewels[i]] = struct{}{}
	}
	for i := range stones {
		if _, ok := m[stones[i]]; ok {
			counter++
		}
	}
	return counter
}

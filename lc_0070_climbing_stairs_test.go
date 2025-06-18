package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/climbing-stairs/
func Test_lc_0070_climbing_stairs(t *testing.T) {
	cases := []struct{ input, result int }{{3, 3}, {4, 5}}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0070_climbing_stairs_memo(c.input))
	}
}

// Решаем рекурсивным перебором с мемоизацией через кеш уже посчитанных результатов.
// На каждом этапе рекурсивно разветвляемся, делая шаг на 1 или 2 ступеньки.
// Рекурсия останавливается, когда мы достигли целевого значения (возвращаем 1) или промахнулись (0).
// Каждый удачно пройденный путь даёт единицу, которые по выходу из рекурсий будут сложены в результат.
func lc_0070_climbing_stairs_memo(n int) int {
	cache := make([]int, n+1)
	for i := range cache {
		cache[i] = -1
	}
	return lc_0070_climbing_stairs_memo_r(n, cache)
}

// Пример перебора до 4:
//	      1               2
//	  2       3        3     4
//	3   4   4   5    4   5
// 4 5

func lc_0070_climbing_stairs_memo_r(n int, cache []int) (res int) {
	if cache[n] != -1 {
		return cache[n]
	}
	a, b := 0, 0
	defer func() { cache[n] = res }()
	switch {
	case n-1 == 0:
		a = 1
	case n-1 > 0:
		a = lc_0070_climbing_stairs_memo_r(n-1, cache)
	}
	switch {
	case n-2 == 0:
		b = 1
	case n-1 > 0:
		b = lc_0070_climbing_stairs_memo_r(n-2, cache)
	}
	return a + b
}

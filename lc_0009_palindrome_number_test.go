package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/palindrome-number
func Test_lc_0009_palindrome_number(t *testing.T) {
	cases := []struct {
		input  int
		output bool
	}{
		{121, true},
		{122, false},
		{111, true},
		{10, false},
		{11, true},
		{101, true},
		{12221, true},
		{1221, true},
		{1021, false},
		{1000000000001, true},
		{1000000000021, false},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0009_palindrome_number(c.input))
	}
}

// 0ms (100%), 5.96MB (98.46%)
// Вычисляем значения высшего и низшего разрядов. Если они не равны, выходим с неудачей.
// Если равны, от отрезаем проверенные разряды с каждой стороны и продолжаем, пока число не "кончится".
func lc_0009_palindrome_number(x int) bool {
	if x < 0 {
		return false
	}
	hi := 1
	for x/hi >= 10 {
		hi *= 10
	}
	for hi >= 1 {
		if x/hi != x%10 {
			return false
		}
		x = x % hi / 10
		hi /= 100
	}
	return true
}

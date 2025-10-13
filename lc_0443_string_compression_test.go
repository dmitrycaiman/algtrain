package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/string-compression
func Test_lc_0443_string_compression(t *testing.T) {
	cases := []struct {
		chars, result string
		length        int
	}{
		{"abbccc", "ab2c3", 5},
		{"aaaaaaaaaaacccccccccc", "a11c10", 6},
		{"abcdefg", "abcdefg", 7},
		{"a", "a", 1},
		{"aaaaaaaaaaaaaaaaaaaa", "a20", 3},
	}
	for _, c := range cases {
		chars := []byte(c.chars)
		assert.Equal(t, c.length, lc_0443_string_compression(chars))
		assert.Equal(t, []byte(c.result), chars[:c.length])
	}
}

// 0/100, 5/39
// Два указателя. Первый указывает на позицию записи результирующей строки, второй идёт по массиву.
// Второй указатель на каждой итерации проверяет, равен ли текущий символ тому, что на первом указателе.
// Если да, то увеличивается счётчик одинаковых символов.
// Если нет, то величина счётчика вносится в результирующую строку (если он не единица),
// а первый указатель сдвигается на количество символов, составляющих счётчик,
// и ещё один раз, записывая новый символ. Счётчик далее сбрасывается к единице. Так до конца строки.
func lc_0443_string_compression(chars []byte) int {
	i := 0
	for j, c := 1, 1; j <= len(chars); j++ {
		if j == len(chars) || chars[j] != chars[i] {
			if c > 1 {
				for _, v := range strconv.Itoa(c) {
					i++
					chars[i] = byte(v)
				}
			}
			if j == len(chars) {
				break
			}
			i++
			chars[i], c = chars[j], 1
		} else {
			c++
		}
	}
	return i + 1
}

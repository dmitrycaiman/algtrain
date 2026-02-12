package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/multiply-strings
func Test_lc_0043_multiply_strings(t *testing.T) {
	cases := []struct {
		num1, num2, result string
	}{
		{"23", "23", "529"},
		{"123", "456", "56088"},
		{"1234", "0", "0"},
		{"0", "1234", "0"},
		{"0", "0", "0"},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0043_multiply_strings(c.num1, c.num2))
	}
}

// Умножаем "в столбик".
func lc_0043_multiply_strings(num1, num2 string) string {
	result := []byte{48}
	for i := range num2 {
		if mul := lc_0043_mul([]byte(num1), num2[len(num2)-1-i]); mul != nil {
			for range i {
				mul = append(mul, 48)
			}
			result = lc_0043_add(result, mul)
		}
	}
	return string(result)
}

// Умножение чисел "по месту": результат стараемся отдать в исходном массиве.
// Если любой из множителей 0, возвращаем nil. Умножаем "в столбик": считаем произведения поразрядно от меньшего,
// сохраняя величину переполнения. Если разряды кончились и осталось переполнение, то вставляем его "вперёд".
func lc_0043_mul(a []byte, b byte) []byte {
	if b -= 48; b == 0 || a[0] == 48 {
		return nil
	}
	var overflow, mul byte
	for i := len(a) - 1; i >= 0; i-- {
		mul = (a[i]-48)*b + overflow
		a[i], overflow = mul%10+48, mul/10
	}
	if overflow != 0 {
		a = append([]byte{overflow + 48}, a...)
	}
	return a
}

// Сложение чисел "по месту": результат стараемся отдать в одном из исходных массивов.
// Берём слагаемые так, чтобы к большему числу прибавлять меньшее. Складываем "в столбик".
// Если индексы меньшего числа закончились, то "в столбик" прибавляем величину оставшегося переполнения.
func lc_0043_add(a, b []byte) []byte {
	if len(a) < len(b) {
		a, b = b, a
	}
	var overflow, sum byte
	for i := 0; ; i++ {
		switch {
		case i < len(b):
			sum = (a[len(a)-1-i] - 48) + (b[len(b)-1-i] - 48) + overflow
		case overflow != 0 && i < len(a):
			sum = (a[len(a)-1-i] - 48) + overflow
		case overflow != 0:
			return append([]byte{overflow + 48}, a...)
		default:
			return a
		}
		a[len(a)-1-i], overflow = sum%10+48, sum/10
	}
}

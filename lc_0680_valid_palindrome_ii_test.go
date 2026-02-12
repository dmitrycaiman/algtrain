package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-palindrome-ii
func Test_lc_0680_valid_palindrome_ii(t *testing.T) {
	cases := []struct {
		input  string
		result bool
	}{
		{"abkcckbta", true},
		{"abkcckbtta", false},
		{"abkcckba", true},
		{"abca", true},
		{"kacabbaak", true},
		{"qcuucuq", true},
		{"axbcbaba", false},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0680_valid_palindrome_ii(c.input))
	}
}

// Обычная проверка на палиндром: идём двумя метками с концов к центру и проверяем равенство символов.
// Также вводится "право на ошибку": один раз при неуспешном сравнении можно сдвинуть одну из меток далее по ходу её движения,
// если это приведёт к устранению ошибки. Если при этом дальнейшие проверки всё равно неуспешны,
// можно попробовать на месте самой первой ошибки сдвинуть уже другую метку в её направлении (если это устраняет ошибку) и попробовать заново.
func lc_0680_valid_palindrome_ii(s string) bool {
	failed, retry, ir, jr := false, false, 0, 0
	for i, j := 0, len(s)-1; i < j; {
		switch {
		case s[i] == s[j]:
			i, j = i+1, j-1
		case failed && !retry: // Пробуем сделать сдвиг в другом направлении.
			i, j, retry = ir, jr, true
		case failed && retry: // Не помогло ни право на ошибку, ни повторная попытка.
			return false
		default: // Сюда мы попадаем единожды, в качестве "права на ошибку".
			switch {
			case j-i == 1: // Первая ошибка возникла у соседних элементов: выходим с успехом, т.к. можно удалить любой из них.
				return true
			case s[i+1] == s[j] && s[i] == s[j-1]: // Сдвиг возможен в обоих направлениях. Можно будет попробовать пройтись повторно.
				i, j, ir, jr = i+2, j-1, i+1, j-2
			case s[i+1] == s[j]: // Сдвиг возможен только левой метки.
				i, j, retry = i+2, j-1, true
			case s[i] == s[j-1]: // Сдвиг возможен только правой метки.
				i, j, retry = i+1, j-2, true
			default: // Сдвиг невозможен.
				return false
			}
			failed = true // Больше нет права на ошибку.
		}
	}
	return true
}

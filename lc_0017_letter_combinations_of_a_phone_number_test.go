package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/letter-combinations-of-a-phone-number
func Test_lc_0017_letter_combinations_of_a_phone_number(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{"", []string{}},
		{"2", []string{"b", "a", "c"}},
		{"25", []string{"bj", "aj", "cj", "bl", "al", "cl", "bk", "ak", "ck"}},
		{"259", []string{
			"bjw", "ajw", "cjw", "blw", "alw", "clw", "bkw", "akw", "ckw",
			"bjx", "ajx", "cjx", "blx", "alx", "clx", "bkx", "akx", "ckx",
			"bjy", "ajy", "cjy", "bly", "aly", "cly", "bky", "aky", "cky",
			"bjz", "ajz", "cjz", "blz", "alz", "clz", "bkz", "akz", "ckz",
		},
		},
	}
	for _, c := range cases {
		assert.ElementsMatch(t, c.output, lc_0017_letter_combinations_of_a_phone_number(c.input))
	}
}

// Рекурсивный перебор. Функция рекурсии получает в аргументах болванку, т.е. постепенно заполняющийся результат.
// На каждом этапе рекурсии болванка копируется, дополняется символами с клавиш текущей позиции и переходит на позицию выше, т.е. следующую клавишу.
// Когда длина болванки достигла длины входых данных, то болванка помещается в хранилище, и рекурсия прекращается.
// В итоге все комбинации будут перебраны и помещены в хранилище, которе возвращается как результат.
func lc_0017_letter_combinations_of_a_phone_number(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return lc_0017(
		nil,
		digits,
		map[byte][]byte{
			'2': {'a', 'b', 'c'},
			'3': {'d', 'e', 'f'},
			'4': {'g', 'h', 'i'},
			'5': {'j', 'k', 'l'},
			'6': {'m', 'n', 'o'},
			'7': {'p', 'q', 'r', 's'},
			'8': {'t', 'u', 'v'},
			'9': {'w', 'x', 'y', 'z'},
		},
		[]string{},
	)
}

func lc_0017(row []byte, set string, m map[byte][]byte, storage []string) []string {
	if l := len(row); l != len(set) {
		for _, v := range m[set[l]] {
			storage = lc_0017(append(row, v), set, m, storage)
		}
		return storage
	}
	return append(storage, string(row))
}

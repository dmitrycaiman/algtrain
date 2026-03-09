package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/convert-a-number-to-hexadecimal
func Test_lc_0405_convert_a_number_to_hexadecimal(t *testing.T) {
	cases := []struct {
		num    int
		result string
	}{
		{15, "f"}, {16, "10"}, {-1, "ffffffff"}, {32, "20"}, {100, "64"}, {255, "ff"},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0405_convert_a_number_to_hexadecimal(c.num))
	}
}

// Последовательно делим на 16 и мапим остатки от деления в символы.
// Если число отрицательное, то инвертируем его противоположное значение и добавляем единицу (представление знака дополнительным кодом).
func lc_0405_convert_a_number_to_hexadecimal(num int) string {
	switch {
	case num == 0:
		return "0"
	case num < 0:
		num = int(^uint32(num*-1)) + 1 // Конвертация в положительное число через дополнительный код.
	}
	m := map[int]byte{0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9', 10: 'a', 11: 'b', 12: 'c', 13: 'd', 14: 'e', 15: 'f'}
	result := []byte{}
	for i := 16; num > 0; {
		result = append(result, m[num%i]) // Собираем символы.
		num /= 16
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 { // Разворачиваем символы, так как собирали в порядке возрастания степени.
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

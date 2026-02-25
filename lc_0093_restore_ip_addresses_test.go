package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/restore-ip-addresses
func Test_lc_0093_restore_ip_addresses(t *testing.T) {
	cases := []struct {
		input   string
		ouptput []string
	}{
		{"127001", []string{"127.0.0.1", "12.70.0.1"}},
		{"1000", []string{"1.0.0.0"}},
		{"10001", []string{"10.0.0.1"}},
	}
	for _, c := range cases {
		assert.ElementsMatch(t, c.ouptput, lc_0093_restore_ip_addresses(c.input))
	}
}

// Итерируемся по строке и откусываем по 1, 2 или 3 символа в качестве значения для октета.
// Формируем октет и рекурсивно обращаемся к остатку строки, следя за длиной и значением октета.
// Валидные результаты сохраняем.
func lc_0093_restore_ip_addresses(s string) []string {
	return lc_0093(0, []byte(s), []byte{}, []string{})
}

const lc_0093_max = int('2')*100 + int('5')*10 + int('5')

func lc_0093(i int, rest, result []byte, storage []string) []string {
	for j := 0; j < 3; j++ { // Сколько символов будем добавлять.
		if len(rest) < 4-i+j { // Если длины строки слишком мало.
			break
		}
		if len(rest) > 10-3*i+j { // Если длины строки слишком много.
			continue
		}
		if j > 0 && rest[0] == '0' { // Нули впереди запрещены.
			break
		}
		if j == 2 && int(rest[0])*100+int(rest[1])*10+int(rest[2]) > lc_0093_max { // Числа не больше 255.
			break
		}
		resultCopy := make([]byte, len(result)) // Копируем строку.
		copy(resultCopy, result)
		switch j { // Добавляем символы.
		case 0:
			resultCopy = append(resultCopy, rest[0])
		case 1:
			resultCopy = append(resultCopy, rest[0], rest[1])
		case 2:
			resultCopy = append(resultCopy, rest[0], rest[1], rest[2])
		}
		if i != 3 { // Если ещё не конец, добавляем точку и рекурсируем.
			resultCopy = append(resultCopy, '.')
			storage = lc_0093(i+1, rest[j+1:], resultCopy, storage)
		} else { // Если конец, то сохраняем результат
			storage = append(storage, string(resultCopy))
		}
	}
	return storage
}

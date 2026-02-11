package main

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/binary-watch
func Test_lc_0401_binary_watch(t *testing.T) {
	cases := []struct {
		turnedOn int
		result   []string
	}{
		{1, []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}},
		{0, []string{}},
	}
	for _, c := range cases {
		// assert.Equal(t, c.result, lc_0401_binary_watch(c.turnedOn))
		lc_0401_binary_watch(c.turnedOn)
	}
}

// Сначала рекурсивно генерируем все возможные комбинации n единиц в рамках 10 позиций.
// По условию 6 позиций на минуты и 4 на часы. В рамках позиций должно быть суммарно ровно n единиц.
// Далее для каждой комбинации формируем время применением маски к первым 6 и последующим 4 позициям.
// Добавляем в результат валидное время (где минуты 0...59 и часы 0...11).
func lc_0401_binary_watch(turnedOn int) []string {
	result := []string{}
	permutations := lc_0401(0, 0, turnedOn, 0, []int{})
	for _, v := range permutations {
		hours := v & 0b1111
		if hours > 11 {
			continue
		}
		minutes := (v & 0b1111110000) >> 4
		if minutes > 59 {
			continue
		}
		result = append(result, fmt.Sprintf("%d:%02d", hours, minutes))
	}
	return result
}

func lc_0401(shifts, cur, max int, result int, storage []int) []int {
	switch {
	case cur == max && shifts == 10:
		return append(storage, result)
	case shifts == 10:
		return storage
	}
	shifts++
	return lc_0401(shifts, cur, max, result<<1, lc_0401(shifts, cur+1, max, result<<1+1, storage))
}

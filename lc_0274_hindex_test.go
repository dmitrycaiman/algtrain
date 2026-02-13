package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/hindex
func Test_lc_0274_hindex(t *testing.T) {
	cases := []struct {
		input  []int
		result int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 3, 1}, 1},
		{[]int{1, 0, 0}, 1},
		{[]int{100}, 1},
		{[]int{0, 5, 0, 9, 7, 8}, 4},
		{[]int{3, 5, 0, 9, 7, 8}, 4},
		{[]int{1, 7, 9, 4}, 3},
		{[]int{12, 56, 77, 22}, 4},
		{[]int{0, 56, 0, 0}, 1},
		{[]int{0, 1, 0, 0}, 1},
		{[]int{0, 1, 5, 0}, 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0274_hindex(c.input))
	}
}

// Сортируем по возрастанию и идём с конца. Два возможных результата:
// - наименьшее число, который больше, чем количество пройденных элементов.
// - наибольшее число, которое меньше или равно количеству пройденных элементов.
// Находим оба результата и возвращаем наилучший (ну или какой нашли).
func lc_0274_hindex(citations []int) int {
	slices.Sort(citations) // Сортируем по возрастанию.
	lastBiggerIndex, firstSmaller, c := -1, -1, 0
	for i := len(citations) - 1; i >= 0; i-- { // Идём с конца в начало.
		if c = citations[i]; c != 0 {
			if c <= len(citations)-i {
				firstSmaller = c
				break // Если мы нашли "первый меньший, чем количество элементов "впереди", то уже точно нашли и "последний больший".
			}
			lastBiggerIndex = i // Нашли элемент, который больше, чем количество элементов "впереди".
		} else {
			break // Как только пошли нули, выходим, так как это не повлияет ни на что.
		}
	}
	switch {
	case lastBiggerIndex != -1 && firstSmaller != -1: // Выбираем или наилучший результат, или один из полученных.
		if l := len(citations) - lastBiggerIndex; l > firstSmaller {
			return l
		}
		return firstSmaller
	case lastBiggerIndex != -1:
		return len(citations) - lastBiggerIndex
	case firstSmaller != -1:
		return firstSmaller
	default:
		return 0 // Попались одни нули.
	}
}

// 0 1 2 3 4
// 0 0 5 7 8 9

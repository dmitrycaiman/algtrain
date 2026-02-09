package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/trapping-rain-water
func Test_lc_0042_trapping_rain_water(t *testing.T) {
	cases := []struct {
		height []int
		result int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{3, 0, 0, 0, 0, 4, 0}, 12},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0042_trapping_rain_water(c.height))
	}
}

// Идём по высотам, кладём "очки движения" в стек при спуске и извлекаем их из стека при подъёме (если возможно).
func lc_0042_trapping_rain_water(height []int) int {
	stack, result := [][]int{}, 0
	for i := 1; i < len(height); i++ {
		points := height[i-1] - height[i] // Количество "очков движения".
		switch {
		case points < 0 && len(stack) != 0: // Отрицательное количество: движемся вверх.
			points *= -1
			for points > 0 && len(stack) != 0 { // Работаем, пока есть очки подъёма и стек не пуст.
				lastDesc := stack[len(stack)-1] // Вынимаем последние очки спуска.
				if lastDesc[0] > points {       // Очков спуска больше полученных очков подъёма.
					result += points * (i - lastDesc[1] - 1) // Все очки подъёма уходят в результат.
					lastDesc[0] -= points                    // Тратим очки подъёма, модифицируя элемент в стеке.
					break                                    // Выходим, так как полностью потратили очки подъёма.
				} else { // Очков спуска меньше полученных очков подъёма:
					result += lastDesc[0] * (i - lastDesc[1] - 1) // Все очки спуска уходят в результат.
					points -= lastDesc[0]                         // Тратим очки подъёма.
					stack = stack[:len(stack)-1]                  // Удаляем элемент в стеке.
				}
			}
		case points > 0: // Положительное количество: движемся вниз.
			stack = append(stack, []int{points, i - 1}) // При движении вниз кладём в стек очки и то место, где произошёл спуск.
		}
	}
	return result
}

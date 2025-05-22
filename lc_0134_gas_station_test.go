package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0134_gas_station(t *testing.T) {
	assert.Equal(t, 3, lc_0134_gas_station([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	assert.Equal(t, 0, lc_0134_gas_station([]int{10, 0, 0, 0, 1}, []int{2, 2, 2, 4, 1}))
	assert.Equal(t, -1, lc_0134_gas_station([]int{10, 0, 0, 0, 0}, []int{2, 2, 2, 4, 1}))
	assert.Equal(t, 0, lc_0134_gas_station([]int{0, 0, 0, 0, 0, 1, 0}, []int{0, 0, 0, 0, 0, 1, 0}))
	assert.Equal(t, 3, lc_0134_gas_station([]int{0, 0, 3, 10, 0, 1, 4}, []int{0, 0, 4, 5, 4, 1, 1}))

}

// Выбираем нулевую колонку и пытаемся проделать от неё циклический путь. Если на колонке N мы не смогли продолжить, то для следующей
// итерации необходимо выбрать колонку N+1 (предыдущие рассматривать нет смысла). Если N+1 выходит за длину массива, то решение
// отсутствует. При совершении полного цикла возвращаем колонку, с которо начинали.
func lc_0134_gas_station(gas []int, cost []int) int {
	l := len(gas)
loop:
	for i := 0; i < l; {
		tank := 0
		for j := range l {
			k := i + j
			if k >= l {
				k -= l
			}
			tank += gas[k]
			if cost[k] > tank {
				if k >= i {
					i = k + 1
				} else {
					return -1
				}
				continue loop
			}
			tank -= cost[k]
		}
		return i
	}
	return -1
}

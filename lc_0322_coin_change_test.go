package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/coin-change
func Test_lc_0322_coin_change(t *testing.T) {
	cases := []struct {
		coins  []int
		amount int
		result int
	}{
		{[]int{4, 3, 2}, 0, 0},
		{[]int{4, 3, 2}, 10, 3},
		{[]int{2, 4, 3}, 11, 3},
		{[]int{2, 5, 10, 1}, 27, 4},
		{[]int{186, 419, 83, 408}, 6249, 20},
		{[]int{5, 1, 2}, 100, 20},
		{[]int{1, 2, 5}, 1000, 200},
	}
	for _, c := range cases {
		assert.Equal(t, c.result, lc_0322_coin_change_table(c.coins, c.amount))
		assert.Equal(t, c.result, lc_0322_coin_change_memo(c.coins, c.amount))
	}
}

// Табличный метод решения. Перебор из глубины с запоминанием наилучшего решения.
// Создаём кеш, элементы которого будут представлять все решения задачи от amount и ниже до 0. Решение F(0) = 0 есть первый элемент кеша.
// Итерируемся по массиву монет и пытаемся найти наилучшее решение для N+1 на основе уже известных: F(N+1) = 1 + F(N-m[i]).
// Если N-m[i] не существует или по данному индексу нет решения (-1), то не выполняем дальнейшие вычисления для элемента на этой итерации.
// Запоминаем наилучшее решение в кеш под индексом N+1 и переходим к поиску решения для N+2, итерируемся заново.
// Ответом будет значение в кеше под индексом, равным amount, т.к. это будет наилучшим решением на основе предыдущих.
func lc_0322_coin_change_table(coins []int, amount int) int {
	cache := make([]int, amount+1)
	for i := 1; i < len(cache); i++ {
		best := -1
		for _, v := range coins {
			if i-v >= 0 {
				cached := cache[i-v]
				if cached != -1 {
					res := 1 + cached
					if res < best || best == -1 {
						best = res
					}
				}
			}
		}
		cache[i] = best
	}
	return cache[amount]
}

// Рекурсивное решение с мемоизацией. По сути перебор с запоминанием уже произведённых вычислений.
// На каждом этапе рекурсии решаем вопрос о том, берём ли элемент в набор или нет.
// Если берём, то рекурсируем в подзадачу, где требуемый результат уменьшен на номинал монеты: F(N, i) = 1 + F(N-m[i]).
// Если не берём, то переходим к следующему элементу: F(N, i) = F(N, i+1).
// Рассчитываем оба результата и наилучший кладём в кеш под индексом [N, i].
// Остановка рекурсии в глубине происходит по выходу за пределы массива монет, либо по точному нахождению решения или взятию значения из кеша.
func lc_0322_coin_change_memo(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	storage := make([][]int, len(coins))
	for i := range storage {
		storage[i] = make([]int, amount)
	}
	return lc_0322_coin_change_memo_check(coins, 0, amount, storage)
}

func lc_0322_coin_change_memo_check(coins []int, i, amount int, storage [][]int) (res int) {
	if i == len(coins) {
		return -1
	}
	if m := storage[i][amount-1]; m != 0 {
		return m
	}
	defer func() { storage[i][amount-1] = res }()
	switch {
	case coins[i] == amount:
		return 1
	case coins[i] > amount:
		return lc_0322_coin_change_memo_check(coins, i+1, amount, storage)
	}
	incl := lc_0322_coin_change_memo_check(coins, i, amount-coins[i], storage) + 1
	excl := lc_0322_coin_change_memo_check(coins, i+1, amount, storage)
	switch {
	case incl == 0 && excl == -1:
		return -1
	case incl == 0 && excl != -1:
		return excl
	case incl != 0 && excl == -1:
		return incl
	}
	if incl < excl {
		return incl
	}
	return excl
}

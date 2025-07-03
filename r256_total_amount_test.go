package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Route256
// «Сумма к оплате»
// Дан массив цен за список продуктов, купленных в магазине. Товары с одинаковой стоимостью считаются одинаковыми.
// Требуется посчитать, сколько потребуется заплатить суммарно за весь поход в магазин при условии,
// что в магазине проводится акция — «купи три одинаковых товара и заплати только за два».
func Test_r256_total_amount(t *testing.T) {
	cases := []struct {
		prices []float64
		total  float64
	}{
		{[]float64{1.5, 1.5, 3.5, 4, 5.1, 1.5, 1.5, 1.5}, 18.6},
		{[]float64{1.1, 4.4, 6.6, 1.1, 1.1, 2.2, 6.6, 6.6, 6.6, 6.6, 1.2, 1.1, 1.1, 2, 5, 7, 8, 4.4, 4.4}, 65},
		{[]float64{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, 30},
		{[]float64{2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3}, 22},
		{[]float64{2, 3, 2, 3, 2, 2, 3, 2, 3, 2, 2, 3}, 22},
		{[]float64{10000}, 10000},
		{[]float64{1, 2, 3, 1, 2, 3, 1, 2, 3}, 12},
		{[]float64{10000, 10000, 10000, 10000, 10000, 10000}, 40000},
		{[]float64{300, 100, 200, 300, 200, 300}, 1100},
	}
	for _, c := range cases {
		assert.Equal(t, c.total, r256_total_amount_1(c.prices))
		assert.Equal(t, c.total, r256_total_amount_2(c.prices))
	}
}

// Создаём мапу, в которой ключ есть цена, значение есть количество товара.
// Из каждого значения выделяем целочисленным делением количество троек и умножаем его на двойную цену. Остаток от деления умножаем на цену.
// Всё складываем. По сути уменьшили все входящие тройки товаров до двоек.
func r256_total_amount_1(prices []float64) (res float64) {
	m := map[float64]int{}
	for _, price := range prices {
		m[price]++
	}
	for price, count := range m {
		res += price * (2*float64(count/3) + float64(count%3))
	}
	return
}

// Решение без использования дополнительной памяти.
// Основано на сортировке входного массива, что позволяет посчитать количество одинаковых товаров прямым проходом.
func r256_total_amount_2(prices []float64) float64 {
	for {
		swap := false
		for i := 0; i < len(prices)-1; i++ {
			if prices[i] > prices[i+1] {
				prices[i], prices[i+1] = prices[i+1], prices[i]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	price, count, res := float64(0), 0, float64(0)
	for _, p := range prices {
		if p != price {
			res += price * (2*float64(count/3) + float64(count%3))
			count = 1
			price = p
		} else {
			count++
		}
	}
	return res + price*(2*float64(count/3)+float64(count%3))
}

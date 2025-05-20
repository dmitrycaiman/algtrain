package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0067_add_binary(t *testing.T) {
	for range 100 {
		a, b := rand.Intn(1000), rand.Intn(1000)
		assert.Equal(t, fmt.Sprintf("%b", a+b), lc_0067_add_binary(fmt.Sprintf("%b", a), fmt.Sprintf("%b", b)))
	}
}

// Производим сложение "в столбик". Выбираем символы на одинаковых позициях каждой из входных строк. Двигаемся влево, т.е. с конца.
// Если какая-то из строк "кончилась", то рассматриваем её значение как "0". В зависимости от суммы байт символов ("0"=48, "1"=49) и флага
// переноса разряда (over) определяем результат в текущем разряде и значение флага переноса разряда для следующей итерации.
// Если в конце вычислений возведён флаг переноса разряда, добавляем к результату "1".
func lc_0067_add_binary(a string, b string) string {
	var m, n, over byte
	var res string
	for i := 1; i <= len(a) || i <= len(b); i++ {
		if i <= len(a) {
			m = a[len(a)-i]
		}
		if i <= len(b) {
			n = b[len(b)-i]
		}
		switch m + n + over {
		case 96:
			res, over = "0"+res, 0
		case 97:
			res, over = "1"+res, 0
		case 98:
			res, over = "0"+res, 1
		case 99:
			res, over = "1"+res, 1
		}
		m, n = 48, 48
	}
	if over != 0 {
		res = "1" + res
	}
	return res
}

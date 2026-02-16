package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var lc_0278_isBadVersion func(version int) bool

// https://leetcode.com/problems/first-bad-version
func Test_lc_0278_first_bad_version(t *testing.T) {
	for i := range 100 {
		lc_0278_isBadVersion = func(version int) bool { return version >= i }
		assert.Equal(t, i, lc_0278_first_bad_version(100))
	}
}

// Ищем точку "перелома" (т.е. где начинаются плохие версии) бинарным поиском.
// Проверяем левый край. Если проверка успешна, то возвращаем его как результат.
// Иначе проверяем середину и правый край:
// - если середина прошла проверку, а правый край нет, то плохие версии начинаются где-то справа.
// - иначе — где-то слева.
// При таких переходах точка перелома
func lc_0278_first_bad_version(n int) int {
	for a, b := 0, n; ; {
		if lc_0278_isBadVersion(a) {
			return a
		}
		m := a + (b-a)/2
		cm, cb := lc_0278_isBadVersion(m), lc_0278_isBadVersion(b)
		switch {
		case !cm && cb:
			a = m + 1
		case cm && cb:
			b = m
		}
	}
}

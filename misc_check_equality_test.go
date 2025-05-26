package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEquality(t *testing.T) {
	assert.True(t, CheckEquality([]string{"a", "b", "b", "c", "d"}, []string{"d", "c", "a", "b", "b"}))
	assert.False(t, CheckEquality([]string{"a", "b", "b", "c", "d"}, []string{"d", "c", "a", "b", "e"}))
	assert.False(t, CheckEquality([]string{"a", "b", "b", "c"}, []string{"d", "c", "a", "b", "b"}))

	assert.True(t, CheckEquality([]int{1, 2, 3, 4}, []int{2, 4, 3, 1}))
	assert.False(t, CheckEquality([]int{1, 2, 3, 4}, []int{2, 4, 3, 5}))
	assert.False(t, CheckEquality([]int{1, 2, 3, 4}, []int{2, 4, 3}))

	s1, s2 := [][]string{{"ab", "fu"}, {"y"}, {"ko", "ka", "ke"}}, [][]string{{"ka", "ke", "ko"}, {"fu", "ab"}, {"y"}}
	assert.True(t, CheckEqualityWithFunc(s1, s2, func(a, b []string) bool { return CheckEquality(a, b) }))

}

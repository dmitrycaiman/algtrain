package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEquality(t *testing.T) {
	assert.True(t, checkEquality([]string{"a", "b", "b", "c", "d"}, []string{"d", "c", "a", "b", "b"}))
	assert.False(t, checkEquality([]string{"a", "b", "b", "c", "d"}, []string{"d", "c", "a", "b", "e"}))
	assert.False(t, checkEquality([]string{"a", "b", "b", "c"}, []string{"d", "c", "a", "b", "b"}))

	assert.True(t, checkEquality([]int{1, 2, 3, 4}, []int{2, 4, 3, 1}))
	assert.False(t, checkEquality([]int{1, 2, 3, 4}, []int{2, 4, 3, 5}))
	assert.False(t, checkEquality([]int{1, 2, 3, 4}, []int{2, 4, 3}))

	s1, s2 := [][]string{{"ab", "fu"}, {"y"}, {"ko", "ka", "ke"}}, [][]string{{"ka", "ke", "ko"}, {"fu", "ab"}, {"y"}}
	assert.True(t, checkEqualityWithFunc(s1, s2, func(a, b []string) bool { return checkEquality(a, b) }))

}

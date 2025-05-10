package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0380_insert_delete_getrandom(t *testing.T) {
	r := Constructor()
	max := 5
	for i := range max {
		assert.True(t, r.Insert(i))
		assert.Equal(t, r.next-1, r.m[i])
		assert.Equal(t, i, r.s[r.next-1])
	}
	for i := range max {
		assert.False(t, r.Insert(i))
	}
	assert.Equal(t, max, r.next)
	for i := range max {
		delIdx := r.m[i]
		assert.True(t, r.Remove(i))
		assert.Zero(t, r.m[i])
		assert.Equal(t, r.s[delIdx], r.s[r.next])
	}
	for i := range max {
		assert.False(t, r.Remove(i))
	}
}

type RandomizedSet struct {
	next int
	s    []int
	m    map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{m: map[int]int{}, s: make([]int, 200_000)}
}

func (slf *RandomizedSet) Insert(val int) bool {
	_, ok := slf.m[val]
	if !ok {
		slf.m[val] = slf.next
		slf.s[slf.next] = val
		slf.next++
	}
	return !ok
}

func (slf *RandomizedSet) Remove(val int) bool {
	del, ok := slf.m[val]
	if ok {
		slf.next--
		last := slf.s[slf.next]
		slf.s[del], slf.m[last] = last, del
		delete(slf.m, val)
	}
	return ok
}

func (slf *RandomizedSet) GetRandom() int {
	return slf.s[rand.Intn(slf.next)]
}

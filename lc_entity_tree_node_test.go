package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeNode(t *testing.T) {
	cases := []string{
		"1,2,3",
		"5,4,8,11,null,13,4,7,2,null,null,5,1",
		"1,2,3,4,null,11,23,null,null,9,0,1",
	}
	for _, s := range cases {
		assert.Equal(t, s, NewTree(s).String())
	}
}

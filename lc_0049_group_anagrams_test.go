package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0049_group_anagrams(t *testing.T) {
	cases := []struct {
		input  []string
		output [][]string
	}{
		{
			input:  []string{"bo", "ob", "ko"},
			output: [][]string{{"bo", "ob"}, {"ko"}},
		},
		{
			input:  []string{"a", "oob", "b", "c", "boo", "obo"},
			output: [][]string{{"oob", "obo", "boo"}, {"a"}, {"b"}, {"c"}},
		},
	}
	for _, c := range cases {
		assert.True(t, CheckEqualityWithFunc(c.output, lc_0049_group_anagrams(c.input), func(a, b []string) bool { return CheckEquality(a, b) }))
	}
}

func lc_0049_group_anagrams(strs []string) [][]string {
	sort := func(in []byte) string {
		for {
			swap := false
			for i := 0; i <= len(in)-2; i++ {
				if in[i+1] > in[i] {
					in[i], in[i+1] = in[i+1], in[i]
					swap = true
				}
			}
			if !swap {
				return string(in)
			}
		}
	}
	m := map[string][]string{}
	for _, s := range strs {
		key := sort([]byte(s))
		m[key] = append(m[key], s)
	}
	result := [][]string{}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

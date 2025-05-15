package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lc_0071_simplify_path(t *testing.T) {
	cases := []struct{ input, output string }{
		{input: "home//dir/.///a/b//c/../../..////", output: "/home/dir"},
		{input: "./home//dir/.///a/b//c/../../..////", output: "/home/dir"},
		{input: "../home//dir/.///a/b//c/../../..////", output: "/home/dir"},
		{input: "////", output: "/"},
		{input: ".////", output: "/"},
		{input: "a////", output: "/a"},
		{input: "a//..//", output: "/"},
		{input: "/.//", output: "/"},
		{input: "./.", output: "/"},
	}
	for _, c := range cases {
		assert.Equal(t, c.output, lc_0071_simplify_path(c.input))
	}
}

func lc_0071_simplify_path(path string) string {
	prevIsSlash, omit, end := path[len(path)-1] == '/', 0, len(path)-1
	res := ""
	save := func(j int) {
		word := path[j : end+1]
		switch word {
		case ".":
		case "..":
			omit++
		default:
			if omit != 0 {
				omit--
			} else {
				res = "/" + word + res
			}
		}
	}
	for i := len(path) - 1; i >= 0; i-- {
		switch {
		case path[i] == '/' && !prevIsSlash:
			prevIsSlash = true
			save(i + 1)
		case path[i] != '/':
			if prevIsSlash {
				end = i
				prevIsSlash = false
			}
			if i == 0 {
				save(i)
			}
		}
	}
	if res == "" {
		return "/"
	}
	return res
}

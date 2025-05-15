package main

import (
	"flag"
	"os"
	"strings"
)

func main() {
	nameForGenerating := flag.String("g", "", "generate template for algo-task testing")
	flag.Parse()
	if name := *nameForGenerating; name != "" {
		generate(name)
	}
}

func generate(name string) {
	os.WriteFile(name+"_test.go", []byte(strings.Replace(template, "$", name, -1)), os.ModePerm)
}

const template = `
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_$(t *testing.T) {
	cases := []struct {}{}
	for _, c := range cases {
		_ = c
		$()
		assert.Equal(t, 1, 1)
	}
}

func $()  {}
`

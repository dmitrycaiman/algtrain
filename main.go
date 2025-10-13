package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fullName := flag.String("g", "", "generate simple template for algo-task testing by full task name")
	leetCodeName := flag.String("l", "", "generate template for LeetCode algo-task testing by name")
	flag.Parse()

	generateSimpleTemplate(fullName)
	generateLeetcodeTemplate(leetCodeName)
}

func generateLeetcodeTemplate(leetCodeName *string) {
	name := *leetCodeName
	if name == "" {
		return
	}
	// 123. Task Name -> lc_0123_task_name
	parts := strings.Split(name, ".")
	if len(parts) != 2 {
		log.Fatal("invalid LeetCode name format")
	}
	num, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	newSlug, prevSym := []byte{}, '-'
	for _, v := range parts[1] {
		switch {
		case v >= 'a' && v <= 'z':
		case v >= 'A' && v <= 'Z':
			v += 'a' - 'A'
		case v >= '0' && v <= '9':
		case v == ' ' && prevSym != '_':
			v = '_'
		default:
			continue
		}
		prevSym = v
		newSlug = append(newSlug, byte(v))
	}
	name = fmt.Sprintf(
		"lc_%04d%s",
		num,
		string(newSlug),
	)
	os.WriteFile(name+"_test.go", []byte(strings.ReplaceAll(simpleTemplate, "$", name)), os.ModePerm)
}

func generateSimpleTemplate(fullName *string) {
	name := *fullName
	if name == "" {
		return
	}
	os.WriteFile(name+"_test.go", []byte(strings.ReplaceAll(simpleTemplate, "$", name)), os.ModePerm)
}

const simpleTemplate = `
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://
func Test_$(t *testing.T) {
	cases := []struct {}{}
	for _, c := range cases {
		_ = c
		$()
		assert.Equal(t, 1, 1)
	}
}

// Описание решения...
func $()  {}
`

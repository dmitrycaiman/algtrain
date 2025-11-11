package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fullName := flag.String("g", "", "generate simple template for algo-task testing by full task name")
	leetCodeName := flag.String("lc", "", "generate template for LeetCode algo-task testing by name")
	codeRunName := flag.String("cr", "", "generate template for CodeRun algo-task testing by name")
	flag.Parse()

	generateSimpleTemplate(fullName)
	generateLeetcodeTemplate(leetCodeName)
	generateCodeRunTemplate(codeRunName)
}

func generateCodeRunTemplate(codeRunName *string) {
	name := *codeRunName
	if name == "" {
		return
	}

	res, err := http.Get(name)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	t := html.NewTokenizer(res.Body)
	for tokenType := t.Next(); tokenType != html.ErrorToken; tokenType = t.Next() {
		if tokenType == html.TextToken {
			token := t.Token()
			if token.Data[0] == '{' && token.Data[len(token.Data)-1] == '}' {
				fields := map[string]any{}
				if err = json.Unmarshal([]byte(token.Data), &fields); err != nil {
					continue
				}
				slug := fields["query"].(map[string]any)["slug"].([]any)[0].(string)
				writeTemplate(
					codeRunTemplate,
					"https://coderun.yandex.ru/problem/"+slug,
					fmt.Sprintf(
						"cr_%04.0f_%s",
						fields["props"].(map[string]any)["pageProps"].(map[string]any)["values"].(map[string]any)["2pxafbfw"].(map[string]any)["number"].(float64),
						strings.ReplaceAll(slug, "-", "_"),
					),
				)
				return
			}
		}
	}
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
	writeTemplate(
		simpleTemplate,
		"https://leetcode.com/problems/"+strings.ReplaceAll(string(newSlug[1:]), "_", "-"),
		fmt.Sprintf(
			"lc_%04d%s",
			num,
			string(newSlug),
		),
	)
}

func generateSimpleTemplate(fullName *string) {
	name := *fullName
	if name == "" {
		return
	}
	writeTemplate(simpleTemplate, "", name)
}

func writeTemplate(template, link, fileName string) {
	os.WriteFile(
		fileName+"_test.go",
		[]byte(strings.ReplaceAll(strings.ReplaceAll(template, "https://", link), "$", fileName)),
		os.ModePerm,
	)
}

const simpleTemplate = `package main

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

const codeRunTemplate = `package main

import (
	"bufio"
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://
func Test_$(t *testing.T) {
	cases := []struct{ input, output string }{
		{
			"",
			"",
		},
	}
	for _, c := range cases {
		w := bytes.NewBuffer(nil)
		$(bytes.NewBufferString(c.input), w)
		result, err := io.ReadAll(w)
		assert.NoError(t, err)
		assert.Equal(t, c.output, string(result))
	}
}

// func main() {
//   $(os.Stdin, os.Stdout)
// }

// Описание решения...
func $(r io.Reader, w io.Writer) {
	reader := bufio.NewReaderSize(r, 1<<20)
	writer := bufio.NewWriterSize(w, 1<<20)
	defer writer.Flush()
	_ = reader
}
`

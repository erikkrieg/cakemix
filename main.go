package main

import (
	"io"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/erikkrieg/cast/internal/values"
)

var mock = `
bool: true
str_quoted: "this is a test"
str_unquoted: this is a test
list:
- 1
- 2
- 3
nested:
  foo: bar
`

var filePath = "./test.json"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	casting, err := os.Open(filePath)
	check(err)
	vals, err := values.Parse(mock)
	check(err)
	check(renderCasting(vals, casting).Close())
}

func renderCasting(data map[interface{}]interface{}, casting io.Reader) *os.File {
	templateBytes, err := ioutil.ReadAll(casting)
	check(err)
	templateString := string(templateBytes)
	t, err := template.New("cast").Parse(templateString)
	check(err)
	newFile, err := os.Create("./tmp/cast.json")
	check(err)
	check(t.Execute(newFile, data))
	return newFile
}
